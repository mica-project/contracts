package misc

import "encoding/json/v2"

type ReturnPlayer struct {
	ContinuousTime int64    `json:"continuous_time"`
	Id             int64    `json:"id"`
	PageId         int64    `json:"page_id"`
	PrizeId        []string `json:"prize_id"`
	PrizeIdSp      []string `json:"prize_id_sp"`
}

type ReturnPlayerList []ReturnPlayer

func (s *ReturnPlayer) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
