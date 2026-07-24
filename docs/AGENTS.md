# AGENTS.md — Working in this repository

Context for AI agents (and humans) contributing to this **Digital.ai Release
container plugin** built with the **Go SDK**. Read this first, then see the
[Plugin Development Guide](PLUGIN_DEVELOPMENT.md) for the full how-to. There is also a
portable [`SKILL.md`](SKILL.md) that routes to both.

## What this project is

- A **container plugin** for Digital.ai Release. It contributes custom **task types**.
- Each task is a Go command in [`my-integration/`](../my-integration/) wired through the
  **[`release-integration-sdk-go`](https://github.com/digital-ai/release-integration-sdk-go)**
  runner (`runner.Execute`, see [`main.go`](../main.go)).
- The compiled binary is copied into a **Docker image** and executed there by Release. The SDK
  runner reads the task input, resolves the matching command, runs its `FetchResult`, and sends
  results back to Release.
- Building produces **two artifacts**: a **plugin zip** (type definitions + icons from
  `resources/`, installed into Release) and a **Docker image** (the compiled Go binary, pushed to
  a registry). The `build.sh` / `build.bat` scripts produce both.

## The one rule you must not break: the type ↔ command naming contract

Release identifies a task by its **type** (e.g. `goContainerExamples.Hello`). At run time the SDK
runner:

1. receives the task type from Release;
2. looks it up in the **command factory** ([`my-integration/cmd/factory.go`](../my-integration/cmd/factory.go)) —
   the `commandHatchery` map keys are the exact type strings;
3. builds the matching command struct and calls its `FetchResult(ctx)` method.

Therefore, three things must stay in lockstep for every task:

- The type string in [`resources/type-definitions.yaml`](../resources/type-definitions.yaml).
- The **same** type string registered as a key in `commandHatchery` (`factory.go`), usually via a
  `const`.
- A command **struct** ([`cmd/commands.go`](../my-integration/cmd/commands.go)) with a
  `FetchResult` implementation ([`cmd/executors.go`](../my-integration/cmd/executors.go)).

A type declared in `type-definitions.yaml` with no matching factory entry (or vice-versa) is a bug.
Input-property JSON tags on the struct (e.g. `` `json:"yourName"` ``) must match the
`input-properties` names in the YAML.

## How to add a task (summary)

1. Declare the type in [`resources/type-definitions.yaml`](../resources/type-definitions.yaml)
   (extend `goContainerExamples.BaseTask`; declare `input-properties` / `output-properties`).
2. Add a struct in [`cmd/commands.go`](../my-integration/cmd/commands.go) with `json`-tagged input
   fields.
3. Register the type constant and a factory entry in
   [`cmd/factory.go`](../my-integration/cmd/factory.go) (`commandHatchery`).
4. Implement `FetchResult(ctx context.Context) (*task.Result, error)` in
   [`cmd/executors.go`](../my-integration/cmd/executors.go); keep the task logic in its own file
   under `cmd/example/`.
5. Add an integration test under `test/testdata/` and register it in `test/integration_test.go`.
6. Bump `VERSION` in `project.properties`, then build.

Inside `FetchResult`:

- Read inputs from the command struct's `json`-tagged fields (populated by the SDK deserializer).
- Build a result with `task.NewResult()` and its fluent setters
  (`.String(name, value)`, comments, etc.).
- **Return a non-nil `error` to fail the task** — the message is surfaced to the user. Wrap errors
  with context; do not swallow them.
- `ctx context.Context` is **required** on every command — it carries cancellation used for abort
  handling.

Use the injected **`releaseClient`** (from `factory.go`) when the task calls the **Release** REST
API, and the injected **`httpClient`** when it calls a **third-party** server connection CI.

## Abort support

Abort is opt-in per task:

1. Register `command.AbortCommand(<existingType>)` in `commandHatchery` (see `hello`).
2. Add a struct for the abort data in `cmd/commands.go`.
3. Implement `FetchResult` for that struct in `cmd/executors.go`.

Always thread `context.Context` through task logic so cancellation propagates.

## Dependencies

- Dependencies are managed with Go modules — [`go.mod`](../go.mod) and [`go.sum`](../go.sum).
- Add a dependency with `go get <module>`; run `go mod tidy` when it is directly relevant.
- The SDK (`release-integration-sdk-go`) is the primary dependency; keep its version consistent
  with the runner behaviour the tasks rely on.

## Build & deploy

`build.sh` / `build.bat` read [`project.properties`](../project.properties)
(`PLUGIN`, `VERSION`, `REGISTRY_URL`, `REGISTRY_ORG`), substitute the `@project.*@` /
`@registry.*@` placeholders in `resources/`, build the zip, and build & push the image.

- `./build.sh` — build zip + image, push image.
- `./build.sh --zip` / `--image` — build just one.
- `./build.sh --upload` — also upload the zip to Release (configure `.xebialabs/config.yaml`).

Treat `build.sh`, `build.bat`, and the [`Dockerfile`](../Dockerfile) as canonical build machinery —
change them only with explicit intent. The normal release edit is bumping `VERSION` in
`project.properties`.

## Tests

- Integration tests live in [`test/integration_test.go`](../test/integration_test.go) and use
  GoConvey. Fixtures and expected output live under `test/testdata/` and `test/fixtures/`.
- Run the suite from the repo root:

```sh
go test ./...             # all tests
go test -v ./test/...     # integration tests, verbose
```

- To add a test: create a folder under `test/testdata/` with `input.json` and `expected.json`,
  register the folder name in the `testsLabels` variable, and add any mock HTTP responses to
  `test/fixtures/` (wired via `test.MockResult{}` in `commandRunner`).

## Conventions & guardrails

- Run `gofmt -w` on edited Go files. Do not hand-format.
- Match the surrounding code style; keep each command small and single-purpose, with its logic in
  its own file under `cmd/example/`.
- Keep `type-definitions.yaml`, the `commandHatchery` keys, and the command structs in lockstep
  (the naming contract), and keep `json` tags aligned with `input-properties` names.
- Return errors with context; preserve `context.Context` cancellation for abort handling.
- Don't hand-edit generated build outputs (`build/`, `tmp/`) — they are produced by the scripts.
