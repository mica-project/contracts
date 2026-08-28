package gun

import "encoding/json/v2"

type GunExpInfo struct {
	Exp string `json:"exp"`
	Lv  string `json:"lv"`
}

type GunExpInfoList []GunExpInfo

func (s *GunExpInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
