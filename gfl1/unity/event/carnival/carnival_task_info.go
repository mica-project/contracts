package carnival

import "encoding/json/v2"

type CarnivalTaskInfo struct {
	CarnivalTypeId string `json:"carnival_type_id"`
	Content        string `json:"content"`
	Count          string `json:"count"`
	Id             string `json:"id"`
	PrizeId        string `json:"prize_id"`
	Title          string `json:"title"`
	Type           string `json:"type"`
}

type CarnivalTaskInfoList []CarnivalTaskInfo

func (s *CarnivalTaskInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
