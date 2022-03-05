package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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

var (
	historyFileName = fmt.Sprintf("%s-history.txt", appName)
	historyFile     = filepath.Join(os.TempDir(), historyFileName)
)

func main() {
	os.Exit(int(Main(os.Args[1:])))
}

func Main(args []string) exitCode {
	var opts options
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = appName
	parser.Usage = appUsage
	args, err := parser.ParseArgs(args)
	if err != nil {
		if flags.WroteHelp(err) {
			return exitCodeOK
		} else {
			log.Println("Argument parsing failed.")
			return exitCodeErr
		}
	}

	if opts.Version {
		fmt.Printf("%s v%s\n", appName, appVersion)
		return exitCodeOK
	}

	if len(args) >= 1 {
		log.Println("Too many arguments.")
		return exitCodeErr
	}

	e := repl()
	return e
}

func repl() exitCode {
	line := liner.NewLiner()
	defer line.Close()

	// Set liner option
	line.SetCtrlCAborts(true)
	line.SetMultiLineMode(true)

	// Load history file
	if f, err := os.Open(historyFile); err == nil {
		line.ReadHistory(f)
	}

	fmt.Printf("%s v%s\nType `help` to show help message. Type `exit` to quit.\n\n", appName, appVersion)

	for {
		code, err := line.Prompt("websh# ")
		if err != nil {
			log.Println(err)
			return exitCodeErr
		}

		if code == "" {
			continue
		}

		if code == "exit" {
			break
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

	// Write history into file
	if f, err := os.Create(historyFile); err != nil {
		log.Println("Error writiing history file:", err)
	} else {
		line.WriteHistory(f)
	}

	return exitCodeOK
}
