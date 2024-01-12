package components

import (
	"context"
	"strings"

	"github.com/atos-digital/DHCW-clinic-outcomes/internal/middleware"
)

func CreateFollowupLabel(label string) string {
	return strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(label, " / ", "-"), " ", "-") + "-followUp")
}

func IsChecked(ctx context.Context, groupName, label string) bool {
	session := middleware.SessionFromContext(ctx)
	checked, ok := session.Values[groupName].([]string)
	if !ok {
		return session.Values[groupName] == label
	}
	for _, v := range checked {
		if v == label {
			return true
		}
	}
	return false
}
