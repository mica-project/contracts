package quest

import "encoding/json/v2"

type DailyExmissionGroup struct {
	Effect       int64 `json:"effect"`
	Id           int64 `json:"id"`
	MissionFloor int64 `json:"mission_floor"`
	MissionId    int64 `json:"mission_id"`
	StateType    int64 `json:"state_type"`
	VehicleId    int64 `json:"vehicle_id"`
}

type DailyExmissionGroupList []DailyExmissionGroup

func (s *DailyExmissionGroup) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
