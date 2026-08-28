package theater

import "encoding/json/v2"

type TheaterArea struct {
	AdvantageDesc           string   `json:"advantage_desc"`
	AdvantageGun            []int64  `json:"advantage_gun"`
	AreaMissionId           int64    `json:"area_mission_id"`
	BasicScore              int64    `json:"basic_score"`
	BattleBackground        string   `json:"battle_background"`
	BossScoreCoef           string   `json:"boss_score_coef"`
	BossScoreDisplay        string   `json:"boss_score_display"`
	Construction            string   `json:"construction"`
	Description             string   `json:"description"`
	Difficulty              int64    `json:"difficulty"`
	DisplayLength           int64    `json:"display_length"`
	EndScore                int64    `json:"end_score"`
	EnemyGroup              []string `json:"enemy_group"`
	EnemyLv                 []int64  `json:"enemy_lv"`
	EnemyScore              []int64  `json:"enemy_score"`
	FightEnvironmentGroup   string   `json:"fight_environment_group"`
	FightType               int64    `json:"fight_type"`
	Id                      int64    `json:"id"`
	MaterialItem            int64    `json:"material_item"`
	Name                    string   `json:"name"`
	OccupiedBossScore       int64    `json:"occupied_boss_score"`
	OccupiedEnemyLv         []int64  `json:"occupied_enemy_lv"`
	OccupiedEnemyScore      []int64  `json:"occupied_enemy_score"`
	SuccessTypeCondition    string   `json:"success_type_condition"`
	TheaterId               int64    `json:"theater_id"`
	TheaterSpareGunNum      int64    `json:"theater_spare_gun_num"`
	TheaterSpareSangvisCost int64    `json:"theater_spare_sangvis_cost"`
	Type                    int64    `json:"type"`
}

type TheaterAreaList []TheaterArea

func (s *TheaterArea) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
