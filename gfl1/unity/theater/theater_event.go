package theater

import "encoding/json/v2"

type TheaterEvent struct {
	ApRecover      int64    `json:"ap_recover"`
	Background     string   `json:"background"`
	CloseTime      string   `json:"close_time"`
	CoreTheater    int64    `json:"core_theater"`
	EndTime        string   `json:"end_time"`
	GiftBackground string   `json:"gift_background"`
	GiftFigure     string   `json:"gift_figure"`
	Id             int64    `json:"id"`
	Name           string   `json:"name"`
	OpeningTime    []int64  `json:"opening_time"`
	PrimaryTheater []int64  `json:"primary_theater"`
	PtGift         []string `json:"pt_gift"`
	StartTime      string   `json:"start_time"`
}

type TheaterEventList []TheaterEvent

func (s *TheaterEvent) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
