# Your SDK (Cosmos Template)

Go **library (SDK)** template for a reusable client. Public API in the root package; optional subpackages for errors or helpers. Use this when the deliverable is a **library** consumed by other projects, not a service or CLI.

## Project structure

```
sdk/
├── client.go              # Main client and options
├── types.go               # Public types (Resource, etc.)
├── client_test.go         # Tests for the public API
├── examples/
│   └── basic/
│       └── main.go        # Example usage
├── go.mod
├── Makefile
└── README.md
```

## Conventions

- **Public API**: Exported types and functions in the root package (or clearly documented subpackages). Avoid breaking changes; use options (e.g. `WithBaseURL`) for new behaviour.
- **Versioning**: Follow [semantic versioning](https://semver.org/). Tag releases (e.g. `v1.0.0`). Go modules use the latest tag by default.
- **Stability**: Prefer a single major version (v1) with additive changes. Use a `CHANGELOG.md` and document deprecations.

## How to use (as a consumer)

1. Add the module to your project:
   ```bash
   go get github.com/your-org/your-sdk@v1.0.0
   ```
2. In your code:

   ```go
   import "github.com/your-org/your-sdk"

   client := sdk.NewClient(
       sdk.WithBaseURL("https://api.example.com"),
       sdk.WithAPIKey("key"),
   )
   if err := client.Ping(ctx); err != nil {
       return err
   }
   resource, err := client.GetResource(ctx, "id1")
   ```

## How to develop (as the SDK author)

1. Copy the template and set the `module` in `go.mod` (e.g. `github.com/your-org/your-sdk`).
2. Replace `client.go` and `types.go` with your real API client and types. Keep options in `client.go` or `options.go`.
3. Add more methods and types; keep the package name (e.g. `sdk`) and avoid breaking the existing API.
4. Run tests and the example:
   ```bash
   make test
   make example   # runs examples/basic/main.go
   ```
5. Tag a release: `git tag v1.0.0 && git push origin v1.0.0`.

## Build and test

```bash
make test    # run tests
make lint    # golangci-lint
make example # run examples/basic
```

## Examples

- `examples/basic`: minimal usage (NewClient, Ping, GetResource).

Add more examples under `examples/` (e.g. `examples/advanced`) and document them in the README.

## License

As per the Cosmos Toolkit project.
