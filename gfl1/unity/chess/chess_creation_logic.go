package chess

import "encoding/json/v2"

type ChessCreationLogic struct {
}

type ChessCreationLogicList []ChessCreationLogic

func (s *ChessCreationLogic) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
