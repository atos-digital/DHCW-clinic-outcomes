package components

import (
	"context"

	"github.com/atos-digital/DHCW-clinic-outcomes/internal/middleware"
)

func CreateFollowupName(name string) string {
	return name + "-followup"
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
