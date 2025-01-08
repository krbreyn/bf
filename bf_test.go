package bf_test

import (
	"testing"

	"github.com/krbreyn/bf"
)

var helloWorldProgram string = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

func TestInterpreter(t *testing.T) {
	t.Run("hello world program", func(t *testing.T) {
		got := bf.RunProgram(helloWorldProgram)

		if got != "Hello World!\n" {
			t.Errorf("wanted 'Hello World!\\n', got %s", got)
		}
	})

	// TODO test loops and other stuff individually i guess
}

func TestDialect(t *testing.T) {
	t.Run("compresses hello world program", func(t *testing.T) {
		got := bf.CompressProgram(helloWorldProgram)
		want := "+8[>+4[>+2>+3>+3>+<4-]>+>+>->2+[<]<-]>2.>-3.+7.2+3.>2.<-.<.+3.-6.-8.>2+.>+2."

		if got != want {
			t.Errorf("got %s ,\nwanted %s", got, want)
		}
	})

	t.Run("decompresses hello world program", func(t *testing.T) {

	})
}
