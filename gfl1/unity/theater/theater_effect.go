package theater

import "encoding/json/v2"

type TheaterEffect struct {
	Description string `json:"description"`
	Duration    int64  `json:"duration"`
	Icon        string `json:"icon"`
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Type        int64  `json:"type"`
}

type TheaterEffectList []TheaterEffect

func (s *TheaterEffect) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
