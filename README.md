[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/PHq8Kfj_)

# Go Installation Guide

This guide provides steps to install Go (Golang) on different operating systems.

## Prerequisites

- Access to a terminal or command prompt
- Stable internet connection

## Installation Steps

### Windows

1. Visit the official Go website: [https://golang.org/dl/](https://golang.org/dl/)
2. Download the Windows installer.
3. Run the installer and follow the installation instructions.
4. Set up environment variables:
- Open Control Panel > System and Security > System.
- Click on "Advanced system settings" and go to the "Advanced" tab.
- Click on "Environment Variables".
- Add a new system variable `GOROOT` pointing to the Go installation directory and add `%GOROOT%\bin` to the `PATH` variable.
5. Open a command prompt and type `go version` to verify the installation.

That's it! You're now ready to start coding in Go.

# Using golangci-lint for Go Code Linting

`golangci-lint` is a powerful Go linter that helps in identifying issues, bugs, and style problems in your Go codebase. Here's how to set it up:

## Installation

### Prerequisites

- Ensure you have Go installed on your system. If not, refer to the Go installation guide provided in this repository.

### Installing golangci-lint

#### Using `go get`

1. Open a terminal or command prompt.
2. Run the following command to install `golangci-lint`:

   ```bash
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

### Installation check :

    ```bash
        golangci-lint --version

### Run code verification :
    ```bash
    golangci-lint run

## Configuration with `.golangci.yml

The `.golangci.yml` file is a configuration file used by `golangci-lint` to adapt linting rules to the specific needs of your project.
