package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// AnkiConnectURL is the local endpoint of the AnkiConnect API used to communicate with Anki
const AnkiConnectURL = "http://localhost:8765"

// Note represents a request body used to build JSON payloads for the AnkiConnect API
type Note struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  struct {
		Note struct {
			DeckName  string `json:"deckName"`
			ModelName string `json:"modelName"`
			Fields    struct {
				Front string `json:"Front"`
				Back  string `json:"Back"`
			} `json:"fields"`
			Tags []string `json:"tags"`
		} `json:"note"`
	} `json:"params"`
}

func main() {
	versionFlag := flag.Bool("version", false, "print version information")
	flag.Parse()

	if *versionFlag {
		printVersion()
		return
	}

	if len(os.Args) <= 3 {
		color.Red(fmt.Sprintf("usage: %s \"deck-name\" \"model-name\" \"word-translate\"", os.Args[0]))
		os.Exit(1)
	}

	deckname := os.Args[1]  // f. e. English
	modelname := os.Args[2] // f.e. Basic

	if deckname == "" && modelname == "" {
		printError("deck name and model name are required")
		os.Exit(1)
	}

	text := os.Args[3]
	lines := strings.Split(text, "\n")

	for _, l := range lines {
		words := strings.Split(l, "-")

		englishWord, translatedWord := strings.TrimSpace(words[0]), strings.TrimSpace(words[1])

		// build a Note struct for the current word pair
		var note Note
		note.Action = "addNote"
		note.Version = 6
		note.Params.Note.DeckName = deckname
		note.Params.Note.ModelName = modelname
		note.Params.Note.Fields.Front = englishWord
		note.Params.Note.Fields.Back = translatedWord
		note.Params.Note.Tags = []string{""}

		noteBytes, _ := json.Marshal(note)

		// create a context with a 20-second timeout for all requests
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// send the note to AnkiConnect
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, AnkiConnectURL, bytes.NewBuffer(noteBytes))
		if err != nil {
			printError(fmt.Sprintf("unable to create a req to anki conntect: %s", err.Error()))
			os.Exit(1)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			printError(fmt.Sprintf("unable to send note: %s", err.Error()))
			os.Exit(1)
		}
		defer res.Body.Close()

		// decode the response and handle possible errors
		resMap := make(map[string]any)
		_ = json.NewDecoder(res.Body).Decode(&resMap)

		er, ok := resMap["error"]
		if ok {
			if er != nil {
				printError(er.(string))
				os.Exit(1)
			}

			color.Green("✓ added: %s", fmt.Sprintf("%s - %s\n", englishWord, translatedWord))
		}
	}
}

func printError(msg string) {
	color.Red("✗ error: %s", msg)
}
