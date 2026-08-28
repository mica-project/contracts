package quest

import "encoding/json/v2"

type CareerQuestGrade struct {
	Description string `json:"description"`
	Grade       int64  `json:"grade"`
	Group       int64  `json:"group"`
	Id          int64  `json:"id"`
	Name        string `json:"name"`
}

type CareerQuestGradeList []CareerQuestGrade

func (s *CareerQuestGrade) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
