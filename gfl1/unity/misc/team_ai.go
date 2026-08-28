package misc

import "encoding/json/v2"

type TeamAi struct {
	AiType      int64  `json:"ai_type"`
	Color       string `json:"color"`
	Description string `json:"description"`
	ForceId     int64  `json:"force_id"`
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Pic         string `json:"pic"`
}

type TeamAiList []TeamAi

func (s *TeamAi) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
