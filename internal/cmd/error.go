package cmd

import (
	"fmt"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		msg := fmt.Sprintf("Fatal error: %+v", err)
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		fmt.Fprintf(os.Stderr, msg)
		//revive:disable-next-line
		os.Exit(1)
	}
}
