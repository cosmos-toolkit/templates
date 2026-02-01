# API Lambda (Cosmos Template)

Go **AWS Lambda** template for **API Gateway HTTP API (v2)**. Single handler routes requests by path and method; build a Linux binary and deploy to Lambda.

## Project structure

```
api-lambda/
├── cmd/
│   └── handler/              # Entrypoint: lambda.Start(handler.HandleAPIGatewayV2)
│       └── main.go
├── internal/
│   └── handler/               # Request routing and responses
│       ├── handler.go
│       └── handler_test.go
├── pkg/
│   └── env/                   # Optional env helpers
├── configs/
├── docs/
│   └── ARCHITECTURE.md
├── go.mod
├── Makefile
└── README.md
```

## Prerequisites

- Go 1.22+
- AWS account (for deployment)

## How to use

1. Copy the template and set the `module` in `go.mod` (e.g. `github.com/your-org/your-lambda`).
2. Replace or extend `internal/handler` with your routes and use cases (optionally add `internal/app` and `internal/domain` as in api-hexagonal).
3. Build for Lambda:
   ```bash
   make build   # bin/bootstrap (Linux amd64)
   make zip     # lambda.zip for upload
   ```
4. Create a Lambda function with runtime **Go 1.x** or **provided.al2023** and handler set to the binary name (e.g. `bootstrap`). Connect it to an API Gateway HTTP API (v2) with a proxy integration.

## Routes (default)

| Method | Path                  | Description      |
| ------ | --------------------- | ---------------- |
| GET    | /health               | Health check     |
| GET    | /api/v1/entities/{id} | Get entity by ID |
| POST   | /api/v1/entities      | Create entity    |

## Build and test

```bash
make build   # bin/bootstrap (Linux)
make test    # run tests
make lint    # golangci-lint
make zip     # lambda.zip
```

## Deployment

- **Manual**: Upload `lambda.zip` (or `bin/bootstrap` in a zip) to Lambda. Set handler to `bootstrap` for provided runtime.
- **SAM**: Add a `template.yaml` with `AWS::Serverless::Function` and `Events: Api` for HTTP API.
- **Terraform / CDK**: Use the same binary; point the Lambda resource to the zip and set the handler accordingly.

## Config

Use environment variables in the Lambda function configuration (or a .env in SAM local). Example: `TABLE_NAME`, `LOG_LEVEL`. See `pkg/env` and `configs/config.example.yaml`.

## License

As per the Cosmos Toolkit project.
