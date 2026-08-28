package misc

import "encoding/json/v2"

type AutoFormationSangvis struct {
	BackTagIds  string   `json:"back_tag_ids"`
	Description string   `json:"description"`
	FrontTagIds string   `json:"front_tag_ids"`
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	TeamInfo    []string `json:"team_info"`
	TeamTagIds  []int64  `json:"team_tag_ids"`
}

type AutoFormationSangvisList []AutoFormationSangvis

func (s *AutoFormationSangvis) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
