package main
import (
	"os"
	"fmt"
);

type InputBuffer struct {
	Buffer string
	BufferLength int
	InputLength int
}


func main() {
	input_buffer := new_input_buffer();

	for ;; {
		print_prompt();
		read_input(input_buffer);

		if input_buffer.Buffer == ".exit" {
			os.Exit(0);
		} else {
			fmt.Printf("Unrecognized command '%s'.\n", input_buffer.Buffer);
		}
	}
}


func new_input_buffer() InputBuffer {
	input_buffer := InputBuffer{};
	input_buffer.Buffer = "";
	input_buffer.BufferLength = 0;
	input_buffer.InputLength = 0;

	return input_buffer;
}

func print_prompt() {
}
func read_input(input_buffer InputBuffer) {
}


