package bf

import (
	"bufio"
	"os"
)

// TODO create a compressed notation dialect that reduces the amount of command repetition
// ex. ++++++++ could be expressed as +8
// in order to avoid conflicts with normal programs that use comments, it should be done by feeding it into a function that outputs the expanded code

func RunProgram(program string) string {
	var tape [30000]byte
	var pointer uint

	var bracketCounter int

	var output []rune

	// alternative way of doing it is to send input in as string alongside program
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
