package exploration

import "encoding/json/v2"

type ExploreScript struct {
	Code string `json:"code"`
	Id   int64  `json:"id"`
	Type int64  `json:"type"`
}

type ExploreScriptList []ExploreScript

func (s *ExploreScript) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
