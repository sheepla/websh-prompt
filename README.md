# 💻 websh-prompt

A command line [websh](https://github.com/jiro4989/websh) client with bash-like interactive UI

## Features

- Run commands on the websh environment without huge Docker image
- Supports bash / GNU Readline like editing, console clearing (`Ctrl-L`), history searching (`Ctrl-R`)

## Usage

### Help message

```
$ websh-prompt
Usage:
  websh-prompt [OPTIONS]

Application Options:
  -V, --version  Show version

Help Options:
  -h, --help     Show this help message
```

### How to use

Just run the `websh-prompt` then you'll be promped. Enter the command you want, the result will be desplayed.
To quit interactive mode, type `exit` or `Ctrl-D`.

```
[you@your-computer]$ websh-prompt
websh# uname -a
Linux 9da3a1b93c29 4.15.0-55-generic #60-Ubuntu SMP Tue Jul 2 18:22:20 UTC 2019 x86_64 x86_64 x86_64 GNU/Linux

websh# figlet -f slant websh
                __         __  
 _      _____  / /_  _____/ /_ 
| | /| / / _ \/ __ \/ ___/ __ \
| |/ |/ /  __/ /_/ (__  ) / / /
|__/|__/\___/_.___/____/_/ /_/ 
                               
websh# exit
[you@your-computer]$ 
```

### Line Editing

To see the line editing key bindings, see the README on [peterh/liner](https://github.com/peterh/liner#readme).

### History retension

The command line execution history is saved in `websh-prompt-history.txt` under the temporary directory of your system. When you start the program, the history is loaded and you can edit the history and re-execute.

## Installation

### Download executable binary

You can download executable binary from release page

> [Latest Release](https://github.com/sheepla/websh-prompt/releases/latest)

### Build from source

Clone this repository then run `go install`.

## Special Thanks

[jiro4989/websh](https://github.com/jiro4989/websh) 

## LICENSE

MIT
