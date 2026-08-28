package quest

import "encoding/json/v2"

type DailyMapModel struct {
	Code         string    `json:"code"`
	HeightRegion []float64 `json:"height_region"`
	Id           int64     `json:"id"`
	Type         int64     `json:"type"`
}

type DailyMapModelList []DailyMapModel

func (s *DailyMapModel) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
