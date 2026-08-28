package squad

import "encoding/json/v2"

type SquadDataDaily struct {
	Content string  `json:"content"`
	Count   int64   `json:"count"`
	Id      int64   `json:"id"`
	Prize   []int64 `json:"prize"`
	Rank    int64   `json:"rank"`
	Title   string  `json:"title"`
	Type    string  `json:"type"`
}

type SquadDataDailyList []SquadDataDaily

func (s *SquadDataDaily) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
