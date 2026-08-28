package chess

import "encoding/json/v2"

type ChessGameConfig struct {
	Id             int64   `json:"id"`
	ParameterName  string  `json:"parameter_name"`
	ParameterType  string  `json:"parameter_type"`
	ParameterValue []int64 `json:"parameter_value"`
}

type ChessGameConfigList []ChessGameConfig

func (s *ChessGameConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
