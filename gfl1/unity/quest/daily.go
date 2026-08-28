package quest

import "encoding/json/v2"

type Daily struct {
	Content         string `json:"content"`
	Count           int64  `json:"count"`
	ExchangePrizeId int64  `json:"exchange_prize_id"`
	Identity        int64  `json:"identity"`
	MissionType     int64  `json:"mission_type"`
	PrizeId         int64  `json:"prize_id"`
	Title           string `json:"title"`
	Type            string `json:"type"`
}

type DailyList []Daily

func (s *Daily) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
