package battle

import "encoding/json/v2"

type FightEnvironmentConfig struct {
	Id                    int64  `json:"id"`
	TransformNumber       string `json:"transform_number"`
	TransformResultAdd    string `json:"transform_result_add"`
	TransformResultDelete string `json:"transform_result_delete"`
	TransformType         int64  `json:"transform_type"`
}

type FightEnvironmentConfigList []FightEnvironmentConfig

func (s *FightEnvironmentConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
