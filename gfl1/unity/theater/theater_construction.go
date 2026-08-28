package theater

import "encoding/json/v2"

type TheaterConstruction struct {
	AreaId         int64  `json:"area_id"`
	Code           string `json:"code"`
	Description    string `json:"description"`
	Effect         string `json:"effect"`
	GroupId        int64  `json:"group_id"`
	Id             int64  `json:"id"`
	MaterialItem   int64  `json:"material_item"`
	MaterialNumber int64  `json:"material_number"`
	MaterialRate   int64  `json:"material_rate"`
	Name           string `json:"name"`
}

type TheaterConstructionList []TheaterConstruction

func (s *TheaterConstruction) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
