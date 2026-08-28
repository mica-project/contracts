package carnival

import "encoding/json/v2"

type CarnivalInfo struct {
	Banner    string   `json:"banner"`
	EndTime   string   `json:"end_time"`
	Id        string   `json:"id"`
	LabelIds  []int64  `json:"label_ids"`
	Name      string   `json:"name"`
	PtGift    []string `json:"pt_gift"`
	PtItemId  string   `json:"pt_item_id"`
	StartTime string   `json:"start_time"`
}

type CarnivalInfoList []CarnivalInfo

func (s *CarnivalInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
