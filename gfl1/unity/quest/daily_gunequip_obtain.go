package quest

import "encoding/json/v2"

type DailyGunequipObtain struct {
	DropDifficulty []string `json:"drop_difficulty"`
	EquipOrGunId   int64    `json:"equip_or_gun_id"`
	GetType        int64    `json:"get_type"`
	Id             int64    `json:"id"`
	Type           int64    `json:"type"`
}

type DailyGunequipObtainList []DailyGunequipObtain

func (s *DailyGunequipObtain) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
