package client_test

import (
	"encoding/json"
	"testing"

	"github.com/tgbotkit/client"
)

func TestEphemeralMessageParametersSerialization(t *testing.T) {
	t.Parallel()

	callbackQueryID := "callback-1"
	replace := true
	markdown := "Working"
	parameters := &client.EphemeralMessageParameters{
		CallbackQueryId:             &callbackQueryID,
		ReceiverUserId:              42,
		ReplaceCallbackQueryMessage: &replace,
	}

	tests := map[string]any{
		"sendMessage": client.SendMessageJSONRequestBody{
			ChatId:                     -100123,
			EphemeralMessageParameters: parameters,
			Text:                       "Working",
		},
		"sendRichMessage": client.SendRichMessageJSONRequestBody{
			ChatId:                     -100123,
			EphemeralMessageParameters: parameters,
			RichMessage:                client.InputRichMessage{Markdown: &markdown},
		},
	}

	for name, body := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			data, err := json.Marshal(body)
			if err != nil {
				t.Fatalf("marshal request: %v", err)
			}

			var decoded map[string]any
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("unmarshal request: %v", err)
			}

			if _, ok := decoded["receiver_user_id"]; ok {
				t.Error("request contains legacy top-level receiver_user_id")
			}
			if _, ok := decoded["callback_query_id"]; ok {
				t.Error("request contains legacy top-level callback_query_id")
			}

			ephemeral, ok := decoded["ephemeral_message_parameters"].(map[string]any)
			if !ok {
				t.Fatal("request does not contain ephemeral_message_parameters object")
			}
			if got := ephemeral["receiver_user_id"]; got != float64(42) {
				t.Errorf("receiver_user_id = %v, want 42", got)
			}
			if got := ephemeral["callback_query_id"]; got != callbackQueryID {
				t.Errorf("callback_query_id = %v, want %q", got, callbackQueryID)
			}
			if got := ephemeral["replace_callback_query_message"]; got != true {
				t.Errorf("replace_callback_query_message = %v, want true", got)
			}
		})
	}
}
