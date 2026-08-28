package quest

import "encoding/json/v2"

type CareerQuest struct {
	Content string `json:"content"`
	Count   int64  `json:"count"`
	GradeId int64  `json:"grade_id"`
	Id      int64  `json:"id"`
	NewType int64  `json:"new_type"`
	PrizeId int64  `json:"prize_id"`
	Sort    int64  `json:"sort"`
	Title   string `json:"title"`
	Type    string `json:"type"`
}

type CareerQuestList []CareerQuest

func (s *CareerQuest) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
