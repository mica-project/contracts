package exploration

import "encoding/json/v2"

type ExploreAffairClient struct {
	AreaId       string   `json:"area_id"`
	Content      []string `json:"content"`
	Id           int64    `json:"id"`
	NecessaryNum []int64  `json:"necessary_num"`
	ScriptType   string   `json:"script_type"`
	Weight       int64    `json:"weight"`
}

type ExploreAffairClientList []ExploreAffairClient

func (s *ExploreAffairClient) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
