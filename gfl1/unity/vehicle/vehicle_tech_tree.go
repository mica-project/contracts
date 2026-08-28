package vehicle

import "encoding/json/v2"

type VehicleTechTree struct {
	BasicEffect int64   `json:"basic_effect"`
	Dodge       int64   `json:"dodge"`
	Hp          int64   `json:"hp"`
	Id          int64   `json:"id"`
	NodeType    int64   `json:"node_type"`
	UnlockCost  []int64 `json:"unlock_cost"`
	UnlockExp   int64   `json:"unlock_exp"`
	VehicleId   int64   `json:"vehicle_id"`
}

type VehicleTechTreeList []VehicleTechTree

func (s *VehicleTechTree) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
