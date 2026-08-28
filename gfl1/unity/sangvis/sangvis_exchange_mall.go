package sangvis

import "encoding/json/v2"

type SangvisExchangeMall struct {
	CostItem  int64  `json:"cost_item"`
	CostNum   int64  `json:"cost_num"`
	EndTime   string `json:"end_time"`
	Id        int64  `json:"id"`
	PrizeId   int64  `json:"prize_id"`
	Quota     int64  `json:"quota"`
	StartTime string `json:"start_time"`
}

type SangvisExchangeMallList []SangvisExchangeMall

func (s *SangvisExchangeMall) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
