# Template Project for Digital.ai Release Integrations

![Go](https://img.shields.io/badge/go-1.26%2B-blue)
[![release-integration-sdk-go](https://img.shields.io/badge/release--integration--sdk--go-GitHub-orange)](https://github.com/digital-ai/release-integration-sdk-go)
![License: MIT](https://img.shields.io/badge/license-MIT-green)

This project serves as a template for developing a Go-based **container plugin**
for Digital.ai Release. Each task is a Go command in [`my-integration/`](my-integration/)
that is compiled into a binary, packaged into a Docker image, and run by Release as a
container task.

The task code is built on the **[`release-integration-sdk-go`](https://github.com/digital-ai/release-integration-sdk-go)** —
commands are wired through its runner (`runner.Execute`) to read inputs, set outputs, and call
the Release APIs. It is the project's main dependency and is pinned in [`go.mod`](go.mod).

Building the project produces **two artifacts**:

- a **plugin zip** — the plugin metadata from `resources/`, installed into Release.
- a **Docker image** — the compiled Go binary, pushed to a container registry and run by Release.

> [!TIP]
> **Writing your own tasks?** Start with the **[Plugin Development Guide](docs/PLUGIN_DEVELOPMENT.md)** —
> it explains how a container plugin works, how to add a task, and how each bundled example was built.

> [!IMPORTANT]
> **Using this as a template?** This README documents the *template itself*. After you create
> your own repo from it, follow [After creating your repository](#after-creating-your-repository)
> to personalize the clone.

## Contents

- [After creating your repository](#after-creating-your-repository)
- [Quick start](#quick-start)
- [Project layout](#project-layout)
- [Prerequisites](#prerequisites)
- [Development](#development)
- [Run Release locally](#run-release-locally)
- [Build & publish](#build--publish)
- [Install the plugin into Release](#install-the-plugin-into-release)
- [First successful run](#first-successful-run)
- [Clean up the local environment](#clean-up-the-local-environment)
- [Related resources](#related-resources)
- [License](#license)

## After creating your repository

The [`release-integration-template-go`](https://github.com/digital-ai/release-integration-template-go)
repository is a template. On its main page, click **Use this template → Create a new repository**.
Then, before developing your integration, complete these steps:

1. Rename the [`my-integration/`](my-integration/) folder (and its package) after your integration
   target. All task logic lives here.
   > **Note:** Go discourages `-` and `_` in package names — keep the package name short, single
   > word, and clear. The `-` in `my-integration` is intentional, for you to refactor.
2. Set `PLUGIN`, `VERSION`, `REGISTRY_URL`, and `REGISTRY_ORG` in
   [`project.properties`](project.properties). Use the naming convention
   `[publisher]-release-[target]-integration` (e.g. `acme-release-example-integration`).
3. Remove or adapt the example tasks in `my-integration/cmd/`,
   [`resources/type-definitions.yaml`](resources/type-definitions.yaml), and `test/`.
4. Update the plugin description and task details in this README.
5. Run `go build ./...` and `go test ./...` before building.

The [`develop-release-integration-go`](docs/SKILL.md) skill guides you (or your AI agent)
through these steps.

## Quick start

From the repository root:

```sh
go build ./...                       # compile everything
docker compose up -d --build         # start Release, the runner, and the local registry
```

Wait for the Release container log to show `Digital.ai Release has started.` Before
building, add `127.0.0.1 container-registry` to your hosts file — this requires
administrator/`sudo` rights (see [Run Release locally](#run-release-locally)). Then
build and install the plugin. The `--upload` step reads your Release server details
from [`.xebialabs/config.yaml`](.xebialabs/config.yaml) (defaults point at the local
server):

```sh
# macOS / Linux
./build.sh --upload
```

```powershell
# Windows (PowerShell or Command Prompt)
.\build.bat --upload
```

The default local Release server is available at <http://localhost:5516>.
After the upload completes, create a template with the **Container Examples: Hello (Go)** task
(`goContainerExamples.Hello`) and run it. Each step is detailed below.

## Project layout

| Path                  | Purpose                                                                       |
|-----------------------|-------------------------------------------------------------------------------|
| `main.go`             | Entry point — wires the SDK runner and command factory. **Ships inside the Docker image.** |
| `my-integration/`     | Task implementations (`cmd/` structs, factory, executors, examples). **This code ships inside the Docker image.** See the [Plugin Development Guide](docs/PLUGIN_DEVELOPMENT.md). |
| `task/`               | Shared task helpers used by the integration (e.g. server connection deserialization). |
| `test/`               | GoConvey integration tests with `testdata/` and `fixtures/`. Not shipped in the image. |
| `resources/`          | Plugin metadata (`type-definitions.yaml`, icons) packaged into the plugin zip. |
| `go.mod` / `go.sum`   | Go module definition and dependency checksums. **Source of truth for the container.** |
| `Dockerfile`          | Builds the container image that runs the tasks.                               |
| `build.sh` / `build.bat` | Builds the plugin zip and the Docker image, and uploads them to Release.    |
| `project.properties`  | Plugin name, version, and registry coordinates used by the build scripts.     |
| `docker-compose.yaml` | A local Dockerized Release server (+ runner + container registry) for testing. |
| `dev-environment/`    | Build contexts and config used by `docker-compose.yaml`.                       |
| `docs/`               | Contributor docs: `PLUGIN_DEVELOPMENT.md` (detailed guide), `AGENTS.md` (conventions/guardrails for AI agents), and `SKILL.md` (portable `develop-release-integration-go` skill that routes to the docs above). |

## Prerequisites

- [Go 1.26+](https://go.dev/)
- [Docker](https://www.docker.com/) — to build and run the container image
- [Git](https://git-scm.com/)

## Development

Write and test tasks with the standard Go toolchain. The container image itself is built from the
module (`go build`) by the [`Dockerfile`](Dockerfile).

### Build and test

```sh
go build ./...            # compile everything
go test ./...             # run all tests
go test -v ./test/...     # run the integration tests, verbose
gofmt -w <files>          # format edited Go files
```

### Add a dependency

Dependencies are managed with Go modules. Add one and update `go.mod` / `go.sum`:

```sh
go get <module>
go mod tidy               # when directly relevant
```

The SDK, [`release-integration-sdk-go`](https://github.com/digital-ai/release-integration-sdk-go),
is the primary dependency.

## Run Release locally

Run a local Release server, its remote runner, and a container registry, using Docker.

```sh
docker compose up -d --build
```

### Configure your `hosts` file

Release must be able to reach the local container registry by name. Add this entry:

- **macOS / Linux** — `/etc/hosts` (requires `sudo`)
- **Windows** — `C:\Windows\System32\drivers\etc\hosts` (run as administrator)

```
127.0.0.1 container-registry
```

## Build & publish

The build scripts read `project.properties`, build the plugin zip from
`resources/`, build the Docker image from the `Dockerfile`, and push the image to
the configured registry.

The `--image`, default, and `--upload` workflows require Docker to be running and
the registry in `REGISTRY_URL` to be reachable. For the local Docker Compose
environment, start the stack first and add `127.0.0.1 container-registry` to your
hosts file as described in [Run Release locally](#run-release-locally). For a remote
registry, make sure Docker is authenticated and that `REGISTRY_URL` and
`REGISTRY_ORG` in [`project.properties`](project.properties) are correct.

| Command                | Result                                                        |
|------------------------|---------------------------------------------------------------|
| `./build.sh`           | Build the zip **and** the image, and push the image.          |
| `./build.sh --zip`     | Build only the plugin zip.                                    |
| `./build.sh --image`   | Build only the Docker image and push it.                      |
| `./build.sh --upload`  | Build the zip and image, push the image, and upload the zip to Release. |

On Windows, use `build.bat` with the same arguments (works in both PowerShell and
Command Prompt), for example:

```powershell
.\build.bat --upload
```

## Install the plugin into Release

**Option A — command line**

Set your Release server details in [`.xebialabs/config.yaml`](.xebialabs/config.yaml)
(the same file used by the [Quick start](#quick-start) `--upload` step), then make sure
the Release server is running and use the command for your platform:

```sh
# macOS / Linux
./build.sh --upload
```

```powershell
# Windows
.\build.bat --upload
```

**Option B — Release UI**

In the Release **Plugin Manager**, upload the zip from `build/`
(named `<PLUGIN>-<VERSION>.zip`, e.g. `release-integration-template-go-0.0.1.zip`
with the current [`project.properties`](project.properties)),
then reload the browser.

## First successful run

The [Quick start](#quick-start) covers the happy path end to end. Once the plugin is
installed, verify the workflow in the Release UI:

1. Open <http://localhost:5516>, create a template, and add
   **Container Examples: Hello (Go)** (`goContainerExamples.Hello`).
2. Run the release and verify the task produces its greeting output.

If a step fails, see [Run Release locally](#run-release-locally),
[Build & publish](#build--publish), and [Install the plugin into Release](#install-the-plugin-into-release)
for the full setup and troubleshooting details.

## Clean up the local environment

When you finish testing, stop the local Release server, runner, and registry:

```sh
docker compose down
```

This stops the containers but preserves the local registry/server data mounted under
`dev-environment/`. To reset the development environment and remove that test state,
run `docker compose down` and remove the generated contents under that directory before
starting the stack again.

## Related resources

- **[Digital.ai Release Go SDK](https://github.com/digital-ai/release-integration-sdk-go)** —
  The SDK powering this template's task runner, commands, and Release API clients.
- **[Digital.ai Go SDK Documentation](https://github.com/digital-ai/release-integration-sdk-go/wiki)** —
  Guide to using the Go SDK and building custom tasks.
- **[SDK Template Project for integration plugins](https://github.com/digital-ai/release-integration-template-go)** —
  A starting point for building custom integrations using Digital.ai Release and Go.
- **[Digital.ai Release documentation](https://docs.digital.ai/)** —
  Product documentation for Digital.ai Release.

## License

See [License.md](License.md).
