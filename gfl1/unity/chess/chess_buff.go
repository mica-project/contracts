package chess

import "encoding/json/v2"

type ChessBuff struct {
	Code         string `json:"code"`
	CreationId   int64  `json:"creation_id"`
	Description  string `json:"description"`
	Duration     int64  `json:"duration"`
	DurationType int64  `json:"duration_type"`
	Id           int64  `json:"id"`
	MaxTier      int64  `json:"max_tier"`
	Name         string `json:"name"`
	Parameter    string `json:"parameter"`
	Type         int64  `json:"type"`
}

type ChessBuffList []ChessBuff

func (s *ChessBuff) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
