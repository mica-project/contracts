package gun

import "encoding/json/v2"

type GunTag struct {
	Description string `json:"description"`
	Id          int64  `json:"id"`
	IsShow      int64  `json:"is_show"`
	Tag         string `json:"tag"`
	TagType     int64  `json:"tag_type"`
}

type GunTagList []GunTag

func (s *GunTag) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
