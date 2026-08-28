package misc

import "encoding/json/v2"

type OrganizationBounty struct {
	Id                int64 `json:"id"`
	OrgId             int64 `json:"org_id"`
	OrganizationLevel int64 `json:"organization_level"`
	Quantity          int64 `json:"quantity"`
	Reward            int64 `json:"reward"`
}

type OrganizationBountyList []OrganizationBounty

func (s *OrganizationBounty) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
