package sangvis

import "encoding/json/v2"

type Sangvis struct {
	Ai                  int64     `json:"ai"`
	AmmoAddWithnumber   int64     `json:"ammo_add_withnumber"`
	ApCost              int64     `json:"ap_cost"`
	ArmorPiercing       int64     `json:"armor_piercing"`
	AssistAttackRange   int64     `json:"assist_attack_range"`
	AttackRangeType     int64     `json:"attack_range_type"`
	Baseammo            int64     `json:"baseammo"`
	Basemre             int64     `json:"basemre"`
	Character           []int64   `json:"character"`
	Code                string    `json:"code"`
	Crit                int64     `json:"crit"`
	CritDmg             int64     `json:"crit_dmg"`
	Dialogue            string    `json:"dialogue"`
	DisplayEnemyTeam    int64     `json:"display_enemy_team"`
	DormScale           int64     `json:"dorm_scale"`
	DynamicPassiveSkill []int64   `json:"dynamic_passive_skill"`
	EatRatio            int64     `json:"eat_ratio"`
	EffectGridEffect    []string  `json:"effect_grid_effect"`
	EnName              string    `json:"en_name"`
	Extra               []string  `json:"extra"`
	Forces              int64     `json:"forces"`
	Formation           int64     `json:"formation"`
	GunTagIds           []int64   `json:"gun_tag_ids"`
	Id                  int64     `json:"id"`
	IllustrationId      int64     `json:"illustration_id"`
	Introduce           []string  `json:"introduce"`
	IsAdditional        int64     `json:"is_additional"`
	LaunchTime          string    `json:"launch_time"`
	MreAddWithnumber    int64     `json:"mre_add_withnumber"`
	Name                string    `json:"name"`
	NormalAttack        int64     `json:"normal_attack"`
	OrgId               int64     `json:"org_id"`
	PictureOffset       []float64 `json:"picture_offset"`
	PictureScale        []float64 `json:"picture_scale"`
	Rank                int64     `json:"rank"`
	RatioDodge          int64     `json:"ratio_dodge"`
	RatioHit            int64     `json:"ratio_hit"`
	RatioHp             int64     `json:"ratio_hp"`
	RatioPow            int64     `json:"ratio_pow"`
	RatioRange          int64     `json:"ratio_range"`
	RatioRate           int64     `json:"ratio_rate"`
	RatioRec            int64     `json:"ratio_rec"`
	RatioSpeed          int64     `json:"ratio_speed"`
	RecommendedTeamIds  []int64   `json:"recommended_team_ids"`
	Resolution          int64     `json:"resolution"`
	SearchRange         int64     `json:"search_range"`
	ShapeScale          []int64   `json:"shape_scale"`
	Skill1              int64     `json:"skill1"`
	Skill2              int64     `json:"skill2"`
	Skill2Type          int64     `json:"skill2_type"`
	Skill3              int64     `json:"skill3"`
	SkillAdvance        int64     `json:"skill_advance"`
	SkillResolution     []int64   `json:"skill_resolution"`
	Type                int64     `json:"type"`
	TypeChip1           []int64   `json:"type_chip1"`
	TypeChip2           []int64   `json:"type_chip2"`
	TypeChip3           []int64   `json:"type_chip3"`
}

type SangvisList []Sangvis

func (s *Sangvis) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
