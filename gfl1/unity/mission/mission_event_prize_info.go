package mission

import "encoding/json/v2"

type MissionEventPrizeInfo struct {
	BossHpBars string `json:"boss_hp_bars"`
	MissionId  string `json:"mission_id"`
	PrizeId    string `json:"prize_id"`
}

type MissionEventPrizeInfoList []MissionEventPrizeInfo

func (s *MissionEventPrizeInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
