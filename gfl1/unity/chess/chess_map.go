package chess

import "encoding/json/v2"

type ChessMap struct {
	Code          string `json:"code"`
	DefaultUnlock int64  `json:"default_unlock"`
	Id            int64  `json:"id"`
	Name          string `json:"name"`
}

type ChessMapList []ChessMap

func (s *ChessMap) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
