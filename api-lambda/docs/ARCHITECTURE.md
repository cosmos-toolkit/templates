# API Lambda (Cosmos Template)

This template is a minimal **AWS Lambda** function for **API Gateway HTTP API (v2)**. The handler receives `APIGatewayV2HTTPRequest` and returns `APIGatewayV2HTTPResponse`.

## Layout

```
api-lambda/
├── cmd/handler/          # Entrypoint: lambda.Start(handler.HandleAPIGatewayV2)
├── internal/handler/     # Request routing and business logic
├── pkg/env/              # Optional env helpers
├── configs/              # Example config / env
├── docs/
├── go.mod
├── Makefile
└── README.md
```

## Flow

1. API Gateway HTTP API (v2) invokes the Lambda with an `APIGatewayV2HTTPRequest` (path, method, body, headers).
2. `HandleAPIGatewayV2` routes by path and method (e.g. `/health`, `/api/v1/entities`, `/api/v1/entities/{id}`).
3. Handler returns `APIGatewayV2HTTPResponse` (status, headers, body).

## Extending

- **Use cases**: Add `internal/app` and `internal/domain` (same pattern as api-hexagonal), and call them from `internal/handler`.
- **Config**: Use `pkg/env.Get("TABLE_NAME", "")` or a config struct loaded from env in `main` or at first request.
- **Other triggers**: Add another `cmd` (e.g. `cmd/sqs`) and handler for SQS, SNS, EventBridge, etc., or use the same binary with multiple Lambda function configs pointing to different handler symbols.

## Build for Lambda

Lambda expects a Linux binary. The Makefile builds for `GOOS=linux GOARCH=amd64` and names the output `bootstrap` (required for the `provided.al2023` / `provided` runtime). Use `make zip` to produce a deployment package.
