# Firefly Programming Language

Firefly is a simple, educational programming language that mimics some aspects of x86_64 assembly. It features custom instructions designed to simulate the behavior of fireflies, making it an engaging way to learn basic programming concepts.

## Features

- Custom instruction set inspired by firefly behavior
- Interactive command-line interface
- Ability to run scripts from files
- Mimics certain aspects of x86_64 assembly structure, including object-subject instruction order

## Getting Started

### For Developers

To run the Firefly interpreter as a developer:

1. **Ensure you have Go installed** on your system.
2. **Clone this repository** using:
    ```bash
    git clone https://github.com/HugeFrog24/firefly-lang.git
    ```
3. **Navigate to the project directory**:
    ```bash
    cd firefly-lang
    ```
4. **Run the interpreter** using:
    ```bash
    go run main.go
    ```

    To run a Firefly script file:
    ```bash
    go run main.go path/to/script.tni
    ```

#### Using Makefile

Developers can also manage build tasks using the provided Makefile:

- **Build the Interpreter**:
    ```bash
    make build
    ```
- **Clean Build Artifacts**:
    ```bash
    make clean
    ```
- **Rebuild the Interpreter**:
    ```bash
    make rebuild
    ```

This approach simplifies the build process and ensures consistency across different development environments.

### For End Users

To run the Firefly interpreter as an end user:

1. **Download the latest release** for your operating system.
2. **Open a terminal or command prompt**.
3. **Navigate to the directory** containing the Firefly executable.
4. **Run the interpreter** in interactive mode:

    ```bash
    ./firefly
    ```

    To run a Firefly script file:

    ```bash
    ./firefly path/to/script.tni
    ```

## Usage

Once the interpreter is running, you can enter Firefly instructions interactively or run them from a script file. Use the `HELP` command to see a list of available instructions and their usage.

### Example

Here's an example of using Firefly in interactive mode:

```assembly
$ ./firefly
Firefly Programming Language v1.0.0
Enter instructions (one per line).
Type 'HELP' for help, 'EXIT' to finish:

> LIGHT TIB, 10       ; Tibik starts with 10 energy
TIB's light initialized with energy 10

> LIGHT NIN, 5        ; Nini starts with 5 energy
NIN's light initialized with energy 5

> GIFT NIN, TIB, 3    ; Nini receives 3 energy from Tibik
NIN received 3 energy from TIB. NIN now has 8 energy, TIB now has 7 energy.

> FLY TIB, 2          ; Tibik flies, losing 2 energy
TIB flew and lost 2 energy. Current energy: 5

> FLY NIN, 1          ; Nini flies, losing 1 energy
NIN flew and lost 1 energy. Current energy: 7

> SHOW LIGHT NIN      ; Displays Nini's energy level
NIN's energy level: 7

> HUG                 ; Both recharge to maximum energy (10)
All fireflies hugged and recharged to full energy (10)

> SHOW LIGHT TIB      ; Displays Tibik's energy level
TIB's energy level: 10

> EXIT
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
