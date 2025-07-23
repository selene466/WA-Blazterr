package main

import (
	"WA-Blazterr/src/dbinit"
	"WA-Blazterr/src/whatsapp"
	"context"
	"fmt"

	waLog "go.mau.fi/whatsmeow/util/log"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	dbinit.Init()
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) StartWhatsapp() {
	whatsapp.Start()
}

func (a *App) ShowWhatsapp() waLog.Logger {
	return whatsapp.GetLogger()
}
