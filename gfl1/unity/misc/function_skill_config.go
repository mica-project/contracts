package misc

import "encoding/json/v2"

type FunctionSkillConfig struct {
	Arguments     int64  `json:"arguments"`
	Code          string `json:"code"`
	Description   string `json:"description"`
	FunctionSkill string `json:"function_skill"`
	Id            int64  `json:"id"`
	IsShow        int64  `json:"is_show"`
	Level         int64  `json:"level"`
	MaxLevel      int64  `json:"max_level"`
	Name          string `json:"name"`
	Type          int64  `json:"type"`
}

type FunctionSkillConfigList []FunctionSkillConfig

func (s *FunctionSkillConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
