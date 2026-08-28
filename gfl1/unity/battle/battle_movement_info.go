package battle

import "encoding/json/v2"

type BattleMovementInfo struct {
	Cd          string  `json:"cd"`
	Distance    string  `json:"distance"`
	Duration    string  `json:"duration"`
	Id          string  `json:"id"`
	IsCharacter string  `json:"is_character"`
	Offset      []int64 `json:"offset"`
	Target      string  `json:"target"`
}

type BattleMovementInfoList []BattleMovementInfo

func (s *BattleMovementInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
