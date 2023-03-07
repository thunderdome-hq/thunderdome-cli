```shell


▄▄▄█████▓ ██░ ██  █    ██  ███▄    █ ▓█████▄ ▓█████  ██▀███  ▓█████▄  ▒█████   ███▄ ▄███▓▓█████ 
▓  ██▒ ▓▒▓██░ ██▒ ██  ▓██▒ ██ ▀█   █ ▒██▀ ██▌▓█   ▀ ▓██ ▒ ██▒▒██▀ ██▌▒██▒  ██▒▓██▒▀█▀ ██▒▓█   ▀ 
▒ ▓██░ ▒░▒██▀▀██░▓██  ▒██░▓██  ▀█ ██▒░██   █▌▒███   ▓██ ░▄█ ▒░██   █▌▒██░  ██▒▓██    ▓██░▒███   
░ ▓██▓ ░ ░▓█ ░██ ▓▓█  ░██░▓██▒  ▐▌██▒░▓█▄   ▌▒▓█  ▄ ▒██▀▀█▄  ░▓█▄   ▌▒██   ██░▒██    ▒██ ▒▓█  ▄ 
  ▒██▒ ░ ░▓█▒░██▓▒▒█████▓ ▒██░   ▓██░░▒████▓ ░▒████▒░██▓ ▒██▒░▒████▓ ░ ████▓▒░▒██▒   ░██▒░▒████▒
  ▒ ░░    ▒ ░░▒░▒░▒▓▒ ▒ ▒ ░ ▒░   ▒ ▒  ▒▒▓  ▒ ░░ ▒░ ░░ ▒▓ ░▒▓░ ▒▒▓  ▒ ░ ▒░▒░▒░ ░ ▒░   ░  ░░░ ▒░ ░
    ░     ▒ ░▒░ ░░░▒░ ░ ░ ░ ░░   ░ ▒░ ░ ▒  ▒  ░ ░  ░  ░▒ ░ ▒░ ░ ▒  ▒   ░ ▒ ▒░ ░  ░      ░ ░ ░  ░
  ░       ░  ░░ ░ ░░░ ░ ░    ░   ░ ░  ░ ░  ░    ░     ░░   ░  ░ ░  ░ ░ ░ ░ ▒  ░      ░      ░   
          ░  ░  ░   ░              ░    ░       ░  ░   ░        ░        ░ ░         ░      ░  ░
                                      ░                       ░                                 

```
![This is Thunderdome](https://media4.giphy.com/media/RFIuO4XWzU8gg/giphy.gif?cid=6104955evj86uj4odo4u1t2mcor2o8515jaqywsmzz1abzug&rid=giphy.gif&ct=g)

`Thunderdome` is a service that will allow people from white-listed email domains to help hummy.social with development.

`Thunderdome cli` is a command line interface for interacting with the thunderdome api. See below for detailed instructions!




## Pre-requisites

A go installation is required for installing the cli. 
If you do not have one, please follow the instructions on https://golang.org/doc/install.
Check if an installation exists using:

```shell 
go version
```

## Installation
Install the latest version of the CLI from GitHub using:

```shell
go install github.com/thunderdome-hq/thunderdome-cli/thunderdome@latest
```

## Usage

Run the following command in the terminal to display how to use the Thunderdome CLI. 

```shell
$ thunderdome --help
```

If the command was not found, the Go output folder must first be added to PATH.
This is done by adding the following line to the used configuration file (e.g. ~/.bashrc, ~/.zshrc, ~/.profile):
    
```shell
export PATH=$PATH:$(go env GOPATH)/bin
```
