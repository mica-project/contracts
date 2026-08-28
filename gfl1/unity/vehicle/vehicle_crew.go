package vehicle

import "encoding/json/v2"

type VehicleCrew struct {
	AtkSpeed            int64    `json:"atk_speed"`
	AvailableUnitType   []string `json:"available_unit_type"`
	BasicEffect         int64    `json:"basic_effect"`
	Description         string   `json:"description"`
	Hit                 int64    `json:"hit"`
	Id                  int64    `json:"id"`
	IfShow              int64    `json:"if_show"`
	Name                string   `json:"name"`
	RequiredLevel       int64    `json:"required_level"`
	RequiredNumber      int64    `json:"required_number"`
	UnitTypeDescription string   `json:"unit_type_description"`
}

type VehicleCrewList []VehicleCrew

func (s *VehicleCrew) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
