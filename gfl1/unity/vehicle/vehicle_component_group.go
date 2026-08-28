package vehicle

import "encoding/json/v2"

type VehicleComponentGroup struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type VehicleComponentGroupList []VehicleComponentGroup

func (s *VehicleComponentGroup) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
