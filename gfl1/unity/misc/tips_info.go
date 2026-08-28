package misc

import "encoding/json/v2"

type TipsInfo struct {
	BgPath         string  `json:"bg_path"`
	Id             string  `json:"id"`
	RegTarget      string  `json:"regTarget"`
	TargetPath     string  `json:"targetPath"`
	TipPos         []int64 `json:"tipPos"`
	TipSize        []int64 `json:"tipSize"`
	TipsID         string  `json:"tipsID"`
	TipsManagePath string  `json:"tipsManagePath"`
	Title          string  `json:"title"`
}

type TipsInfoList []TipsInfo

func (s *TipsInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
