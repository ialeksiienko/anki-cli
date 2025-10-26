package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

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
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s \"english-англійська\"\n", os.Args[0])
		return
	}

	text := os.Args[1]
	lines := strings.Split(text, "\n")

	for _, l := range lines {
		words := strings.Split(l, "-")

		englishWord, translatedWord := strings.TrimSpace(words[0]), strings.TrimSpace(words[1])

		var note Note
		note.Action = "addNote"
		note.Version = 6
		note.Params.Note.DeckName = "English"
		note.Params.Note.ModelName = "Basic"
		note.Params.Note.Fields.Front = englishWord
		note.Params.Note.Fields.Back = translatedWord
		note.Params.Note.Tags = []string{""}

		noteBytes, _ := json.Marshal(note)
		res, err := http.Post("http://localhost:8765", "application/json", bytes.NewBuffer(noteBytes))
		if err != nil {
			fmt.Printf("unable to send note: %s\n", err.Error())
			return
		}
		defer res.Body.Close()

		resMap := make(map[string]any)
		_ = json.NewDecoder(res.Body).Decode(&resMap)

		fmt.Printf("%+v\n", resMap)

		fmt.Printf("Added: %s - %s\n", englishWord, translatedWord)
	}
}
