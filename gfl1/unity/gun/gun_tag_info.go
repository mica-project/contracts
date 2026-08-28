package gun

import "encoding/json/v2"

type GunTagInfo struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	IsShow      string `json:"is_show"`
	Tag         string `json:"tag"`
	TagType     string `json:"tag_type"`
}

type GunTagInfoList []GunTagInfo

func (s *GunTagInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
