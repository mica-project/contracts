package furniture

import "encoding/json/v2"

type FurnitureClassesInfo struct {
	BonusDescription string   `json:"bonus_description"`
	BonusNumber      string   `json:"bonus_number"`
	BonusPicture     string   `json:"bonus_picture"`
	Code             string   `json:"code"`
	Description      string   `json:"description"`
	Id               string   `json:"id"`
	IsShowed         string   `json:"is_showed"`
	KBonus           []string `json:"k_bonus"`
	Name             string   `json:"name"`
	Placement        string   `json:"placement"`
	Rank             string   `json:"rank"`
	Years            string   `json:"years"`
}

type FurnitureClassesInfoList []FurnitureClassesInfo

func (s *FurnitureClassesInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
