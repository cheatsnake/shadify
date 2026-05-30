## 🦖 This repository is not maintained and exists only for history

# Shadify

<a href="https://goreportcard.com/report/github.com/cheatsnake/shadify"><img src="https://goreportcard.com/badge/github.com/cheatsnake/shadify" alt="Go report"/></a>
<a href="https://github.com/cheatsnake/shadify/releases"><img src="https://img.shields.io/github/v/release/cheatsnake/shadify.svg" alt="GitHub repo size"/></a>
<a href="https://img.shields.io/github/repo-size/cheatsnake/shadify?color=blue"><img src="https://img.shields.io/github/repo-size/cheatsnake/shadify?color=blue" alt="GitHub repo size"/></a>
<a href="https://img.shields.io/github/license/cheatsnake/shadify?color=orange"><img src="https://img.shields.io/github/license/cheatsnake/shadify?color=orange" alt="GitHub repo license"/></a>
<a href="https://github.com/cheatsnake/shadify"><img src="https://img.shields.io/badge/status-archived-orange.svg?style=flat" alt="GitHub repo issues"/></a>

<p align="center"><img src="/docs/images/shadify.png" alt="Promo"/></p>

Shadify is a powerful REST API service provides a collection of different puzzle types, like crosswords, Sudoku, word search and so on. The API allows users to generate data for puzzles, check the correctness of solutions, and configure various parameters to change the difficulty of the puzzles.

## 📃 Documentation

All documentation is available [here](https://github.com/cheatsnake/shadify/blob/master/docs/documentation.md).

The service is divided into independent modules. Each module starts with a brief description of what the module is oriented at (be it a game, a puzzle, a task, etc.). This is followed by a detailed description of each HTTP interface, with descriptions of the possible parameters and return responses.

## 🚀 Server startup

1. Clone this repository onto your computer:

```sh
git clone https://github.com/cheatsnake/shadify.git
```

2. Inside the project, run this command to install the necessary packages:

```sh
go mod download
```

> Make sure you have already [installed Go](https://go.dev) on your computer.

3. Start the server by running the last command:

```sh
go run cmd/server/main.go
```

> The server will start at the address http://localhost:5000

For easy testing, use a ready-made collection for [Insomnia](https://insomnia.rest). Open `Settings` > `Data` > `Import Data` > `From URL` and paste link to the [insomnia.shadify.json](./insomnia.shadify.json) file. Enjoy!

## 🐳 Docker container startup

-   Running a Docker container for development (after any file changes, server will be restarts):

```sh
docker compose up
```

-   Running a Docker container for production:

```sh
docker build -t shadify . --target prod
```

```sh
docker run --rm -p 5000:5000 --name shadify shadify
```

<div align="center">Made with &#9829;</div>
<div align="center"><a href="https://github.com/cheatsnake/shadify/blob/master/LICENSE">LICENSE</a> 2022-2023</div>
