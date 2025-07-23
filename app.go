package main

import (
	"WA-Blazterr/src/dbsqlite"
	dialog "WA-Blazterr/src/utils"
	"WA-Blazterr/src/whatsapp"
	"context"
	"fmt"

	waLog "go.mau.fi/whatsmeow/util/log"
)

// App struct
type App struct {
	ctx      context.Context
	whatsapp *whatsapp.Whatsapp
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		whatsapp: whatsapp.NewWhatsapp(context.Background()),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	if err := dbsqlite.Init(); err != nil {
		dialog.ErrorDialog(a.ctx, "Error Local Database", err.Error())
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ShowWhatsapp() waLog.Logger {
	return whatsapp.GetLogger()
}
