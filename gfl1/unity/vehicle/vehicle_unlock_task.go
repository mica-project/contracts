package vehicle

import "encoding/json/v2"

type VehicleUnlockTask struct {
	Count       string `json:"count"`
	Description string `json:"description"`
	GroupType   int64  `json:"group_type"`
	Id          int64  `json:"id"`
	TaskType    string `json:"task_type"`
	VehicleId   int64  `json:"vehicle_id"`
}

type VehicleUnlockTaskList []VehicleUnlockTask

func (s *VehicleUnlockTask) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
