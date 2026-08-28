package carnival

import "encoding/json/v2"

type CarnivalTask struct {
	CarnivalTypeId int64  `json:"carnival_type_id"`
	Content        string `json:"content"`
	Count          int64  `json:"count"`
	Id             int64  `json:"id"`
	PrizeId        int64  `json:"prize_id"`
	Title          string `json:"title"`
	Type           string `json:"type"`
}

type CarnivalTaskList []CarnivalTask

func (s *CarnivalTask) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
