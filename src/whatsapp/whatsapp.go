package whatsapp

import (
	"WA-Blazterr/src/dbsqlite"
	"WA-Blazterr/src/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"

	"github.com/mdp/qrterminal/v3"
)

type Whatsapp struct {
	ctx     context.Context
	appCtx  context.Context
	running bool
}

func (w *Whatsapp) SetAppContext(c context.Context) {
	w.appCtx = c
}

func newLogger(name string) waLog.Logger {
	return waLog.Stdout(name, "INFO", true)
}

var Log = newLogger("WA-Blazterr")

func GetLogger() waLog.Logger {
	return Log
}

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		Log.Infof("Received a message! %s", v.Message.GetConversation())
	}
}

func NewWhatsapp(ctx context.Context) *Whatsapp {
	return &Whatsapp{
		ctx:     ctx,
		running: false,
	}
}

func (w *Whatsapp) Start() {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	ctx := context.Background()
	dbFilePath, err := dbsqlite.FileDB()
	if err != nil {
		utils.ErrorDialog(w.appCtx, "Error Local Database", err.Error())
	}
	container, err := sqlstore.New(ctx, "sqlite3", "file:"+dbFilePath+"?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				// Render the QR code here
				// e.g. qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				Log.Infof("QR code: %s", evt.Code)
				qrterminal.GenerateHalfBlock(string(evt.Code), qrterminal.L, os.Stdout)
			} else {
				Log.Infof("Login event: %s", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}

	w.running = true
	utils.WhatsappEvent(w.appCtx, "running")

	// HTTP handler "/api/send"
	http.HandleFunc("/api/send", func(w http.ResponseWriter, r *http.Request) {
		phone := r.URL.Query().Get("phone")
		msg := r.URL.Query().Get("msg")

		if phone == "" || msg == "" {
			http.Error(w, "Missing phone or msg", http.StatusBadRequest)
			return
		}

		jid := types.NewJID(phone, types.DefaultUserServer)

		client.SendMessage(context.Background(), jid, &waE2E.Message{
			Conversation: proto.String(msg),
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "✅ Message sent to %s", phone)
	})

	// Start HTTP server
	http.ListenAndServe("0.0.0.0:4140", nil)

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}
