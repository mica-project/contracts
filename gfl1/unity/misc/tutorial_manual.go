package misc

import "encoding/json/v2"

type TutorialManual struct {
	ChildIds []int64  `json:"child_ids"`
	Code     string   `json:"code"`
	Id       int64    `json:"id"`
	SubTitle []string `json:"sub_title"`
	Title    string   `json:"title"`
	Type     int64    `json:"type"`
}

type TutorialManualList []TutorialManual

func (s *TutorialManual) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
