package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

// HandleAPIGatewayV2 routes API Gateway HTTP API (v2) requests to the appropriate handler.
func HandleAPIGatewayV2(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	path := strings.TrimSuffix(req.RawPath, "/")
	if path == "" {
		path = "/"
	}
	method := req.RequestContext.HTTP.Method

	switch {
	case method == http.MethodGet && (path == "/health" || path == "/api/v1/health"):
		return health(ctx)
	case method == http.MethodGet && strings.HasPrefix(path, "/api/v1/entities/"):
		id := strings.TrimPrefix(path, "/api/v1/entities/")
		return getEntity(ctx, id)
	case method == http.MethodPost && path == "/api/v1/entities":
		return createEntity(ctx, req.Body)
	default:
		return jsonResponse(http.StatusNotFound, map[string]string{"error": "not found"})
	}
}

func health(ctx context.Context) (events.APIGatewayV2HTTPResponse, error) {
	body, _ := json.Marshal(map[string]string{
		"status": "ok",
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}, nil
}

func getEntity(ctx context.Context, id string) (events.APIGatewayV2HTTPResponse, error) {
	if id == "" {
		return jsonResponse(http.StatusBadRequest, map[string]string{"error": "id required"})
	}
	// Replace with your use case (e.g. call app.Service.GetEntity)
	entity := map[string]interface{}{
		"id":         id,
		"created_at": time.Now().UTC().Format(time.RFC3339),
		"updated_at": time.Now().UTC().Format(time.RFC3339),
	}
	return jsonResponse(http.StatusOK, entity)
}

func createEntity(ctx context.Context, body string) (events.APIGatewayV2HTTPResponse, error) {
	var input struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal([]byte(body), &input); err != nil || input.ID == "" {
		return jsonResponse(http.StatusBadRequest, map[string]string{"error": "invalid body, id required"})
	}
	// Replace with your use case (e.g. call app.Service.CreateEntity)
	entity := map[string]interface{}{
		"id":         input.ID,
		"created_at": time.Now().UTC().Format(time.RFC3339),
		"updated_at": time.Now().UTC().Format(time.RFC3339),
	}
	return jsonResponse(http.StatusCreated, entity)
}

func jsonResponse(status int, v interface{}) (events.APIGatewayV2HTTPResponse, error) {
	body, _ := json.Marshal(v)
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}, nil
}
