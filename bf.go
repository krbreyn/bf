package bf

import (
	"bufio"
	"fmt"
	"os"
)

func RunProgram(program string) string {
	var tape [30000]byte
	var pointer uint

	var bracketCounter int

	var output []rune

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(program); i++ {
		switch program[i] {

		// no bounds checking, that is undefined behavior

		case '>':
			pointer++

		case '<':
			pointer--

		case '+':
			tape[pointer]++

		case '-':
			tape[pointer]--

		case '.':
			output = append(output, rune(tape[pointer]))

		case ',':
			//cheating
			fmt.Print(string(output))
			output = []rune{}

			in, _, err := reader.ReadRune()
			if err != nil {
				panic(err)
			}
			tape[pointer] = byte(in)

		case '[':
			if tape[pointer] == 0 {
				bracketCounter++
				for program[i] != ']' || bracketCounter != 0 {
					i++
					switch program[i] {
					case '[':
						bracketCounter++
					case ']':
						bracketCounter--
					}
				}
			}

		case ']':
			if tape[pointer] != 0 {
				bracketCounter++
				for program[i] != '[' || bracketCounter != 0 {
					i--
					switch program[i] {
					case ']':
						bracketCounter++
					case '[':
						bracketCounter--
					}
				}
			}

		}
	}

	return string(output)
}
