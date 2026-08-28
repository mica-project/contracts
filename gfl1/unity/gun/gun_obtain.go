package gun

import "encoding/json/v2"

type GunObtain struct {
	Description []string `json:"description"`
}

type GunObtainList []GunObtain

func (s *GunObtain) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
