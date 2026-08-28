package sangvis

import "encoding/json/v2"

type SangvisCharacterType struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type SangvisCharacterTypeList []SangvisCharacterType

func (s *SangvisCharacterType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
