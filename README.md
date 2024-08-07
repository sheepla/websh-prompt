<div align="right">

![CI](https://github.com/sheepla/websh-prompt/actions/workflows/ci.yml/badge.svg)
![Relase](https://github.com/sheepla/websh-prompt/actions/workflows/release.yml/badge.svg)

<a href="https://github.com/sheepla/websh-prompt/releases/latest">

![Latest Release](https://img.shields.io/github/v/release/sheepla/websh-prompt?style=flat-square)

</a>

</div>


<div align="center">

# 💻 websh-prompt (Archived)

*The **websh** service is no longer available. Thank you very much for the operation so far. With the end of service, this repository is archived as well.*

</div>


<div align="center">

A command line [websh](https://github.com/jiro4989/websh) client with bash-like interactive UI

![](https://img.shields.io/static/v1?label=Language&message=Go&color=blue&style=flat-square)
![](https://img.shields.io/static/v1?label=License&message=MIT&color=blue&style=flat-square)

</div>


<div align="center">
<img src="https://user-images.githubusercontent.com/62412884/159151159-d3b18f14-f714-43fc-a43d-0e81bbf9d9f8.png" width="70%">
</div>


## Features

- Run commands on the websh environment without huge Docker image
- Supports bash / GNU Readline like line editing, console clearing (`Ctrl-L`), history searching (`Ctrl-R`), and `Tab` completion

## Usage

### Help message

```
Usage:
  websh-prompt [OPTIONS]

Application Options:
  -V, --version  Show version
  -t, --test     Test websh server status

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

### Built in commands

- `help`: Show help message.
- `exit`: Quit interactive UI
- `version`: Show version
- `ping`: Test websh server status

## Installation

### Download executable binary

You can download executable binary from release page

> [Latest Release](https://github.com/sheepla/websh-prompt/releases/latest)

### Use GitHub release installer tools

These tools make it easy to install executable binaries from GitHub Release.

with [ghg](https://github.com/Songmu/ghg):

```bash
ghg get sheepla/websh-prompt  # Install
ls -l $(ghg bin)/websh-prompt # It will exists executable
```

with [relma](https://github.com/jiro4989/relma):

Copy download link URL from [Latest Release](https://github.com/sheepla/websh-prompt/releases/latest) page, then run below.


```bash
relma init                           # Setup
relma install {{DOWNLOAD_LINK_URL}}  # Install
ls -l $(relma root)/bin/websh-prompt # It will exists executable
```

with [gh-install](https://github.com/redraw/gh-install)

```bash
gh install sheepla/websh-prompt # Install
ls -l ~/.local/bin/websh-prompt # It will exists executable
```

### Build from source

Clone this repository then run `go install`.
Developing on `v1.17.7 linux/amd64`.

## Special Thanks

[jiro4989/websh](https://github.com/jiro4989/websh) 

## Contributing

Welcome 💖

## LICENSE

[MIT](./LICENSE)

## Author

[Sheepla](https://github.com/sheepla)
