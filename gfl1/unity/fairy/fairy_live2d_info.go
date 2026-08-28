package fairy

import "encoding/json/v2"

type FairyLive2dInfo struct {
	Code     string  `json:"code"`
	FitFairy string  `json:"fit_fairy"`
	Id       string  `json:"id"`
	Motions  []int64 `json:"motions"`
	Skin     string  `json:"skin"`
}

type FairyLive2dInfoList []FairyLive2dInfo

func (s *FairyLive2dInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
