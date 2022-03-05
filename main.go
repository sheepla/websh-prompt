package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/mattn/go-colorable"
	"github.com/peterh/liner"
	"github.com/sheepla/websh-prompt/client"
)

const (
	appName    = "websh-prompt"
	appUsage   = "[OPTIONS]"
	appVersion = "0.0.1"
)

type exitCode int

const (
	exitCodeOK exitCode = iota
	exitCodeErr
)

type options struct {
	Version bool `short:"V" long:"version" description:"Show version"`
}

var (
	stdout = colorable.NewColorableStdout()
	stderr = colorable.NewColorableStderr()
)

func main() {
	os.Exit(int(Main(os.Args[1:])))
}

func Main(args []string) exitCode {
	var opts options
	parser := flags.NewParser(&opts, flags.Default)
	args, err := parser.ParseArgs(args)
	if err != nil {
		if flags.WroteHelp(err) {
			return exitCodeOK
		} else {
			fmt.Fprintln(os.Stderr, "Argument parsing failed.")
			return exitCodeErr
		}
	}

	if len(args) >= 1 {
		fmt.Fprintln(os.Stderr, "Too many arguments.")
		return exitCodeErr
	}

	parser.Name = appName
	parser.Usage = appUsage

	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	for {
		code, err := line.Prompt("websh# ")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return exitCodeErr
		}
		if code == "" {
			continue
		}

		if code == "exit" {
			return exitCodeOK
		}

		p := client.Param{
			Code: code,
		}
		result, err := client.Post(p)
		if err != nil {
			continue
		}
		if result.Stdout != "" {
			fmt.Fprintln(stdout, result.Stdout)
		}
		if result.Stderr != "" {
			fmt.Fprintln(stderr, result.Stderr)
		}

		line.AppendHistory(code)
	}
}
