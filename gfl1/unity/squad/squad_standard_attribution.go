package squad

import "encoding/json/v2"

type SquadStandardAttribution struct {
	AttributeType        string  `json:"attribute_type"`
	BasicRate            int64   `json:"basic_rate"`
	CpuRate              int64   `json:"cpu_rate"`
	CpuStandardAttribute float64 `json:"cpu_standard_attribute"`
	Id                   int64   `json:"id"`
	Name                 string  `json:"name"`
	RoleId               int64   `json:"role_id"`
	StandardAttribute    float64 `json:"standard_attribute"`
}

type SquadStandardAttributionList []SquadStandardAttribution

func (s *SquadStandardAttribution) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
