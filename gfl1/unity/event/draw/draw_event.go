package draw

import "encoding/json/v2"

type DrawEvent struct {
	AmountCoordinate []int64  `json:"amount_coordinate"`
	BgRes            string   `json:"bg_res"`
	CanTenDraws      int64    `json:"can_ten_draws"`
	DropIds          []string `json:"drop_ids"`
	EndTime          string   `json:"end_time"`
	Id               int64    `json:"id"`
	IsShow           int64    `json:"is_show"`
	ItemId           int64    `json:"item_id"`
	StartTime        string   `json:"start_time"`
	TitleRes         string   `json:"title_res"`
	Type             int64    `json:"type"`
}

type DrawEventList []DrawEvent

func (s *DrawEvent) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
