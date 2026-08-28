package misc

import "encoding/json/v2"

type SimulationBattleReward struct {
	GiftId    string `json:"gift_id"`
	ItemId    string `json:"item_id"`
	MissionId int64  `json:"mission_id"`
}

type SimulationBattleRewardList []SimulationBattleReward

func (s *SimulationBattleReward) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
