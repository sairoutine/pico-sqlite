package main
import (
	"os"
	"fmt"
);

func main() {
	input_buffer := new_input_buffer();

	for ;; {
		print_prompt();
		read_input(input_buffer);

		if input_buffer.buffer == ".exit" {
			os.Exit(0);
		} else {
			fmt.Printf("Unrecognized command '%s'.\n", input_buffer.buffer);
		}
	}
}


func new_input_buffer() int {

	return 0;
}

func print_prompt() {
}
func read_input(input_buffer int) {
}


