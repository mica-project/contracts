package fairy

import "encoding/json/v2"

type FairyLive2dMotionsInfo struct {
	Camera      string `json:"camera"`
	Id          string `json:"id"`
	MotionName  string `json:"motion_name"`
	Name        string `json:"name"`
	Probability string `json:"probability"`
	Type        string `json:"type"`
}

type FairyLive2dMotionsInfoList []FairyLive2dMotionsInfo

func (s *FairyLive2dMotionsInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
