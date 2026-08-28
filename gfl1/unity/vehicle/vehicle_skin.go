package vehicle

import "encoding/json/v2"

type VehicleSkin struct {
	ClassId      int64    `json:"class_id"`
	Code         string   `json:"code"`
	Id           int64    `json:"id"`
	Introduction []string `json:"introduction"`
	IsShow       int64    `json:"is_show"`
	ItemId       int64    `json:"item_id"`
	Name         string   `json:"name"`
	Obtain       string   `json:"obtain"`
	Type         int64    `json:"type"`
	VehicleId    string   `json:"vehicle_id"`
}

type VehicleSkinList []VehicleSkin

func (s *VehicleSkin) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
