package mission

import "encoding/json/v2"

type MissionWinStepControl struct {
	Id      int64  `json:"id"`
	NextId  string `json:"next_id"`
	WinStep string `json:"win_step"`
}

type MissionWinStepControlList []MissionWinStepControl

func (s *MissionWinStepControl) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
