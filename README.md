# Firefly Programming Language

Firefly is a simple, educational programming language that mimics some aspects of x86_64 assembly. It features custom instructions designed to simulate the behavior of fireflies, making it an engaging way to learn basic programming concepts.

## Features

- Custom instruction set inspired by firefly behavior
- Interactive command-line interface
- Ability to run scripts from files
- Mimics certain aspects of x86_64 assembly structure

## Getting Started

### For Developers

To run the Firefly interpreter as a developer:

1. Ensure you have Go installed on your system.
2. Clone this repository using:
    ```
    git clone https://github.com/HugeFrog24/firefly-lang.git
    ```
3. Navigate to the project directory:
    ```
    cd firefly-lang
    ```
4. Run the interpreter using:
    ```
    go run main.go
    ```

    To run a Firefly script file:
    ```
    go run main.go path/to/script.tni
    ```

### For End Users

To run the Firefly interpreter as an end user:

1. Download the latest release for your operating system.
2. Open a terminal or command prompt.
3. Navigate to the directory containing the Firefly executable.
4. Run the interpreter in interactive mode:

    ```
    ./firefly
    ```

    To run a Firefly script file:

    ```
    ./firefly path/to/script.tni
    ```

## Usage

Once the interpreter is running, you can enter Firefly instructions interactively or run them from a script file. Use the `HELP` command to see a list of available instructions and their usage.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
