package components

import (
	"context"

	"github.com/atos-digital/DHCW-clinic-outcomes/internal/middleware"
)

func CreateFollowupName(name string) string {
	return name + "-followup"
}

func IsChecked(ctx context.Context, key, value string) bool {
	session := middleware.SessionFromContext(ctx)
	data, ok := session.Values["outcomes-form-data"].(map[string]string)
	if !ok {
		return false
	}
	return data[key] == value
}
