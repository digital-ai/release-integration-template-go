---
name: develop-release-integration-go
description: Build, test, and maintain this container-based integration plugin for Digital.ai Release (Go SDK). Use for project setup, defining task types and server connections in resources/type-definitions.yaml, implementing tasks in my-integration/cmd, writing GoConvey integration tests, and the build/deploy workflow. Routes to the detailed guide; does not duplicate it.
license: MIT
metadata:
  audience: developers
  workflow: digital-ai-release
---

# Develop a Release container plugin (Go)

A portable, tool-neutral skill for working in this template. It is concise on purpose: the
**authoritative content lives in the repo docs**, and this file routes you to the right one so
nothing is duplicated or drifts.

- **[AGENTS.md](AGENTS.md)** — conventions and guardrails. **Read first.**
- **[PLUGIN_DEVELOPMENT.md](PLUGIN_DEVELOPMENT.md)** — the detailed guide (architecture, type
  definitions, command patterns, examples, dev environment, troubleshooting, Kubernetes).
- **[README.md](../README.md)** — setup, build, and install commands.

## What this covers

Setting up the project · defining task types and server connections in
`resources/type-definitions.yaml` · implementing tasks in `my-integration/cmd` (struct + factory +
`FetchResult`) · writing GoConvey integration tests · building the plugin zip + Docker image and
installing into Release.

## The one rule you must not break

A task's **type** maps to a Go **command** through the factory. The same type string must appear in
`resources/type-definitions.yaml`, as a key in the `commandHatchery` map in
[`cmd/factory.go`](../my-integration/cmd/factory.go), and be backed by a struct in
[`cmd/commands.go`](../my-integration/cmd/commands.go) with a `FetchResult` in
[`cmd/executors.go`](../my-integration/cmd/executors.go). Struct `json` tags must match the YAML
`input-properties` names.
→ [details](AGENTS.md#the-one-rule-you-must-not-break-the-type--command-naming-contract)

## Route the request

| To… | Go to |
|-----|-------|
| Set up / configure a fresh clone | [README → Development](../README.md#development); name the plugin in `project.properties`; replace the template README with the starter (`mv README-plugin.md README.md`); rename the `my-integration` package for your target. |
| Add a new task | [PLUGIN_DEVELOPMENT.md → Add a new task](PLUGIN_DEVELOPMENT.md#add-a-new-task--step-by-step) (declare type → struct → factory → `FetchResult` → test → build). |
| Add abort logic | [PLUGIN_DEVELOPMENT.md → Abort support](PLUGIN_DEVELOPMENT.md#abort-support). |
| Understand the SDK task API or examples | [PLUGIN_DEVELOPMENT.md → Anatomy of a task](PLUGIN_DEVELOPMENT.md#anatomy-of-a-task) and [The example tasks explained](PLUGIN_DEVELOPMENT.md#the-example-tasks-explained). |
| Build / deploy / install | [AGENTS.md → Build & deploy](AGENTS.md#build--deploy) and [README → Build & publish](../README.md#build--publish). |
| Run the dev server or fix a stuck server | [PLUGIN_DEVELOPMENT.md → The development environment](PLUGIN_DEVELOPMENT.md#the-development-environment) and [Troubleshooting](PLUGIN_DEVELOPMENT.md#troubleshooting). |

## Always

- Keep `resources/type-definitions.yaml`, the `commandHatchery` keys, and the command structs in
  lockstep (the naming contract); keep `json` tags aligned with `input-properties` names.
- Thread `context.Context` through task logic so abort/cancellation works.
- Return errors with context; don't swallow them. Run `gofmt -w` on edited files.
- Run `go test ./...` before building.
