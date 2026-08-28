package misc

import "encoding/json/v2"

type Live2dMotions struct {
	Id          int64  `json:"id"`
	MotionName  string `json:"motion_name"`
	Name        string `json:"name"`
	Probability int64  `json:"probability"`
	Type        int64  `json:"type"`
}

type Live2dMotionsList []Live2dMotions

func (s *Live2dMotions) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
