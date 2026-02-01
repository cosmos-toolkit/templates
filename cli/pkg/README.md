# Plug-and-play commands and packages

CLI commands and shared utilities live here or in `internal/commands`. The root command wires subcommands in `internal/cli/root.go`.

## Adding commands

### From internal (default)

1. Create a new file in `internal/commands/` (e.g. `migrate.go`).
2. Define a `*cobra.Command` and register it in `internal/cli/root.go`:
   ```go
   rootCmd.AddCommand(commands.MigrateCmd)
   ```

### From pkg (plug-and-play)

1. Create a package under `pkg/cli/commands/` (e.g. `pkg/cli/commands/migrate/command.go`).
2. Export a function that returns `*cobra.Command`:
   ```go
   package migrate
   func Command() *cobra.Command { return migrateCmd }
   ```
3. In `internal/cli/root.go`, import and add:
   ```go
   import "github.com/your-org/your-cli/pkg/cli/commands/migrate"
   rootCmd.AddCommand(migrate.Command())
   ```

This way you can add optional command sets (e.g. `pkg/cli/commands/db` for migrate/seed) without touching core logic.

## Shared packages

Use `pkg/` for utilities shared across commands or with other apps (API, worker):

- **pkg/env** – environment variables / .env (used by commands that need config).
- **pkg/logger** – structured logging (inject in commands that need it).
- **pkg/database** – DB client (used by migrate/seed commands).

Commands stay thin: parse flags, call a service or pkg, and write to stdout/stderr.
