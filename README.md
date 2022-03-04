# 💻 websh-prompt

A command line [websh](https://github.com/jiro4989/websh) client with bash-like interactive UI

## Features

- Run commands on the websh environment without huge Docker image
- Supports bash / GNU Readline like keybindings, killrings, console clearing


## Usage

help message:

```
$ websh-prompt
Usage:
  websh-prompt [OPTIONS]

Application Options:
  -V, --version  Show version

Help Options:
  -h, --help     Show this help message
```

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


## Installation

Clone this repository then run `go install`.

## Thanks

[jiro4989/websh](https://github.com/jiro4989/websh) 

## LICENSE

MIT
