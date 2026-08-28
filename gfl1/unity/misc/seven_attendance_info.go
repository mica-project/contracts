package misc

import "encoding/json/v2"

type SevenAttendanceInfo struct {
	Day     string `json:"day"`
	ItemIds string `json:"item_ids"`
	Type    string `json:"type"`
}

type SevenAttendanceInfoList []SevenAttendanceInfo

func (s *SevenAttendanceInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
