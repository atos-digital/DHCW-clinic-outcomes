package components

import (
	"context"

	"github.com/atos-digital/DHCW-clinic-outcomes/internal/middleware"
)

type LabelOpts struct {
	Label    string
	Required bool
	Tooltip  string
	Bold     bool
}

func DefaultLabel(label string) LabelOpts {
	return LabelOpts{
		Label: label,
	}
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
