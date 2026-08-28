package misc

import "encoding/json/v2"

type CommanderRankingTypes struct {
	ClassId int64   `json:"class_id"`
	Id      int64   `json:"id"`
	Title   string  `json:"title"`
	Weight  float64 `json:"weight"`
}

type CommanderRankingTypesList []CommanderRankingTypes

func (s *CommanderRankingTypes) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
