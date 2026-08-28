package gun

import "encoding/json/v2"

type GunTagTypeInfo struct {
	EnName   string `json:"en_name"`
	TypeName string `json:"type_name"`
}

type GunTagTypeInfoList []GunTagTypeInfo

func (s *GunTagTypeInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
