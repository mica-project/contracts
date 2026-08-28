package sangvis

import "encoding/json/v2"

type SangvisChip struct {
	Code                string   `json:"code"`
	Des                 []string `json:"des"`
	DevBatteryNum       int64    `json:"dev_battery_num"`
	DevTime             int64    `json:"dev_time"`
	Id                  int64    `json:"id"`
	Name                string   `json:"name"`
	PassiveMissionSkill string   `json:"passive_mission_skill"`
	Type                int64    `json:"type"`
}

type SangvisChipList []SangvisChip

func (s *SangvisChip) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
