package main
import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
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


func new_input_buffer() *InputBuffer {
	input_buffer := InputBuffer{};
	input_buffer.Buffer = "";
	input_buffer.BufferLength = 0;
	input_buffer.InputLength = 0;

	return &input_buffer;
}

func print_prompt() {
	fmt.Printf("db > ");
}

func read_input(input_buffer *InputBuffer) {
	r := bufio.NewReader(os.Stdin);
	line, err := r.ReadString('\n')

	if err == io.EOF {
		return;
	} else if err != nil {
		fmt.Printf("Error reading input\n");
		os.Exit(1);
	}

	// Ignore trailing newline
	input_buffer.Buffer = strings.TrimRight(line, "\n");
	input_buffer.InputLength = len(input_buffer.Buffer);
}


