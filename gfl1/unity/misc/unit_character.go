package misc

import "encoding/json/v2"

type UnitCharacter struct {
	CharacterDes string `json:"character_des"`
	Id           int64  `json:"id"`
	Name         string `json:"name"`
}

type UnitCharacterList []UnitCharacter

func (s *UnitCharacter) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
