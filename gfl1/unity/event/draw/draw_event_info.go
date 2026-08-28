package draw

import "encoding/json/v2"

type DrawEventInfo struct {
	AmountCoordinate []int64  `json:"amount_coordinate"`
	BgRes            string   `json:"bg_res"`
	CanTenDraws      string   `json:"can_ten_draws"`
	DropIds          []string `json:"drop_ids"`
	EndTime          string   `json:"end_time"`
	Id               string   `json:"id"`
	IsShow           string   `json:"is_show"`
	ItemId           string   `json:"item_id"`
	StartTime        string   `json:"start_time"`
	TitleRes         string   `json:"title_res"`
	Type             string   `json:"type"`
}

type DrawEventInfoList []DrawEventInfo

func (s *DrawEventInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
