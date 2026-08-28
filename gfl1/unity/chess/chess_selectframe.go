package chess

import "encoding/json/v2"

type ChessSelectframe struct {
	Code           []string `json:"code"`
	Des            string   `json:"des"`
	Id             int64    `json:"id"`
	ItemlanguageId []int64  `json:"itemlanguage_id"`
	Preview        []string `json:"preview"`
	SelectNum      int64    `json:"select_num"`
}

type ChessSelectframeList []ChessSelectframe

func (s *ChessSelectframe) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
