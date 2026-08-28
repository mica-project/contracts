package carnival

import "encoding/json/v2"

type CarnivalTaskType struct {
	Content string `json:"content"`
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Type    string `json:"type"`
}

type CarnivalTaskTypeList []CarnivalTaskType

func (s *CarnivalTaskType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
