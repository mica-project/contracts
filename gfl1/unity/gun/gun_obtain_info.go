package gun

import "encoding/json/v2"

type GunObtainInfo struct {
	Description []string `json:"description"`
}

type GunObtainInfoList []GunObtainInfo

func (s *GunObtainInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
