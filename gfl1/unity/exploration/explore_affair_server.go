package exploration

import "encoding/json/v2"

type ExploreAffairServer struct {
	Content    []string `json:"content"`
	Id         int64    `json:"id"`
	ScriptType string   `json:"script_type"`
}

type ExploreAffairServerList []ExploreAffairServer

func (s *ExploreAffairServer) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
