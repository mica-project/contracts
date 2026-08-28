package theater

import "encoding/json/v2"

type TheaterSelection struct {
	Id                  int64  `json:"id"`
	Name                string `json:"name"`
	ScoutMaterial       int64  `json:"scout_material"`
	ScoutMaterialNumber string `json:"scout_material_number"`
	ScoutPt             int64  `json:"scout_pt"`
}

type TheaterSelectionList []TheaterSelection

func (s *TheaterSelection) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
