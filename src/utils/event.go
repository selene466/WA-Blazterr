package utils

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func WhatsappEvent(ctx context.Context, inp string) {
	runtime.EventsEmit(ctx, "whatsappEvent", inp)
}
