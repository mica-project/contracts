package chess

import "encoding/json/v2"

type ChessModel struct {
	Id              int64   `json:"id"`
	Name            string  `json:"name"`
	PlayerNum       int64   `json:"player_num"`
	Prize           []int64 `json:"prize"`
	TeamDes         string  `json:"team_des"`
	TeamNum         int64   `json:"team_num"`
	TeamPlayerlimit string  `json:"team_playerlimit"`
}

type ChessModelList []ChessModel

func (s *ChessModel) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
