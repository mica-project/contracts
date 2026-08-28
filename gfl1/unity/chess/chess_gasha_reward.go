package chess

import "encoding/json/v2"

type ChessGashaReward struct {
	Gift       string `json:"gift"`
	Id         int64  `json:"id"`
	TicketsNum int64  `json:"tickets_num"`
	Type       int64  `json:"type"`
}

type ChessGashaRewardList []ChessGashaReward

func (s *ChessGashaReward) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
