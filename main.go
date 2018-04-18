package main
import (
	"os"
	"fmt"
	"bufio"
	"io"
);

type InputBuffer struct {
	Buffer []byte
	BufferLength int
	InputLength int
}
// enum
type MetaCommandResult int;
const (
    META_COMMAND_SUCCESS MetaCommandResult = iota
    META_COMMAND_UNRECOGNIZED_COMMAND
)

type PrepareResult int;
const (
    PREPARE_SUCCESS PrepareResult = iota
    PREPARE_UNRECOGNIZED_STATEMENT
)

type Statement int;
const (
    STATEMENT_INSERT Statement = iota
    STATEMENT_SELECT
)

func main() {
	input_buffer := new_input_buffer();

	for ;; {
		print_prompt();
		read_input(input_buffer);

		if string(input_buffer.Buffer[0]) == "." {
			switch (do_meta_command(input_buffer)) {
				case (META_COMMAND_SUCCESS):
					continue;
				case (META_COMMAND_UNRECOGNIZED_COMMAND):
					fmt.Printf("Unrecognized command '%s'\n", input_buffer.Buffer);
			}
		}

		var statement Statement;
		switch (prepare_statement(input_buffer, &statement)) {
			case (PREPARE_SUCCESS):
				break;
			case (PREPARE_UNRECOGNIZED_STATEMENT):
				fmt.Printf("Unrecognized keyword at start of '%s'\n", input_buffer.Buffer);
				continue;
		}

		execute_statement(&statement);
		fmt.Printf("Executed.\n");
	}
}


func new_input_buffer() *InputBuffer {
	input_buffer := InputBuffer{};
	input_buffer.Buffer = []byte{};
	input_buffer.BufferLength = 0;
	input_buffer.InputLength = 0;

	return &input_buffer;
}

func print_prompt() {
	fmt.Printf("db > ");
}

func read_input(input_buffer *InputBuffer) {
	r := bufio.NewReader(os.Stdin);
	line, err := r.ReadBytes('\n')

	if err == io.EOF {
		return;
	} else if err != nil {
		fmt.Printf("Error reading input\n");
		os.Exit(1);
	}

	// Ignore trailing newline
	line_length := len(line) - 1;
	input_buffer.Buffer = line[0:line_length];
	input_buffer.InputLength = line_length;
}

func do_meta_command(input_buffer *InputBuffer) MetaCommandResult {
	return META_COMMAND_SUCCESS;
}
func prepare_statement(input_buffer *InputBuffer, statement *Statement) PrepareResult {
	return PREPARE_SUCCESS;
}
func execute_statement(statement *Statement) {
}

