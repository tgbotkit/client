package client_test

import (
	"encoding/json"
	"testing"

	"github.com/tgbotkit/client"
)

func TestBotAPI105RequestSerialization(t *testing.T) {
	t.Parallel()

	enabled := true
	callbackQueryMode := "inline"
	settings, err := json.Marshal(client.SetMySettingsJSONRequestBody{
		CallbackQueryMode:    &callbackQueryMode,
		SupportsGuestQueries: &enabled,
	})
	if err != nil {
		t.Fatalf("marshal setMySettings request: %v", err)
	}
	if got, want := string(settings), `{"callback_query_mode":"inline","supports_guest_queries":true}`; got != want {
		t.Errorf("setMySettings request = %s, want %s", got, want)
	}

	widget, err := json.Marshal(client.SendWidgetJSONRequestBody{
		AllowQueries: &enabled,
		ChatId:       -100123,
		Html:         "<b>Hello</b>",
	})
	if err != nil {
		t.Fatalf("marshal sendWidget request: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(widget, &decoded); err != nil {
		t.Fatalf("unmarshal sendWidget request: %v", err)
	}
	if got := decoded["chat_id"]; got != float64(-100123) {
		t.Errorf("chat_id = %v, want -100123", got)
	}
	if got := decoded["html"]; got != "<b>Hello</b>" {
		t.Errorf("html = %v, want %q", got, "<b>Hello</b>")
	}
	if got := decoded["allow_queries"]; got != true {
		t.Errorf("allow_queries = %v, want true", got)
	}
}
