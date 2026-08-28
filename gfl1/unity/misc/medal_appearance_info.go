package misc

import "encoding/json/v2"

type MedalAppearanceInfo struct {
	BackgroundCode string `json:"background_code"`
	Id             string `json:"id"`
	PicCode        string `json:"pic_code"`
}

type MedalAppearanceInfoList []MedalAppearanceInfo

func (s *MedalAppearanceInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
