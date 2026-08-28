package chess

import "encoding/json/v2"

type ChessCreationPerform struct {
	Code                 string  `json:"code"`
	DestinationType      int64   `json:"destination_type"`
	Duration             string  `json:"duration"`
	Id                   int64   `json:"id"`
	RouteHight           int64   `json:"route_hight"`
	RouteType            int64   `json:"route_type"`
	Scale                []int64 `json:"scale"`
	Speed                int64   `json:"speed"`
	SpinVelocity         int64   `json:"spin_velocity"`
	StartType            int64   `json:"start_type"`
	TriggerCreation      []int64 `json:"trigger_creation"`
	TriggerCreationDelay []int64 `json:"trigger_creation_delay"`
}

type ChessCreationPerformList []ChessCreationPerform

func (s *ChessCreationPerform) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
