package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/tgbotkit/client"
)

func ExampleNewClient() {
	serverURL, err := client.NewServerUrlTelegramBotAPIEndpointSubstituteBotTokenWithYourBotToken(
		client.ServerUrlTelegramBotAPIEndpointSubstituteBotTokenWithYourBotTokenBotTokenVariable("<bot_token>"),
	)
	if err != nil {
		return
	}

	c, err := client.NewClient(
		serverURL,
		client.WithHTTPClient(&http.Client{}),
		client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("X-Request-Id", "example")
			return nil
		}),
	)
	if err != nil {
		return
	}

	_ = c
}

func ExampleClientWithResponses_GetMeWithResponse() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/bot<bot_token>/getMe" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true,"result":{"id":123,"is_bot":true,"first_name":"TgBotKit"}}`))
	}))
	defer ts.Close()

	serverURL, err := client.NewServerUrlTelegramBotAPIEndpointSubstituteBotTokenWithYourBotToken(
		client.ServerUrlTelegramBotAPIEndpointSubstituteBotTokenWithYourBotTokenBotTokenVariable("<bot_token>"),
	)
	if err != nil {
		return
	}

	baseURL, err := url.Parse(serverURL)
	if err != nil {
		return
	}
	testURL, err := url.Parse(ts.URL)
	if err != nil {
		return
	}
	baseURL.Scheme = testURL.Scheme
	baseURL.Host = testURL.Host

	c, err := client.NewClientWithResponses(
		baseURL.String(),
		client.WithHTTPClient(ts.Client()),
	)
	if err != nil {
		return
	}

	resp, err := c.GetMeWithResponse(context.Background())
	if err != nil {
		return
	}
	if resp.JSON200 != nil {
		fmt.Println(resp.JSON200.Result.FirstName)
	}

	// Output: TgBotKit
}
