package message

import "errors"

var (
	ErrInvalidValidation = errors.New("failed to validate arguments")

	ErrMessageNotFound = errors.New("message not found")
	ErrChatNotFound    = errors.New("chat not found")
	ErrDeleteMessage   = errors.New("failed to delete message")
	ErrInsertMessage   = errors.New("failed to insert message")
	ErrAccessDenied    = errors.New("access denied")
)
