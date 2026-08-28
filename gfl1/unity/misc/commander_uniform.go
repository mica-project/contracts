package misc

import "encoding/json/v2"

type CommanderUniform struct {
	Code         string `json:"code"`
	ColorIconId  string `json:"color_icon_id"`
	ColorNormal  string `json:"color_normal"`
	Description  string `json:"description"`
	Icon         string `json:"icon"`
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Type         int64  `json:"type"`
	UniformClass string `json:"uniform_class"`
}

type CommanderUniformList []CommanderUniform

func (s *CommanderUniform) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
