# anki-cli

Simple CLI tool to add "word - translation" pairs to Anki decks using AnkiConnect.

## Features

- Add single or multiple words in the format `word - translation`
- Supports custom deck names and note types

## Requirements

- [Anki](https://apps.ankiweb.net/) with [AnkiConnect](https://ankiweb.net/shared/info/2055492159) plugin installed and running
- Go 1.18+

## Installation

### 1 option:

Clone this repository:

```bash
git clone https://github.com/ialeksiienko/anki-cli.git
cd anki-cli
```

Build the CLI:

```bash
go build -o anki-cli main.go
```

### 2 option:

Download a binary from [Releases](https://github.com/ialeksiienko/anki-cli/releases)

### 3 option:

Go install:

```bash
go install github.com/ialeksiienko/anki-cli@latest
```

## Usage

Add a single word:

```bash
./anki-cli "deck-name" "model-name" "word - translation"
```

Add multiple words (newline-separated):

```bash
./anki-cli "deck-name" "model-name" "word - translation
word1 - translation1
word2 - translation2"
```

## How it works

The CLI sends a JSON request to AnkiConnect at http://localhost:8765.
It uses the "Front"/"Back" fields by default.
