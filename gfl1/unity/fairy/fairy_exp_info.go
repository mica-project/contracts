package fairy

import "encoding/json/v2"

type FairyExpInfo struct {
	Id string `json:"id"`
}

type FairyExpInfoList []FairyExpInfo

func (s *FairyExpInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
