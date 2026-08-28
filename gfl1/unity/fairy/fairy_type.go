package fairy

import "encoding/json/v2"

type FairyType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type FairyTypeList []FairyType

func (s *FairyType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
