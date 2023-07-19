# PromptPal CLI[![goreleaser](https://github.com/PromptPal/cli/actions/workflows/release.yaml/badge.svg)](https://github.com/PromptPal/cli/actions/workflows/release.yaml)

> **Warning**
> This project is currently in an early stage of development and may not be suitable for production use. Use with caution and expect frequent updates and changes.

PromptPal CLI is a command line interface tool designed in Golang to streamline prompt management and code generation for the PromptPal project. It allows developers to initialize a configuration file and generate code structures for seamless integration with Golang and Node.js applications.

## Features

- **Configuration Initialization**: The CLI provides the `promptpal init` command to initialize a configuration file. This file allows users to specify the server URL and authentication details required for communication with the PromptPal server.

- **Code Generation**: The CLI offers the `promptpal g` command to generate code structures based on the configuration file. By utilizing the configuration settings, the CLI retrieves prompt information from the PromptPal server and automatically generates code snippets tailored to Golang and Node.js applications.

- **Seamless Integration**: The generated code structures can be easily integrated into Golang and Node.js applications, enabling developers to leverage the prompt actions provided by the PromptPal project.

## Installation (Example for Unix-like Operating Systems)

To install the PromptPal CLI, follow these steps:

1. Go to the [Release](https://github.com/PromptPal/cli/releases) page on GitHub.

2. Download the appropriate prebuilt binary package for your operating system and architecture.

3. Unzip or extract the downloaded package. In the extracted folder, you will find the `promptpal` binary.

4. Open a terminal and navigate to the extracted folder containing the `promptpal` binary.

5. To make the `promptpal` command accessible system-wide, run the following command:

   ```bash
   sudo mv ./promptpal /usr/local/bin
   ```

This command moves the promptpal binary to the /usr/local/bin directory, which is typically included in the system's PATH. You may need to enter your administrator password for the command to execute successfully.

You can now use the promptpal command in your terminal to interact with the PromptPal CLI.
Please note that the steps above are an example for Unix-like operating systems. For Windows systems, the process may differ.

## Usage
The usage instructions for the PromptPal CLI remain the same as described in the previous sections of this README.

### Configuration Initialization
To initialize the configuration file, use the following command:

```bash
promptpal init
```

This command will prompt you to enter the server URL and authentication token for the PromptPal server. It will create a configuration file (promptpal.yaml) in the current directory with the specified details.

#### Configuration File Example

An example of a promptpal.yaml configuration file:

```yaml
input:
    http:
        url: "http://localhost:7788"
        # token: d6e9a6b170784fdfb4ef54417a32f391
        token: "@env.PROMPTPAL_API_TOKEN"
output:
    schema: "./schema.g.json"
    go_types:
        prefix: PP
        output: "./example/types.g.go"
        package_name: "ppp"
    typescript_types:
        prefix: PP
        output: "./example/types.g.ts"
```

In this example, you need to specify the input.http.url and input.http.token according to your PromptPal server configuration. Additionally, you can modify the output section to customize the generated code's schema file, Golang types, and TypeScript types.

### Code Generation
To generate code structures based on the configuration file, use the following command:

```bash
promptpal g
```

The CLI will read the configuration from the promptpal.yaml file and retrieve prompt information from the PromptPal server. It will then generate the necessary code structures according to the specified output paths and configurations.

### Help
For more information on available commands and options, use the --help flag.

```bash
promptpal --help
```
## To Contributors

Thank you for considering contributing to the PromptPal CLI project! Your contributions are valuable in improving the tool and making it more robust. Here are the steps to follow for releasing a new version:

1. Set up a local `$GITHUB_TOKEN` environment variable. This token will be used for authenticating with the GitHub API during the release process.

2. Set up an Apple personal token by visiting [https://appleid.apple.com/account/manage](https://appleid.apple.com/account/manage). This token is required for certain operations related to building and signing macOS binaries. For more details, please refer to the documentation of [gon](https://github.com/mitchellh/gon#troubleshooting).

3. Merge the "latest release" merge request into the master branch. This merge request should include all the necessary changes for the new release, including bug fixes, new features, and documentation updates.

4. Wait for the new release to be created on GitHub. Once the release is created, run `git pull` on your local machine to ensure that you have the latest changes.

5. Run the following command in your local macOS environment to upload all the artifacts:

   ```bash
   goreleaser --clean
   ```
This command will build the project, create release artifacts, and upload them to the GitHub release page.

Please note that these steps are specific to the release process of the PromptPal CLI project. If you have any questions or need further assistance, feel free to reach out to us.

We appreciate your contributions and look forward to your continued support!