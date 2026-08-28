package battle

import "encoding/json/v2"

type AllyTeam struct {
	Ai             string  `json:"ai"`
	AiContent      string  `json:"ai_content"`
	Code           string  `json:"code"`
	EnemyPanelType int64   `json:"enemy_panel_type"`
	EnemyTeamId    int64   `json:"enemy_team_id"`
	Guns           []int64 `json:"guns"`
	Id             int64   `json:"id"`
	InitialType    int64   `json:"initial_type"`
	Name           string  `json:"name"`
	NoBattleDamage int64   `json:"no_battle_damage"`
	UiImageIcon    string  `json:"ui_image_icon"`
}

type AllyTeamList []AllyTeam

func (s *AllyTeam) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
