# CLI (Cosmos Template)

Go CLI template with **Cobra** and **plug-and-play commands**: add subcommands in `internal/commands` or in `pkg/cli/commands` and register them on the root command. Same layout and conventions as the API/worker templates (internal, pkg, configs).

## Project structure

```
cli/
├── cmd/
│   └── cli/                 # Entrypoint
│       └── main.go
├── internal/
│   ├── cli/                 # Root command, wires subcommands
│   │   └── root.go
│   └── commands/            # Subcommands (version, run, …)
│       ├── version.go
│       └── run.go
├── pkg/                     # Shared packages and optional command sets
│   ├── env/
│   └── README.md            # How to add commands (internal or pkg)
├── configs/
├── docs/
│   └── ARCHITECTURE.md
├── go.mod
├── Makefile
└── README.md
```

## Plug-and-play commands

- **Default**: add a new file in `internal/commands/` (e.g. `migrate.go`) and in `internal/cli/root.go`: `rootCmd.AddCommand(commands.MigrateCmd)`.
- **From pkg**: add `pkg/cli/commands/migrate/command.go` that returns `*cobra.Command`, then in root: `rootCmd.AddCommand(migrate.Command())`.
- No other code changes; just register the command on root.

See `pkg/README.md` for step-by-step and shared packages (env, logger, database).

## How to use

1. Copy the template and set the `module` in `go.mod` (e.g. `github.com/your-org/your-cli`).
2. Add or replace commands in `internal/commands/` (migrate, seed, export, etc.).
3. Build and run:
   ```bash
   make build
   ./bin/cli version
   ./bin/cli run --name "Cosmos"
   # or: go run ./cmd/cli run -n "Cosmos"
   ```

## Built-in commands

| Command   | Description                             |
| --------- | --------------------------------------- |
| `version` | Print CLI version                       |
| `run`     | Sample task (e.g. `run --name "world"`) |

## Build and tests

```bash
make build   # bin/cli
make run     # run default (help)
make test    # run tests
make install # install to $GOBIN
make hello   # example: run --name "Cosmos"
```

## License

As per the Cosmos Toolkit project.
