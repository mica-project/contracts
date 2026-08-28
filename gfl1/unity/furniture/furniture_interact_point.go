package furniture

import "encoding/json/v2"

type FurnitureInteractPoint struct {
	Direction int64  `json:"direction"`
	EndAction string `json:"end_action"`
	GunAction string `json:"gun_action"`
	Id        int64  `json:"id"`
}

type FurnitureInteractPointList []FurnitureInteractPoint

func (s *FurnitureInteractPoint) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
