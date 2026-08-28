package sangvis

import "encoding/json/v2"

type SangvisSign struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type SangvisSignList []SangvisSign

func (s *SangvisSign) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
