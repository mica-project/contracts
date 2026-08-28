package misc

import "encoding/json/v2"

type DormAi struct {
	Actions    []int64 `json:"actions"`
	MinTime    []int64 `json:"min_time"`
	TimeWeight []int64 `json:"time_weight"`
	UpRate     []int64 `json:"up_rate"`
}

type DormAiList []DormAi

func (s *DormAi) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
