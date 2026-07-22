<!--
  Starter README for a plugin generated from the Release integration template.
  After cloning the template, replace the template's README with this file:

      mv README-plugin.md README.md      # (Windows: move /Y README-plugin.md README.md)

  Then fill in the placeholders below: the title, the one-line description, and the
  plugin name in `project.properties`.
-->

# <Your Plugin Name>

![Go](https://img.shields.io/badge/go-1.26%2B-blue)
[![release-integration-sdk-go](https://img.shields.io/badge/release--integration--sdk--go-GitHub-orange)](https://github.com/digital-ai/release-integration-sdk-go)
![License: MIT](https://img.shields.io/badge/license-MIT-green)

<!-- One line on what this plugin does. -->
A Digital.ai Release **container plugin**. It contributes custom **task types** — each is a
Go command in [`my-integration/`](my-integration/) that is compiled into a binary, packaged into a
Docker image, and run by Release as a container task. Built from the
[Release integration template](https://github.com/digital-ai/release-integration-template-go).

The task code is built on the **[`release-integration-sdk-go`](https://github.com/digital-ai/release-integration-sdk-go)** —
commands are wired through its runner (`runner.Execute`) to read inputs, set outputs, and call the
Release APIs. It is the project's main dependency and is pinned in [`go.mod`](go.mod).

Building this project produces **two artifacts**:

- a **plugin zip** — the plugin metadata from `resources/`, installed into Release.
- a **Docker image** — the compiled Go binary, pushed to a container registry and run by Release.

> [!TIP]
> **Adding or changing tasks?** See the **[Plugin Development Guide](docs/PLUGIN_DEVELOPMENT.md)** —
> it explains how a container plugin works, how to add a task, and the type ↔ command naming contract.

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

On Windows, use `build.bat` with the same arguments, for example:

```powershell
.\build.bat --upload
```

## Install the plugin into Release

**Option A — command line**

Set your Release server details in [`.xebialabs/config.yaml`](.xebialabs/config.yaml),
then make sure the Release server is running and use the command for your platform:

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
(named `<PLUGIN>-<VERSION>.zip`, using the values in [`project.properties`](project.properties)),
then reload the browser.

## First successful run

Use this sequence to verify the complete local workflow:

1. Start the local Release, runner, and registry with `docker compose up -d --build`.
2. Wait for `Digital.ai Release has started.` in the Release container logs.
3. Add `127.0.0.1 container-registry` to your hosts file.
4. Build and install the plugin with `./build.sh --upload` or `.\build.bat --upload` on Windows.
5. Open <http://localhost:5516>, create a template, and add one of your task types.
6. Run the release and verify the task produces the expected output.

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
  The SDK powering this plugin's task runner, commands, and Release API clients.
- **[Digital.ai Go SDK Documentation](https://docs.digital.ai/release/docs/how-to/overview-go-sdk)** —
  Guide to using the Go SDK and building custom tasks.
- **[Digital.ai Release documentation](https://docs.digital.ai/)** —
  Product documentation for Digital.ai Release.

## License

See [LICENSE.md](LICENSE.md).
