package misc

import "encoding/json/v2"

type TipsManageInfo struct {
	Id             string  `json:"id"`
	RegTarget      string  `json:"regTarget"`
	TargetPath     string  `json:"targetPath"`
	TipManagePos   []int64 `json:"tipManagePos"`
	TipManageSize  []int64 `json:"tipManageSize"`
	TipsManagePath string  `json:"tipsManagePath"`
	Title          string  `json:"title"`
}

type TipsManageInfoList []TipsManageInfo

func (s *TipsManageInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
