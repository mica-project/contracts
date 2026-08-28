package skin

import "encoding/json/v2"

type SkinBonusInfo struct {
	ActiveNum   string   `json:"active_num"`
	Description []string `json:"description"`
	Icon        string   `json:"icon"`
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Prize       string   `json:"prize"`
	SkinTheme   string   `json:"skin_theme"`
	Type        string   `json:"type"`
}

type SkinBonusInfoList []SkinBonusInfo

func (s *SkinBonusInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
