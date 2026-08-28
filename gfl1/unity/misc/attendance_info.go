package misc

import "encoding/json/v2"

type AttendanceInfo struct {
	AttendanceType string `json:"attendance_type"`
	Day            string `json:"day"`
	EndTime        string `json:"end_time"`
	Mp             string `json:"mp"`
	StartTime      string `json:"start_time"`
}

type AttendanceInfoList []AttendanceInfo

func (s *AttendanceInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
