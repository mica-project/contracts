package squad

import "encoding/json/v2"

type SquadColor struct {
	FliterPic  string  `json:"fliter_pic"`
	FliterText string  `json:"fliter_text"`
	Id         int64   `json:"id"`
	RankWeight []int64 `json:"rank_weight"`
	Rgb        []int64 `json:"rgb"`
}

type SquadColorList []SquadColor

func (s *SquadColor) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
