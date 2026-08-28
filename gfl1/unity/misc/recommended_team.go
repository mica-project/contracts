package misc

import "encoding/json/v2"

type RecommendedTeam struct {
	AutoFormationId        int64  `json:"auto_formation_id"`
	AutoFormationSangvisId int64  `json:"auto_formation_sangvis_id"`
	Id                     int64  `json:"id"`
	Name                   string `json:"name"`
}

type RecommendedTeamList []RecommendedTeam

func (s *RecommendedTeam) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
