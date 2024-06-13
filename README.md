# _Template Project for Digital.ai Release Integrations_

_This project serves as a template for developing a Go-based container plugin._

_See [How to create a new project](#how-to-create-a-new-project) below_

---

# Digital.ai Release integration to TARGET by PUBLISHER

⮕ Insert description here ⬅

---
## How to build and run

This section describes the quickest way to get a setup with Release to test containerized plugins using the SDK Development environment. For a production setup, please refer to the documentation. <!-- XXX insert link to documentation -->

### Prerequisites

You need to have the following installed in order to develop Go-based container tasks for Release using this project:

* Go 1.21
* Git
* Docker

### Start Release and Release Remote Runner

We will run Release and Release Remote Runner within a local Docker environment. In the development setup, the Release will trigger execution of containerized task on local Docker run Remote Runner.

Start the Release and Remote Runner environment with the following command

```commandline
cd dev-environment
docker compose up -d --build
```

>**Note:** Before running you can change default password fore `remote-runner` user in `dev-environment/digitalai-release-setup/secrets.xlvals` if needed - be sure to set a password with special char, numeric value, upper case letter and long enough, or secure enough to be up to security compliancy in Release.

### Configure your `hosts` file

The Release server needs to be able to find the container images of the integration you are creating. In order to do so the development setup has its own registry running inside Docker. Add the address of the registry to your local machine's `hosts` file.

**Unix / macOS**

Add the following entry to `/etc/hosts` (sudo privileges is required to edit):

    127.0.0.1 container-registry

**Windows**

Add the following entry to `C:\Windows\System32\drivers\etc\hosts` (Run as administrator permission is required to edit):

    127.0.0.1 container-registry


### Build & publish the plugin

Build will generate a zip and a docker image pushed to the registry defined in `project.properties`

1. Set properties in `project.properties`
2. Run build script to build the plugin zip and publish the image to registry:

**Unix / macOS**

```commandline
sh build.sh 
```

**Windows**

```commandline
build.bat 
```
The above command builds the zip, creates the container image, and then pushes the image to the configured registry.

`build.bat --zip` Builds the zip.

`build.bat --image` Creates the container image, and then pushes the image to the configured registry.

### Install plugin into Release

There are two ways to install the plugin into Release.

**Install plugin via commandline**

Update the Release server details in `.xebialabs/config.yaml`

Run the command for Unix / macOS:
```commandline
sh build.sh --upload 
```

Run the command for Windows:
```commandline
build.bat --upload 
```
The above command builds the zip and image and uploads the zip to the release server.

**Install plugin via Release server UI**

In the Release UI, use the Plugin Manager interface to upload the zip from `build`.
The zip takes the name of the project, for example `release-integration-template-go-0.0.1.zip`.

Then:
* Refresh the UI by pressing Reload in the browser.

### 5. Test it!

Create a template with the task **Go Container Example: Hello** and run it!

### 6. Clean up

Stop the development environment with the following command:

    docker compose down

---

## How to create a new project

The  [release-integration-template-go](https://github.com/digital-ai/release-integration-template-go) repository is a template project.

On the main page of this repository, click **Use this template** button, and select **Create new repository**. This will create a duplicate of this project to start developing your own container-based integration.

**Naming conventions**

- `my-integration` folder (as well as the package `my_integration`) should be renamed after the integration target name. 
All task logic should be implemented inside this folder.\
(**Note:** *Go doesn't encourage usage of `-` and `_` in package names, try to keep package name short, single word, but still clear. In this example `-` was used on intention with intention for you to refactor it.*)


Use the following naming convention for developing Digital.ai Release integration plugins:

    [publisher]-release-[target]-integration

Where publisher would be the name of your company.

For example:

    acme-release-example-integration

### Repository configuration

In the new project, update `project.properties` with the name of the integration plugin

```commandline
cd acme-release-example-integration
```

Change the following line in `project.properties`:

```
PLUGIN=acme-release-example-integration
...
```
### Add a new task

1. Add task type to `type-definitions.yaml`.
2. Add task struct with input parameters to `cmd/commands.go`.
3. Add task type to constants and command factory in `cmd/factory.go`.
4. Add `FetchResult()` implementation of command in `cmd/executors.go` and add task logic.

**_NOTE:_** Although task logic is inside `FetchResult()` method, it is a good practice to create a new file for each new task (See examples).

### Add abort logic for a task

View examples at [Abort Example](my-integration/cmd/example)

1. Define abort command for a task in `cmd/factory.go`. Use following syntax `command.AbortCommand(NAME_OF_EXISTING_COMAND): func...` (See example for `hello`)
2. In `cmd/commands.go`, define a struct that will hold the necessary data for abort execution.
3. In `cmd/executors.go`, define a method on the newly created struct which implements `FetchResults`

**_NOTE:_** Make sure to include context.Context in your methods as it is now required because of changes made to support abort functionality.




### Integration tests

Integration tests execution is implemented in `test/integration_test.go` using Convey. 

#### Add a new integration test:
1. Create a new folder inside `test/testdata`.
2. Add `input.json` (provided input) and `expected.json` (expected output) files inside the new folder.
3. Add the folder name to `testsLabels` variable in `test/integration_test.go`.

#### Add a new mock HTTP response:
1. Add JSON file with corresponding response to `test/fixtures`.
2. Add `test.MockResult{}` to `commandRunner` in `test/integration_test.go`.
