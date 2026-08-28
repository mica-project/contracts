package chess

import "encoding/json/v2"

type ChessScorelevel struct {
	Code          string   `json:"code"`
	Id            int64    `json:"id"`
	Name          string   `json:"name"`
	ScoreCeilling int64    `json:"score_ceilling"`
	ScoreGet      []string `json:"score_get"`
}

type ChessScorelevelList []ChessScorelevel

func (s *ChessScorelevel) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
