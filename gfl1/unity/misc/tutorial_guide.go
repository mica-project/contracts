package misc

import "encoding/json/v2"

type TutorialGuide struct {
	ChildIds  []int64 `json:"child_ids"`
	Id        int64   `json:"id"`
	SubTitle  string  `json:"sub_title"`
	Title     string  `json:"title"`
	TitleCode string  `json:"title_code"`
	Type      int64   `json:"type"`
}

type TutorialGuideList []TutorialGuide

func (s *TutorialGuide) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
