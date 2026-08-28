package furniture

import "encoding/json/v2"

type FurnitureEstablishInfo struct {
	Description       string  `json:"description"`
	EstablishId       string  `json:"establish_id"`
	EstablishLv       string  `json:"establish_lv"`
	EstablishName     string  `json:"establish_name"`
	FurnitureId       string  `json:"furniture_id"`
	FurniturePostion  []int64 `json:"furniture_postion"`
	RoomId            string  `json:"room_id"`
	UpdateFurnitureId string  `json:"update_furniture_id"`
}

type FurnitureEstablishInfoList []FurnitureEstablishInfo

func (s *FurnitureEstablishInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
