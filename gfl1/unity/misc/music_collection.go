package misc

import "encoding/json/v2"

type MusicCollection struct {
	Code   []string `json:"code"`
	Detail int64    `json:"detail"`
	Id     int64    `json:"id"`
	Name   string   `json:"name"`
	Type   string   `json:"type"`
}

type MusicCollectionList []MusicCollection

func (s *MusicCollection) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
