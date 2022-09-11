package sms

import (
	"time"
)

type json map[string]string
type sms_request_bool_exp map[string]interface{}

// SendSmsInput represents send sms input payload
type SendSmsInput struct {
	ClientName string      `json:"client_name,omitempty" graphql:"client_name"`
	TemplateID string      `json:"template_id,omitempty" graphql:"template_id"`
	Content    string      `json:"content,omitempty" graphql:"content"`
	Recipient  []Recipient `json:"recipient" graphql:"recipient"`
	SendAfter  time.Time   `json:"send_after,omitempty" graphql:"send_after"`
	Save       bool        `json:"save"`
	Locale     string      `json:"locale"`
}

// Recipient represents the international phone number
type Recipient struct {
	PhoneCode   int    `json:"phone_code"`
	PhoneNumber string `json:"phone_number"`
}

// SendSmsResponse represents sms response from external service
type SendSmsResponse struct {
	Success   bool        `json:"success" graphql:"success"`
	RequestID *string     `json:"request_id,omitempty" graphql:"request_id"`
	MessageID string      `json:"message_id,omitempty" graphql:"message_id"`
	Error     interface{} `json:"error,omitempty" graphql:"error"`
}

// SendSmsOutput represents the summary result of sending sms
type SendSmsOutput struct {
	Responses    []SendSmsResponse `json:"responses" graphql:"responses"`
	SuccessCount int               `json:"success_count" graphql:"success_count"`
	FailureCount int               `json:"failure_count" graphql:"failure_count"`
}
