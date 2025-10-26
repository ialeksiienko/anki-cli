# anki-go

Simple CLI tool to add "word - translation" pairs to Anki decks using AnkiConnect.

## Features

- Add single or multiple words in the format `word - translation`
- Automatically creates the deck if it doesn't exist
- Supports custom deck names and the default "Basic" note type

## Requirements

- [Anki](https://apps.ankiweb.net/) with [AnkiConnect](https://ankiweb.net/shared/info/2055492159) plugin installed and running
- Go 1.18+

## Installation

Clone this repository:

```bash
git clone https://github.com/yourusername/anki-go.git
cd anki-go
```

## Build the CLI:

```bash
go build -o anki-go main.go
```

## Usage

Add a single word:

```bash
./anki-go "word - translation"
```

Add multiple words (newline-separated):

```bash
./anki-go "word - translation
word1 - translation1
word2 - translation2"
```

## How it works

The CLI sends a JSON request to AnkiConnect at http://localhost:8765.
It uses the "Basic" note type with "Front"/"Back" fields by default.
