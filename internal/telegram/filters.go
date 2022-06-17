package telegram

import (
	tm "github.com/and3rson/telemux/v2"
)

// isEditCommand filters that match "edit" regular expression.
// For example, isEditCommand will handle commands like "/e_42"
func isEditCommand() tm.FilterFunc {
	return tm.And(tm.IsAnyCommandMessage(), tm.HasRegex("(/e_(\\d*))"))
}

// isCloseCommand filters that match "close" regular expression.
// For example, isCloseCommand will handle commands like "/c_69"
func isCloseCommand() tm.FilterFunc {
	return tm.And(tm.IsAnyCommandMessage(), tm.HasRegex("(/c_(\\d*))"))
}

func isPageCallback() tm.FilterFunc {
	return tm.And(tm.IsCallbackQuery(), tm.HasRegex("(page_(\\d*))"))
}
