package misc

import "encoding/json/v2"

type Organization struct {
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Name        string `json:"name"`
}

type OrganizationList []Organization

func (s *Organization) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
