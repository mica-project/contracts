package misc

import "encoding/json/v2"

type CommanderEmoji struct {
	Content string `json:"content"`
	Id      int64  `json:"id"`
	IsShow  int64  `json:"is_show"`
}

type CommanderEmojiList []CommanderEmoji

func (s *CommanderEmoji) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
