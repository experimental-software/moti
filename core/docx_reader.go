package core

import (
	"regexp"
	"strings"

	"github.com/nguyenthenguyen/docx"
	"github.com/orisano/gosax"
)

func ReadDocxFile(path string) []string {
	var result []string

	r, _ := docx.ReadDocxFile(path)

	content := r.Editable().GetContent()

	reader := strings.NewReader(content)

	r2 := gosax.NewReader(reader)
	for {
		event, err := r2.Event()
		if err != nil {
			//log.Fatal(err)
		}
		if event.Type() == gosax.EventEOF {
			break
		}
		text := string(event.Bytes)
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		re := regexp.MustCompile(`<.*?>`)
		match := re.FindStringSubmatch(text)
		if len(match) == 1 {
			continue
		}
		result = append(result, text)
	}

	return result
}
