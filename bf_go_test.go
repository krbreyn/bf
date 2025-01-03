package bf_go_test

import (
	"testing"

	"github.com/krbreyn/bf_go"
)

func TestInterpreter(t *testing.T) {
	t.Run("hello world program", func(t *testing.T) {
		program := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
		output := bf_go.RunProgram(program)

		if output != "Hello World!\n" {
			t.Errorf("wanted 'Hello World!\\n', got %s", output)
		}
	})
}
