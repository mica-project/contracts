package misc

import "encoding/json/v2"

type ShareContentInfo struct {
	Description []string `json:"description"`
	Id          string   `json:"id"`
	Platform    string   `json:"platform"`
	Title       string   `json:"title"`
}

type ShareContentInfoList []ShareContentInfo

func (s *ShareContentInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
