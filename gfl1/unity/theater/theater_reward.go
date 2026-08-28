package theater

import "encoding/json/v2"

type TheaterReward struct {
	Id             int64    `json:"id"`
	PrizeId        []string `json:"prize_id"`
	Rank           string   `json:"rank"`
	TheaterEventId int64    `json:"theater_event_id"`
	Type           int64    `json:"type"`
}

type TheaterRewardList []TheaterReward

func (s *TheaterReward) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
