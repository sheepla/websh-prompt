package main

import (
	"fmt"
	"log"
	"os"

	"github.com/peterh/liner"
	"github.com/sheepla/websh-prompt/client"
	// "github.com/sheepla/websh-prompt/client"
)

func main() {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	for {
		code, err := line.Prompt("websh# ")
		if err != nil {
			log.Fatal(err)
		}
		if code == "" {
			continue
		}

        if code == "exit" {
            return
        }

		p := &client.Param{
			Code: code,
		}
		result, err := client.Post(*p)
		if err != nil {
			continue
		}
		if result.Stdout != "" {
			fmt.Fprintln(os.Stdout, result.Stdout)
		}
		if result.Stderr != "" {
			fmt.Fprintln(os.Stderr, result.Stderr)
		}
	}
}
