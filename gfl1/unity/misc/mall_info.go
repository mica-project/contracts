package misc

import "encoding/json/v2"

type MallInfo struct {
	ClassificationId string `json:"classification_id"`
	Discount         string `json:"discount"`
	Gemprice         string `json:"gemprice"`
	Id               string `json:"id"`
	Mp               string `json:"mp"`
	Type             string `json:"type"`
}

type MallInfoList []MallInfo

func (s *MallInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
