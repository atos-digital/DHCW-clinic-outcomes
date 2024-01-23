package components

import (
	"context"
	"encoding/json"

	"github.com/atos-digital/DHCW-clinic-outcomes/internal/middleware"
)

func IsChecked(ctx context.Context, key, value string) bool {
	session := middleware.SessionFromContext(ctx)
	var data map[string]string
	b, ok := session.Values["outcomes-form-data"]
	if !ok {
		return false
	}
	json.Unmarshal([]byte(b.(string)), &data)
	return data[key] == value
}
