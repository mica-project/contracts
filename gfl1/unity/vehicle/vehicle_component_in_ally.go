package vehicle

import "encoding/json/v2"

type VehicleComponentInAlly struct {
	ComponentId int64  `json:"component_id"`
	HeavyDamage int64  `json:"heavy_damage"`
	Id          int64  `json:"id"`
	Level       int64  `json:"level"`
	Skill       string `json:"skill"`
}

type VehicleComponentInAllyList []VehicleComponentInAlly

func (s *VehicleComponentInAlly) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
