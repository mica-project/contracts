package quest

import "encoding/json/v2"

type Weekly struct {
	Content         string `json:"content"`
	Count           int64  `json:"count"`
	ExchangePrizeId int64  `json:"exchange_prize_id"`
	Identity        int64  `json:"identity"`
	PrizeId         int64  `json:"prize_id"`
	Title           string `json:"title"`
	Type            string `json:"type"`
}

type WeeklyList []Weekly

func (s *Weekly) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
