package misc

import "encoding/json/v2"

type AdjutantSkin struct {
	Code      string `json:"code"`
	Id        int64  `json:"id"`
	ItemId    int64  `json:"item_id"`
	Name      string `json:"name"`
	ObtainTxt string `json:"obtain_txt"`
	Type      string `json:"type"`
}

type AdjutantSkinList []AdjutantSkin

func (s *AdjutantSkin) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
