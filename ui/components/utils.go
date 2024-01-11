package components

import "strings"

func CreateFollowupLabel(label string) string {
	return strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(label, " / ", "-"), " ", "-") + "-followUp")
}
