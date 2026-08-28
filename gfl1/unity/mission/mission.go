package mission

import "encoding/json/v2"

type Mission struct {
	BossTeamId                  string   `json:"boss_team_id"`
	Difficulty                  int64    `json:"difficulty"`
	DifficultyRecommend         []string `json:"difficulty_recommend"`
	DifficultyRecommendAddendum string   `json:"difficulty_recommend_addendum"`
	EnemyAiType                 int64    `json:"enemy_ai_type"`
	ExpParameter                int64    `json:"exp_parameter"`
	ExpectEnemyDieNum           int64    `json:"expect_enemy_die_num"`
	ExpectTurn                  int64    `json:"expect_turn"`
	FogColor                    string   `json:"fog_color"`
	FogLength                   []int64  `json:"fog_length"`
	GuideCdn                    string   `json:"guide_cdn"`
	Id                          int64    `json:"id"`
	IntegrationSwitch           string   `json:"integration_switch"`
	LimitSangvis                int64    `json:"limit_sangvis"`
	LimitSquad                  int64    `json:"limit_squad"`
	LimitVehicle                int64    `json:"limit_vehicle"`
	MapInformation              []string `json:"map_information"`
	MapResName                  string   `json:"map_res_name"`
	MissionTip                  string   `json:"mission_tip"`
	Name                        string   `json:"name"`
	RecommendGunLevel           int64    `json:"recommend_gun_level"`
	RecommendedTeamIds          string   `json:"recommended_team_ids"`
	RoundConfig                 []string `json:"round_config"`
	SpecialSpotId               string   `json:"special_spot_id"`
	Sub                         int64    `json:"sub"`
	SupportAvailable            int64    `json:"support_available"`
	TurnDuration                int64    `json:"turn_duration"`
	Type                        string   `json:"type"`
	WinSpotId                   string   `json:"win_spot_id"`
}

type MissionList []Mission

func (s *Mission) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
