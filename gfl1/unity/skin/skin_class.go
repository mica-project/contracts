package skin

import "encoding/json/v2"

type SkinClass struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	ThemeType int64  `json:"theme_type"`
}

type SkinClassList []SkinClass

func (s *SkinClass) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
