package misc

import "encoding/json/v2"

type SevenSpendpointInfo struct {
	Day        string `json:"day"`
	Gem        string `json:"gem"`
	Mp         string `json:"mp"`
	SpendPoint string `json:"spend_point"`
	Type       string `json:"type"`
}

type SevenSpendpointInfoList []SevenSpendpointInfo

func (s *SevenSpendpointInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
