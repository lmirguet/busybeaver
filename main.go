package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PRINT_ALL_STATES = false

// Transition represents a single transition in the Turing machine.
type Transition struct {
	Read      int
	Write     int
	Move      string
	NextState string
}

// TuringMachineProgram represents the entire Turing machine program.
type TuringMachineProgram map[string]map[int]Transition

// ParseFile parses the input file and returns the Turing machine program.
func ParseFile(filePath string) (TuringMachineProgram, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	program := make(TuringMachineProgram)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ";")
		if len(parts) != 5 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		state := parts[0]
		read, _ := strconv.Atoi(parts[1])
		write, _ := strconv.Atoi(parts[2])
		move := parts[3]
		nextState := parts[4]

		if _, exists := program[state]; !exists {
			program[state] = make(map[int]Transition)
		}

		program[state][read] = Transition{
			Read:      read,
			Write:     write,
			Move:      move,
			NextState: nextState,
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return program, nil
}

// Execute one instruction then return the new currentNode and currentState
func ExecuteInstruction(program TuringMachineProgram, tape *DoublyLinkedList, currentNode *Node, currentState string) (*Node, string) {
	transition := program[currentState][currentNode.value]
	currentNode.value = transition.Write
	if transition.Move == "RIGHT" {
		currentNode = currentNode.next
		if currentNode == nil {
			tape.AppendRight(0)
			currentNode = tape.tail
		}
	}
	if transition.Move == "LEFT" {
		currentNode = currentNode.prev
		if currentNode == nil {
			tape.AppendLeft(0)
			currentNode = tape.head
		}
	}
	currentState = transition.NextState
	return currentNode, currentState
}

// Execute the Program... be careful, it may not return
func ExecuteProgram(program TuringMachineProgram, tape *DoublyLinkedList, currentNode *Node, currentState string) {
	i := 0
	for {
		i = i + 1
		currentNode, currentState = ExecuteInstruction(program, tape, currentNode, currentState)
		if PRINT_ALL_STATES {
			fmt.Print(i, " : ", currentState, " : ")
			tape.Display(60)
		}
		if currentState == "STOP" {
			if !PRINT_ALL_STATES {
				fmt.Print("Finished after ", i, " steps")
			}
			return
		}
	}
}

func main() {
	program, err := ParseFile("sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// create a tape as a doubly linked list and create a first node 0
	tape := &DoublyLinkedList{}
	tape.AppendRight(0)
	currentNode := tape.head

	// initial state is "A"
	currentState := "A"

	ExecuteProgram(program, tape, currentNode, currentState)
}
