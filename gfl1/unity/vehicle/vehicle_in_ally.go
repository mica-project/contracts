package vehicle

import "encoding/json/v2"

type VehicleInAlly struct {
	HArmorPiercing  int64    `json:"H_armor_piercing"`
	LArmorPiercing  int64    `json:"L_armor_piercing"`
	AtkSpeed        int64    `json:"atk_speed"`
	CommonComponent []string `json:"common_component"`
	Component       []string `json:"component"`
	Crew            []string `json:"crew"`
	Exp             int64    `json:"exp"`
	HeavyDamage     int64    `json:"heavy_damage"`
	Hit             int64    `json:"hit"`
	Hp              int64    `json:"hp"`
	Id              int64    `json:"id"`
	Life            int64    `json:"life"`
	LightDamage     int64    `json:"light_damage"`
	Precision       int64    `json:"precision"`
	Reload          int64    `json:"reload"`
	UnlockedNodes   []int64  `json:"unlocked_nodes"`
	VehicleId       int64    `json:"vehicle_id"`
}

type VehicleInAllyList []VehicleInAlly

func (s *VehicleInAlly) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
