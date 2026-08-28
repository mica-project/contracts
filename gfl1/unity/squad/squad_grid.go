package squad

import "encoding/json/v2"

type SquadGrid struct {
	Code       string  `json:"code"`
	Grid       []int64 `json:"grid"`
	GridNumber int64   `json:"grid_number"`
	Id         int64   `json:"id"`
	RankWeight []int64 `json:"rank_weight"`
}

type SquadGridList []SquadGrid

func (s *SquadGrid) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
