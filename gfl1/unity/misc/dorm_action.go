package misc

import "encoding/json/v2"

type DormAction struct {
	Id          int64  `json:"id"`
	ShadowExist int64  `json:"shadow_exist"`
	SpineName   string `json:"spine_name"`
}

type DormActionList []DormAction

func (s *DormAction) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
