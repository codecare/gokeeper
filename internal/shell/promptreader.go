package shell

import (
	"bufio"
	"os"
)

func ReadInput() string {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return scanner.Text()
	}

	return ""
}
