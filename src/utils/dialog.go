package utils

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func ErrorDialog(ctx context.Context, title string, message string) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   title,
		Message: message,
	})
}

func InfoDialog(ctx context.Context, title string, message string) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   title,
		Message: message,
	})
}
