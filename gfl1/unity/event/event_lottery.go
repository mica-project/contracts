package event

import "encoding/json/v2"

type EventLottery struct {
	BgCode     string  `json:"bg_code"`
	EndTime    string  `json:"end_time"`
	Id         int64   `json:"id"`
	LeftCode   string  `json:"left_code"`
	LotteryIds []int64 `json:"lottery_ids"`
	StartTime  string  `json:"start_time"`
}

type EventLotteryList []EventLottery

func (s *EventLottery) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
