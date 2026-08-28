package vehicle

import "encoding/json/v2"

type VehicleComponentRollType struct {
	Group    int64  `json:"group"`
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	RollCode string `json:"roll_code"`
	Type     int64  `json:"type"`
}

type VehicleComponentRollTypeList []VehicleComponentRollType

func (s *VehicleComponentRollType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
