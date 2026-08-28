package quest

import "encoding/json/v2"

type DailyMapSpot struct {
	Id          int64  `json:"id"`
	MapId       int64  `json:"map_id"`
	ModelCode   string `json:"model_code"`
	ModelHeight string `json:"model_height"`
	SpotId      int64  `json:"spot_id"`
	SpotType    int64  `json:"spot_type"`
}

type DailyMapSpotList []DailyMapSpot

func (s *DailyMapSpot) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
