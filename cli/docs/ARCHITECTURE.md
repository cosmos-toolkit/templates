# CLI plug-and-play architecture

This template is a **CLI** built with Cobra. Commands are **pluggable**: add them in `internal/commands` or in `pkg/cli/commands` and register them on the root command in `internal/cli/root.go`.

## Structure

```
                    +------------------+
                    | cmd/cli          |  main.go → cli.Execute()
                    +--------+---------+
                             |
                    +--------v---------+
                    | internal/cli    |  root command, wires subcommands
                    +--------+---------+
                             |
                    +--------v---------+
                    | internal/commands|  version, run, ... (or pkg/cli/commands)
                    +------------------+
```

### cmd/cli

- **main.go**: calls `cli.Execute()`. No other logic.

### internal/cli

- **root.go**: root `*cobra.Command`, short/long description, and `init()` adding subcommands.
- To plug a command: `rootCmd.AddCommand(commands.YourCmd)` or `rootCmd.AddCommand(migrate.Command())` from pkg.

### internal/commands

- One file per command (e.g. `version.go`, `run.go`).
- Each file exports a `*cobra.Command` and optionally uses `init()` to add flags.
- Replace or extend with your own commands (migrate, seed, export, etc.).

### pkg/cli/commands (optional)

- Optional command sets as packages (e.g. `pkg/cli/commands/migrate`, `pkg/cli/commands/db`).
- Each package exposes `Command() *cobra.Command` and is added to root in `internal/cli/root.go`.
- Keeps the core CLI small and lets you add feature-specific command bundles.

### pkg (shared)

- **pkg/env**: env vars / .env for commands that need config.
- Other shared utilities (logger, database client) used by commands.

## Naming (Go)

- **Directories**: lowercase, single word (`cli`, `commands`).
- **Files**: `snake_case.go` for multiple words (`db_migrate.go`).
- **Commands**: exported `*cobra.Command` (e.g. `VersionCmd`, `RunCmd`).

## Adding a new command

1. Create `internal/commands/hello.go` (or `pkg/cli/commands/hello/command.go`).
2. Define the command and flags.
3. In `internal/cli/root.go`: `rootCmd.AddCommand(commands.HelloCmd)` (or `hello.Command()`).

No other code changes required (plug-and-play).
