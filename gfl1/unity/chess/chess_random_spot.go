package chess

import "encoding/json/v2"

type ChessRandomSpot struct {
	Id         int64  `json:"id"`
	SpotEffect string `json:"spot_effect"`
}

type ChessRandomSpotList []ChessRandomSpot

func (s *ChessRandomSpot) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
