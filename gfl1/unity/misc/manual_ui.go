package misc

import "encoding/json/v2"

type ManualUi struct {
	CodeA        string `json:"code_a"`
	CodeB        string `json:"code_b"`
	DescriptionA string `json:"description_a"`
	DescriptionB string `json:"description_b"`
	DescriptionC string `json:"description_c"`
	Id           int64  `json:"id"`
}

type ManualUiList []ManualUi

func (s *ManualUi) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
