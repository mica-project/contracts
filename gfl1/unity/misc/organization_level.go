package misc

import "encoding/json/v2"

type OrganizationLevel struct {
	Description    string `json:"description"`
	Icon           string `json:"icon"`
	Id             int64  `json:"id"`
	Level          int64  `json:"level"`
	OrganizationId int64  `json:"organization_id"`
}

type OrganizationLevelList []OrganizationLevel

func (s *OrganizationLevel) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
