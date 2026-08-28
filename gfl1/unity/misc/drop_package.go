package misc

import "encoding/json/v2"

type DropPackage struct {
	Id int64 `json:"id"`
}

type DropPackageList []DropPackage

func (s *DropPackage) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
