package quest

import "encoding/json/v2"

type DailyMap struct {
	Coordinate string `json:"coordinate"`
	Id         int64  `json:"id"`
	MapLabel   int64  `json:"map_label"`
	Neighbor   string `json:"neighbor"`
	Start      int64  `json:"start"`
	Type       string `json:"type"`
}

type DailyMapList []DailyMap

func (s *DailyMap) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
