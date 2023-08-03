# KeyManager

KeyManager is a Command-Line Interface (CLI) tool developed in pure Go for managing keys. This open-source project is aimed at providing a simple and efficient way to create, copy, track history, and delete keys. Future versions of the project intend to support various types of secrets, RSA keys, and more.

## Features

1. Create a new key
2. Copy an existing key
3. Print copy history for a key
4. Delete a key with confirmation

## Design Considerations

In designing KeyManager, we aimed for a solution that is self-contained. All the data is stored within the application itself, with no external dependencies or the need for any additional storage solutions. This decision was intentional, as it provides the ease of portability and ensures that users don't need to worry about setting up or configuring additional databases or storage solutions.

## Libraries Used

KeyManager leverages a couple of essential libraries to offer a user-friendly experience:

- [Cobra](https://github.com/spf13/cobra): This library is used for creating powerful modern CLI applications, and it drives the command-line interactions in KeyManager.

- [Pretty-Table](https://github.com/jedib0t/go-pretty): This library is used for creating formatted and visually appealing printouts in the terminal.

## Installation

### Install from source

1. Clone the repository:
```sh
git clone https://github.com/VladimirsDuhovs/keyManager.git
```
2. Navigate to the project directory:

```sh
cd keyManager
```

3. Build the project:

```sh
go build
```

### Install with a script

Each installation script is tailored for a specific operating system and architecture. To install the latest version of KeyManager for your specific system, use the relevant script from the table below:

```sh
curl -fsSL https://raw.githubusercontent.com/VladimirsDuhovs/keyManager/main/installation_scripts/<installation_script_for_your_os_architecture>.sh | sh
```

This script will download the appropriate binary and place it in your `/usr/local/bin` directory.

## Installation Scripts

| OS    | Architecture | Script |
| ------|:------------:| ------:|
| Darwin (Mac) | AMD64 | [Download](https://raw.githubusercontent.com/VladimirsDuhovs/keyManager/main/installation_scripts/install_keyManager_darwin_amd64.sh) |
| Darwin (Mac) | ARM64 | [Download](https://raw.githubusercontent.com/VladimirsDuhovs/keyManager/main/installation_scripts/install_keyManager_darwin_arm64.sh) |
| Linux  | AMD64 | [Download](https://raw.githubusercontent.com/VladimirsDuhovs/keyManager/main/installation_scripts/install_keyManager_linux_amd64.sh) |
| Linux  | ARM64 | [Download](https://raw.githubusercontent.com/VladimirsDuhovs/keyManager/main/installation_scripts/install_keyManager_linux_arm64.sh) |

## Commands

Here is a list of available commands:

- Create a new key:

```sh
keyManager create --key <key_name>
```

- Copy a key to a specific path:

```sh
keyManager copy --key <key_name> --output <path_to_copy>
```

- Print the copy history of a key:

```sh
keyManager history --key <key_name>
```

- Delete a key (requires user confirmation):
```sh
keyManager delete --key <key_name>
```

## Contributing

All the contributions from the open-source community are welcome. After all it is the community that makes all of this possible.

Please read the [CONTRIBUTING.md](CONTRIBUTING.md) file for more information.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Author

- Vladimirs Duhovs - Initial work

Feel free to reach out if you have any questions, or just want to say hi. Feedback and suggestions are always welcome.
