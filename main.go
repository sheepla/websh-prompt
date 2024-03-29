package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
	"github.com/mattn/go-colorable"
	"github.com/peterh/liner"
	"github.com/sheepla/websh-prompt/client"
)

const (
	appName    = "websh-prompt"
	appUsage   = "[OPTIONS]"
	appVersion = "0.0.5"
)

const helpMessage = `
COMMANDS
  exit    Quit interactive UI
  help    Show help message
  version Show version
  ping    Test websh server status

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
	exitCodeErrArgs
	exitCodeErrPrompt
	exitCodeErrPing
	exitCodeErrPost
)

type options struct {
	Version bool `short:"V" long:"version" description:"Show version"`
	Test    bool `short:"t" long:"test" description:"Test websh server status"`
}

var (
	stdout = colorable.NewColorableStdout()
	stderr = colorable.NewColorableStderr()
)

var (
	historyFileName = fmt.Sprintf("%s-history.txt", appName)
	historyFile     = filepath.Join(os.TempDir(), historyFileName)
)

var commands = []string{"exit", "help", "version", "ping"}

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
		}
		log.Println("Argument parsing failed.")
		return exitCodeErrArgs
	}

	if opts.Version {
		fmt.Fprintf(stdout, "%s v%s\n", appName, appVersion)

		return exitCodeOK
	}

	if opts.Test {
		result, err := ping()
		if err != nil {
			log.Println(err)
			return exitCodeErrPing
		}

		fmt.Println(result.Status)

		return exitCodeOK
	}

	if len(args) >= 1 {
		log.Println("Too many arguments.")
		return exitCodeErrArgs
	}

	return repl()
}

//nolint:funlen
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
		return []string{}
	})

	// Load history file
	if f, err := os.Open(historyFile); err == nil {
		if _, err := line.ReadHistory(f); err != nil {
			fmt.Fprintln(stderr, err)
		}
	}

	fmt.Printf("%s v%s\nType `help` to show help message. Type `exit` to quit.\n\n", appName, appVersion)

REPL:
	for {
		code, err := line.Prompt("# ")
		if err != nil {
			if errors.Is(err, io.EOF) {
				return exitCodeOK
			}

			log.Println(err)
			return exitCodeErrPrompt
		}

		switch code {
		case "exit":
			break REPL
		case "help":
			fmt.Fprint(stdout, helpMessage)

			continue
		case "version":
			fmt.Fprintf(stdout, "%s v%s\n", appName, appVersion)

			continue
		case "":

			continue
		case "ping":
			result, err := ping()
			if err != nil {
				fmt.Fprintln(stderr, color.HiRedString(err.Error()))
			}
			if result.Status != "" {
				fmt.Fprintln(stdout, result.Status)
			}

			continue
		default:
			result, err := run(code)
			if err != nil {
				fmt.Fprintln(stderr, err)

				continue
			}

			if result.Stdout != "" {
				fmt.Fprintln(stdout, result.Stdout)
			}
			if result.Stderr != "" {
				fmt.Fprintln(stderr, color.HiRedString(result.Stderr))
			}

			printPrompt(result)
			line.AppendHistory(code)

			continue
		}

	}

	// Write history into file
	if f, err := os.Create(historyFile); err != nil {
		log.Println(err)
	} else {
		line.WriteHistory(f)
	}

	return exitCodeOK
}

func run(code string) (*client.Result, error) {
	return client.Post(client.Param{
		Code: code,
	})
}

func ping() (*client.PingResult, error) {
	return client.Ping()
}

func printPrompt(result *client.Result) {
	fmt.Fprintf(stdout,
		"%s %s\n",
		color.New(color.FgBlue, color.Bold).Sprint("websh"),
		color.New(color.FgCyan, color.Bold).Sprint(result.ElapsedTime),
	)
}
