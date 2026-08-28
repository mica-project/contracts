package vehicle

import "encoding/json/v2"

type VehicleComponentType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type VehicleComponentTypeList []VehicleComponentType

func (s *VehicleComponentType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
