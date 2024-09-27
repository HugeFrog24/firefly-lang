package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Constants
const (
	maxEnergy    = 10
	numRegisters = 4 // Define the number of registers (fireflies) available
)

// Fixed register names
var registerNames = []string{"TIB", "NIN", "ABC", "XYZ"}

// registers holds the energy levels for each firefly
var registers = make(map[string]int)

// InstructionDefinition defines the structure for instruction specifications
type InstructionDefinition struct {
	Name    string
	MinArgs int
	MaxArgs int
	Usage   string
}

// instructionDefinitions maps each instruction to its definition
var instructionDefinitions = map[string]InstructionDefinition{
	"LIGHT": {
		Name:    "LIGHT",
		MinArgs: 2,
		MaxArgs: 2,
		Usage:   "LIGHT <name> <energy> - Set firefly's energy (0-10)",
	},
	"GIFT": {
		Name:    "GIFT",
		MinArgs: 3,
		MaxArgs: 3,
		Usage:   "GIFT <giver> <receiver> <amount> - Transfer energy between fireflies",
	},
	"FLY": {
		Name:    "FLY",
		MinArgs: 2,
		MaxArgs: 2,
		Usage:   "FLY <name> <amount> - Reduce firefly's energy by flying",
	},
	"HUG": {
		Name:    "HUG",
		MinArgs: 0,
		MaxArgs: 0,
		Usage:   "HUG - Recharge all fireflies to full energy",
	},
	"SHOW": {
		Name:    "SHOW",
		MinArgs: 2,
		MaxArgs: 2,
		Usage:   "SHOW LIGHT <name> - Display firefly's current energy",
	},
	"HELP": {
		Name:    "HELP",
		MinArgs: 0,
		MaxArgs: 0,
		Usage:   "HELP - Display usage information",
	},
	"CLEAR": {
		Name:    "CLEAR",
		MinArgs: 0,
		MaxArgs: 0,
		Usage:   "CLEAR - Clear the terminal",
	},
	"DIM": {
		Name:    "DIM",
		MinArgs: 1,
		MaxArgs: 1,
		Usage:   "DIM <name> - Reset firefly's energy to 0",
	},
}

func main() {
	fmt.Println("Firefly Programming Language v1.0.0")

	if len(os.Args) > 2 {
		fmt.Println("Error: Too many arguments")
		printUsage()
		os.Exit(1)
	}

	// Initialize registers (energy levels) to 0
	for _, name := range registerNames {
		registers[name] = 0
	}

	var instructions []string
	var err error

	if len(os.Args) == 2 {
		filename := os.Args[1]
		if filepath.Ext(filename) != ".tni" {
			fmt.Println("Error: File must have .tni extension")
			printUsage()
			os.Exit(1)
		}
		instructions, err = readInstructionsFromFile(filename)
	} else {
		instructions, err = readInstructionsFromStdin()
	}

	if err != nil {
		fmt.Printf("Error reading instructions: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Executing Firefly Language Instructions:")
	for _, instruction := range instructions {
		fmt.Printf("\nExecuting: %s\n", instruction)
		if err := executeInstruction(instruction); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func readInstructionsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var instructions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, ";") {
			instructions = append(instructions, line)
		}
	}
	return instructions, scanner.Err()
}

func readInstructionsFromStdin() ([]string, error) {
	var instructions []string
	scanner := bufio.NewScanner(os.Stdin)

	// Use newline characters to format the prompt
	fmt.Println("Enter instructions (one per line).\nType 'HELP' for help, 'EXIT' to finish:")
	fmt.Println() // This will print a blank line explicitly

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if strings.ToUpper(line) == "EXIT" {
			break
		}
		if strings.ToUpper(line) == "HELP" {
			printUsage()
			continue
		}
		if line != "" && !strings.HasPrefix(line, ";") {
			instructions = append(instructions, line)
			// Execute the instruction immediately
			if err := executeInstruction(line); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
	return instructions, scanner.Err()
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  1. Run with a .tni file:")
	fmt.Println("     go run main.go <filename.tni>")
	fmt.Println("  2. Run interactively:")
	fmt.Println("     go run main.go")
	fmt.Println("\nFirefly Language Instructions:")
	for _, def := range instructionDefinitions {
		fmt.Printf("  %s\n", def.Usage)
	}
	fmt.Println("\nComments start with a semicolon (;)")
	fmt.Println("\nNotes:")
	fmt.Printf("  - <name> must be one of the predefined three-letter names: %s\n", strings.Join(registerNames, ", "))
	fmt.Println("  - Energy levels are stored in registers (0-10)")
}

func executeInstruction(instruction string) error {
	// Remove inline comments
	instruction = strings.Split(instruction, ";")[0]
	parts := strings.Fields(instruction)
	if len(parts) == 0 {
		return fmt.Errorf("empty instruction")
	}

	op := strings.ToUpper(parts[0])

	// Retrieve the instruction definition
	def, exists := instructionDefinitions[op]
	if !exists {
		return fmt.Errorf("unknown instruction: %s", op)
	}

	// Sanitize arguments by removing trailing commas
	args := make([]string, len(parts[1:]))
	for i, arg := range parts[1:] {
		args[i] = strings.TrimRight(arg, ",")
	}

	// Validate the number of arguments using a helper function
	if len(args) < def.MinArgs {
		return fmt.Errorf("incomplete %s instruction. Usage: %s", op, def.Usage)
	}
	if len(args) > def.MaxArgs {
		return fmt.Errorf("too many arguments for %s instruction. Usage: %s", op, def.Usage)
	}

	// Execute based on instruction type
	switch op {
	case "LIGHT":
		name := strings.ToUpper(args[0])
		energy, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid energy value '%s'. Use a number between 0 and %d", args[1], maxEnergy)
		}
		return setEnergy(name, energy)
	case "GIFT":
		giver := strings.ToUpper(args[0])
		receiver := strings.ToUpper(args[1])
		amount, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("invalid gift amount: %v", err)
		}
		return gift(giver, receiver, amount)
	case "FLY":
		name := strings.ToUpper(args[0])
		amount, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid fly amount: %v", err)
		}
		return fly(name, amount)
	case "HUG":
		return hug()
	case "SHOW":
		if strings.ToUpper(args[0]) != "LIGHT" {
			return fmt.Errorf("invalid SHOW instruction format. Usage: %s", def.Usage)
		}
		name := strings.ToUpper(args[1])
		return showLight(name)
	case "HELP":
		printUsage()
		return nil
	case "CLEAR":
		return clearTerminal()
	case "DIM":
		if len(args) != 1 {
			return fmt.Errorf("invalid DIM instruction. Usage: %s", def.Usage)
		}
		name := strings.ToUpper(args[0])
		return dim(name)
	default:
		return fmt.Errorf("unknown instruction: %s", op)
	}
}

func setEnergy(name string, energy int) error {
	if _, exists := getRegisterIndex(name); !exists {
		return fmt.Errorf("unknown firefly name '%s'. Use one of: %s", name, strings.Join(registerNames, ", "))
	}

	if energy < 0 || energy > maxEnergy {
		return fmt.Errorf("energy must be between 0 and %d, got %d", maxEnergy, energy)
	}

	if registers[name] != 0 {
		return fmt.Errorf("%s's light is already initialized with energy %d", name, registers[name])
	}

	registers[name] = energy
	fmt.Printf("%s's light initialized with energy %d\n", name, energy)
	return nil
}

func gift(giver, receiver string, amount int) error {
	if _, existsGiver := getRegisterIndex(giver); !existsGiver {
		return fmt.Errorf("invalid giver: %s. Use predefined names: %s", giver, strings.Join(registerNames, ", "))
	}
	if _, existsReceiver := getRegisterIndex(receiver); !existsReceiver {
		return fmt.Errorf("invalid receiver: %s. Use predefined names: %s", receiver, strings.Join(registerNames, ", "))
	}

	giverEnergy := registers[giver]

	if amount < 0 || amount > giverEnergy {
		return fmt.Errorf("invalid gift amount: %d", amount)
	}

	registers[giver] -= amount
	registers[receiver] = min(registers[receiver]+amount, maxEnergy)
	fmt.Printf("%s gave %d energy to %s\n", giver, amount, receiver)
	fmt.Printf("%s now has %d energy, %s now has %d energy\n", giver, registers[giver], receiver, registers[receiver])
	return nil
}

func fly(name string, amount int) error {
	if _, exists := getRegisterIndex(name); !exists {
		return fmt.Errorf("unknown firefly: %s. Use predefined names: %s", name, strings.Join(registerNames, ", "))
	}

	energy := registers[name]
	if amount < 0 || amount > energy {
		return fmt.Errorf("invalid fly amount: %d", amount)
	}

	registers[name] -= amount
	fmt.Printf("%s flew and lost %d energy. Current energy: %d\n", name, amount, registers[name])
	return nil
}

func hug() error {
	for name := range registers {
		registers[name] = maxEnergy
	}
	fmt.Printf("All fireflies hugged and recharged to full energy (%d)\n", maxEnergy)
	return nil
}

func showLight(name string) error {
	if _, exists := getRegisterIndex(name); !exists {
		return fmt.Errorf("unknown firefly: %s. Use predefined names: %s", name, strings.Join(registerNames, ", "))
	}

	energy := registers[name]
	if energy == 0 {
		fmt.Printf("%s's light is not initialized (energy level: 0)\n", name)
	} else {
		fmt.Printf("%s's energy level: %d\n", name, energy)
	}
	return nil
}

func dim(name string) error {
	if _, exists := getRegisterIndex(name); !exists {
		return fmt.Errorf("unknown firefly: %s. Use predefined names: %s", name, strings.Join(registerNames, ", "))
	}

	if registers[name] == 0 {
		return fmt.Errorf("%s's energy is already at 0", name)
	}

	registers[name] = 0
	fmt.Printf("%s's energy has been reset to 0\n", name)
	return nil
}

func getRegisterIndex(name string) (int, bool) {
	for i, regName := range registerNames {
		if regName == name {
			return i, true
		}
	}
	return -1, false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// clearTerminal clears the terminal screen
func clearTerminal() error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	return cmd.Run()
}
