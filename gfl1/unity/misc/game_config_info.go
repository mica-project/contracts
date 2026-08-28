package misc

import "encoding/json/v2"

type GameConfigInfo struct {
	ClientRequire  string `json:"client_require"`
	ParameterName  string `json:"parameter_name"`
	ParameterType  string `json:"parameter_type"`
	ParameterValue string `json:"parameter_value"`
}

type GameConfigInfoList []GameConfigInfo

func (s *GameConfigInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
