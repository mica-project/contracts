package misc

import "encoding/json/v2"

type MailContent struct {
	Content []string `json:"content"`
	Id      int64    `json:"id"`
	Title   string   `json:"title"`
}

type MailContentList []MailContent

func (s *MailContent) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
