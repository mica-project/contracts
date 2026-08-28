package fairy

import "encoding/json/v2"

type FairyTypeInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type FairyTypeInfoList []FairyTypeInfo

func (s *FairyTypeInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
