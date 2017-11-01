package gociba

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// WordExplanation wraps explanation for a word
type WordExplanation struct {
	UsPronounciation string
	UkPronounciation string
	UsMp3            string
	UkMp3            string
	Meaning          string
	Spelling         string
}

// String formats the word explanation
func (w *WordExplanation) String() string {
	buffer := bytes.Buffer{}
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "    ")
	encoder.Encode(w)
	return string(buffer.Bytes())
}

// LookupWord finds a word from iciba.com
func LookupWord(word string) (*WordExplanation, error) {
	url := fmt.Sprintf("http://www.iciba.com/%v", word)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	explanation := &WordExplanation{Spelling: word}

	// find pronounciation
	doc.Find(".base-speak span span").Each(func(i int, s *goquery.Selection) {
		text := strings.Trim(s.Text(), " \n")
		script, _ := s.Parent().Find("i[ms-on-mouseover]").First().Attr("ms-on-mouseover")
		mp3URL := strings.Split(script, "'")[1]
		if strings.Contains(text, "英") {
			explanation.UkPronounciation = strings.Split(text, " ")[1]
			explanation.UkMp3 = mp3URL
		} else if strings.Contains(text, "美") {
			explanation.UsPronounciation = strings.Split(text, " ")[1]
			explanation.UsMp3 = mp3URL
		}
	})

	// find explanation
	text := doc.Find(".base-list").First().Text()
	text = strings.Trim(text, " \n")
	text = strings.Replace(text, "\n", " ", -1)
	text = strings.Replace(text, "  ", "", -1)
	explanation.Meaning = text

	return explanation, nil
}
