package misc

import "encoding/json/v2"

type CommanderClass struct {
	Code              string   `json:"code"`
	DesYear           string   `json:"des_year"`
	Description       []string `json:"description"`
	GroupId           string   `json:"group_id"`
	Id                int64    `json:"id"`
	IsClass           int64    `json:"is_class"`
	Name              string   `json:"name"`
	Path              string   `json:"path"`
	Source            int64    `json:"source"`
	SourceDescription string   `json:"source_description"`
}

type CommanderClassList []CommanderClass

func (s *CommanderClass) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
