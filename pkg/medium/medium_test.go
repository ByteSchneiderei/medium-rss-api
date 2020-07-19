package medium_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/ByteSchneiderei/medium-rss-api/pkg/medium"
)

func TestTokenizeHTML(t *testing.T) {
	htmlContent := `<h1>This is a Test</h1><p>test for <a href="#">tokenizeHMTL</a></p>`
	expected := &Node{
		Tag:   "body",
		Text:  "",
		Attrs: []map[string]string{},
		Children: []*Node{
			{
				Tag:   "h1",
				Text:  "",
				Attrs: []map[string]string{},
				Children: []*Node{
					{
						Tag:      "",
						Text:     "This is a Test",
						Attrs:    []map[string]string{},
						Children: []*Node{},
					},
				},
			},
			{
				Tag:   "p",
				Text:  "",
				Attrs: []map[string]string{},
				Children: []*Node{
					{
						Tag:      "",
						Text:     "test for ",
						Attrs:    []map[string]string{},
						Children: []*Node{},
					},
					{
						Tag:  "a",
						Text: "",
						Attrs: []map[string]string{
							{"key": "href", "value": "#"},
						},
						Children: []*Node{
							{
								Tag:      "",
								Text:     "tokenizeHMTL",
								Attrs:    []map[string]string{},
								Children: []*Node{},
							},
						},
					},
				},
			},
		},
	}

	expectedJSON, _ := json.Marshal(expected)

	result, err := TokenizeHTML(htmlContent)
	resultJSON, _ := json.Marshal(result)

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, string(expectedJSON), string(resultJSON), "JSON string result must be similar")
}
