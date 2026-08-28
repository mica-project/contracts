package carnival

import "encoding/json/v2"

type CarnivalLabelInfo struct {
	CarnivalTasks []int64 `json:"carnival_tasks"`
	Id            string  `json:"id"`
	LabelText     string  `json:"label_text"`
	StartTime     string  `json:"start_time"`
}

type CarnivalLabelInfoList []CarnivalLabelInfo

func (s *CarnivalLabelInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
