package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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

const helpMessage = `
COMMANDS
    exit    Quit interactive UI
    help    Show help message
    version Show version

KEY BINDINGS
  Ctrl-A, Home          Move cursor to beginning of line
  Ctrl-E, End           Move cursor to end of line
  Ctrl-B, Left          Move cursor one character left
  Ctrl-F, Right         Move cursor one character right
  Ctrl-Left, Alt-B      Move cursor to previous word
  Ctrl-Right, Alt-F     Move cursor to next word
  Ctrl-D, Del           (if line is not empty) Delete character under cursor
  Ctrl-D                (if line is empty) End of File - usually quits application
  Ctrl-C                Reset input (create new empty prompt)
  Ctrl-L                Clear screen (line is unmodified)
  Ctrl-T                Transpose previous character with current character
  Ctrl-H, BackSpace     Delete character before cursor
  Ctrl-W, Alt-BackSpace Delete word leading up to cursor
  Alt-D                 Delete word following cursor
  Ctrl-K                Delete from cursor to end of line
  Ctrl-U                Delete from start of line to cursor
  Ctrl-P, Up            Previous match from history
  Ctrl-N, Down          Next match from history
  Ctrl-R                Reverse Search history (Ctrl-S forward, Ctrl-G cancel)
  Ctrl-Y                Paste from Yank buffer (Alt-Y to paste next yank instead)
  Tab                   Next completion
  Shift-Tab             (after Tab) Previous completion
`

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

var commands = []string{"exit", "help", "version"}

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
		version()
		return exitCodeOK
	}

	if len(args) >= 1 {
		log.Println("Too many arguments.")
		return exitCodeErr
	}

	e := repl()
	return e
}

func version() {
	fmt.Printf("%s v%s\n", appName, appVersion)
}

func repl() exitCode {
	line := liner.NewLiner()
	defer line.Close()

	// Set liner option
	line.SetCtrlCAborts(true)
	line.SetMultiLineMode(true)

	// Set tab completion
	line.SetCompleter(func(line string) (c []string) {
		for _, n := range commands {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

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

		if code == "help" {
			fmt.Print(helpMessage)
			continue
		}

		if code == "exit" {
			break
		}

		if code == "version" {
			version()
			continue
		}

		p := client.Param{
			Code: code,
		}
		result, err := client.Post(p)
		if err != nil {
			log.Println(err)
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
