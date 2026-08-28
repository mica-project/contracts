package misc

import "encoding/json/v2"

type ChatFixPhrases struct {
	Content  string `json:"content"`
	Group    int64  `json:"group"`
	Id       int64  `json:"id"`
	Type     int64  `json:"type"`
	TypeName string `json:"type_name"`
}

type ChatFixPhrasesList []ChatFixPhrases

func (s *ChatFixPhrases) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
