package furniture

import "encoding/json/v2"

type FurnitureClasses struct {
	BonusDescription string   `json:"bonus_description"`
	BonusNumber      int64    `json:"bonus_number"`
	BonusPicture     string   `json:"bonus_picture"`
	Code             string   `json:"code"`
	Description      string   `json:"description"`
	Id               int64    `json:"id"`
	IsShowed         int64    `json:"is_showed"`
	KBonus           []string `json:"k_bonus"`
	Name             string   `json:"name"`
	Placement        int64    `json:"placement"`
	Rank             int64    `json:"rank"`
	Years            int64    `json:"years"`
}

type FurnitureClassesList []FurnitureClasses

func (s *FurnitureClasses) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
