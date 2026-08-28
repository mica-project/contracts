package misc

import "encoding/json/v2"

type ReturnPlayerTask struct {
	ActivityBonus int64  `json:"activity_bonus"`
	Count         int64  `json:"count"`
	Describe      string `json:"describe"`
	Id            int64  `json:"id"`
	PrizeId       int64  `json:"prize_id"`
	Title         string `json:"title"`
	Type          string `json:"type"`
}

type ReturnPlayerTaskList []ReturnPlayerTask

func (s *ReturnPlayerTask) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
