package vehicle

import "encoding/json/v2"

type VehicleSkinClass struct {
	Desc       string   `json:"desc"`
	Id         int64    `json:"id"`
	Name       string   `json:"name"`
	PicCode    []string `json:"pic_code"`
	VehicleIds []int64  `json:"vehicle_ids"`
}

type VehicleSkinClassList []VehicleSkinClass

func (s *VehicleSkinClass) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
