package quest

import "encoding/json/v2"

type DailyMapRoute struct {
	MapId    int64   `json:"map_id"`
	MapLabel int64   `json:"map_label"`
	Route    []int64 `json:"route"`
}

type DailyMapRouteList []DailyMapRoute

func (s *DailyMapRoute) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
