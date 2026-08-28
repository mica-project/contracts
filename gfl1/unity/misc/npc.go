package misc

import "encoding/json/v2"

type Npc struct {
	Ai              int64  `json:"ai"`
	Code            string `json:"code"`
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	OrgIntroduction string `json:"org_introduction"`
	Title           string `json:"title"`
	UnlockText      string `json:"unlock_text"`
}

type NpcList []Npc

func (s *Npc) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
