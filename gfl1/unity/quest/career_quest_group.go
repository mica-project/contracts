package quest

import "encoding/json/v2"

type CareerQuestGroup struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Type  int64  `json:"type"`
}

type CareerQuestGroupList []CareerQuestGroup

func (s *CareerQuestGroup) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
