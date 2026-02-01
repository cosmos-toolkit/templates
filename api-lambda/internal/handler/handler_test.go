package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleAPIGatewayV2_Health(t *testing.T) {
	req := events.APIGatewayV2HTTPRequest{
		RawPath: "/health",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: http.MethodGet,
			},
		},
	}
	resp, err := HandleAPIGatewayV2(context.Background(), req)
	if err != nil {
		t.Fatalf("HandleAPIGatewayV2: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("health: got status %d", resp.StatusCode)
	}
	var out map[string]string
	if err := json.Unmarshal([]byte(resp.Body), &out); err != nil {
		t.Fatalf("body: %v", err)
	}
	if out["status"] != "ok" {
		t.Errorf("status: got %q", out["status"])
	}
}

func TestHandleAPIGatewayV2_CreateEntity(t *testing.T) {
	body := `{"id":"e1"}`
	req := events.APIGatewayV2HTTPRequest{
		RawPath: "/api/v1/entities",
		Body:    body,
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: http.MethodPost,
			},
		},
	}
	resp, err := HandleAPIGatewayV2(context.Background(), req)
	if err != nil {
		t.Fatalf("HandleAPIGatewayV2: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("create: got status %d, body %s", resp.StatusCode, resp.Body)
	}
}
