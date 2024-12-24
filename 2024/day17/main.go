package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println("This is main for 2024/day02")
	filename := "./data_test.in"
	filename = "./data.in"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	computer := Computer{
		InstructionPointer: 0,
	}
	m := regexp.MustCompile(`\d+`)
	scanner.Scan()
	value := m.FindAllString(scanner.Text(), -1)[0]
	integer, _ := strconv.Atoi(value)
	computer.A = integer
	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	programRegex := regexp.MustCompile(`Program: (.*)`)
	computer.Program = strings.Split(programRegex.FindAllStringSubmatch(scanner.Text(), -1)[0][1], ",")
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result exercise 1: ", exercise1(computer))
	fmt.Println("Result exercise 2: ", exercise2(computer))
}

type Computer struct {
	InstructionPointer int
	Program            []string
	A                  int
	B                  int
	C                  int
	Output             []string
	Clock              int
}

func (c *Computer) GetOutput() string {
	return strings.Join(c.Output, ",")
}

func (c *Computer) isHalted() bool {
	return c.InstructionPointer >= len(c.Program)
}

var OPCODES map[int]string = map[int]string{
	0: "A = A / COMBO<<2",
	1: "B = B ^ LITERAL",
	2: "B = COMBO % 8",
	3: "jmp LITERAL if A==0",
	4: "B = B ^ C",
	5: "out(COMBO)",
	6: "B = A / COMBO<<2",
	7: "C = A / COMBO<<2",
}

func (c *Computer) parseOpCodes() {
	ip := 0
	for ip < len(c.Program) {
		opCode, _ := strconv.Atoi(c.Program[ip])
		literalOperand, _ := strconv.Atoi(c.Program[ip+1])

		str := OPCODES[opCode]
		comboOperand := literalOperand
		if comboOperand == 4 {
			comboOperand = c.A
		} else if comboOperand == 5 {
			comboOperand = c.B
		} else if comboOperand == 6 {
			comboOperand = c.C
		}
		fmt.Println(opCode, literalOperand, "|", str, "COMBO=", comboOperand, "LITERAL=", literalOperand)
		ip += 2
	}
}

func (c *Computer) nextInstruction(debug bool) {
	opCode, _ := strconv.Atoi(c.Program[c.InstructionPointer])
	literalOperand, _ := strconv.Atoi(c.Program[c.InstructionPointer+1])

	if debug {
		fmt.Printf("- OpCode: %d %d\n", opCode, literalOperand)
	}

	comboOperand := literalOperand
	if comboOperand == 4 {
		comboOperand = c.A
	} else if comboOperand == 5 {
		comboOperand = c.B
	} else if comboOperand == 6 {
		comboOperand = c.C
	}

	result := 0.0
	switch opCode {
	case 0:
		result = float64(c.A) / math.Pow(2, float64(comboOperand))
		c.A = int(result)
	case 1:
		c.B = c.B ^ literalOperand
	case 2:
		c.B = comboOperand % 8
	case 3:
		if c.A != 0 {
			c.InstructionPointer = literalOperand - 2
		}
	case 4:
		c.B = c.B ^ c.C
	case 5:
		c.Output = append(c.Output, strconv.Itoa(comboOperand%8))
	case 6:
		result = float64(c.A) / math.Pow(2, float64(comboOperand))
		c.B = int(result)
	case 7:
		result = float64(c.A) / math.Pow(2, float64(comboOperand))
		c.C = int(result)
	default:
		os.Exit(4)
	}

	c.InstructionPointer += 2
	c.Clock++
}

func (c *Computer) PrintState() {
	fmt.Printf("State:\tRegister A: %d\n\tRegister B: %d\n\tRegister C: %d\n\tProgram: %s\n\tIP:%d\n\tClock: %d\n\tOutput: %s\n\n", c.A, c.B, c.C, strings.Join(c.Program, " "), c.InstructionPointer, c.Clock, c.GetOutput())
}

func exercise1(c Computer) string {
	// c.PrintState()
	for !c.isHalted() {
		c.nextInstruction(false)
	}
	return strings.Join(c.Output, ",")
}

// hardcoded solution for my input
func solve(program []string, aValue int) int {
	fmt.Println(program, aValue)
	if len(program) == 0 {
		return aValue
	}
	for k := range 8 {
		// in each loop the R_A value is divided by 2**3
		// start at the end, and increase with leftShift
		// or with
		a := (aValue << 3) | k
		println(a)
		b := a % 8
		b = b ^ 7
		c := a >> b
		b = b ^ c
		b = b ^ 4
		if strconv.Itoa(b%8) == program[len(program)-1] {
			sub := solve(program[:len(program)-1], a)
			if sub == -1 {
				continue
			}
			return sub
		}
	}
	return -1
}

func exercise2(comp Computer) int {
	comp.parseOpCodes()
	// Outputs
	/*
	   2 4 | B = COMBO % 8 COMBO= 53437164 LITERAL= 4
	   1 7 | B = B ^ LITERAL COMBO= 7 LITERAL= 7
	   7 5 | C = A / COMBO<<2 COMBO= 0 LITERAL= 5
	   4 1 | B = B ^ C COMBO= 1 LITERAL= 1
	   1 4 | B = B ^ LITERAL COMBO= 53437164 LITERAL= 4
	   5 5 | out(COMBO) COMBO= 0 LITERAL= 5
	   0 3 | A = A / COMBO<<2 COMBO= 3 LITERAL= 3
	   3 0 | jmp LITERAL if A==0 COMBO= 0 LITERAL= 0
	*/

	// Simplifying instructions:
	// B = A % 8
	// B = B ^ 7
	// C = A >> B
	// B = B ^ C
	// B = B ^ 4
	// out(B % 8)
	// A = A >> 3
	// jmp 0 if A!=0 COMBO= 0 LITERAL= 0

	return solve(comp.Program, 1)
}
