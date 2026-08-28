package squad

import "encoding/json/v2"

type Squad struct {
	AdvancedBonus           int64    `json:"advanced_bonus"`
	AmmoPart                int64    `json:"ammo_part"`
	AssistArmorPiercing     int64    `json:"assist_armor_piercing"`
	AssistDamage            int64    `json:"assist_damage"`
	AssistDefBreak          int64    `json:"assist_def_break"`
	AssistHit               int64    `json:"assist_hit"`
	AssistReload            int64    `json:"assist_reload"`
	AssistType              int64    `json:"assist_type"`
	AtkSpeed                int64    `json:"atk_speed"`
	AttackRange             int64    `json:"attack_range"`
	Baseammo                int64    `json:"baseammo"`
	Basemre                 int64    `json:"basemre"`
	BasicRate               int64    `json:"basic_rate"`
	BattleAssistRange       []int64  `json:"battle_assist_range"`
	Code                    string   `json:"code"`
	CpuId                   int64    `json:"cpu_id"`
	CpuRate                 int64    `json:"cpu_rate"`
	CritDamage              int64    `json:"crit_damage"`
	CritRate                int64    `json:"crit_rate"`
	Damage                  int64    `json:"damage"`
	DestroyCoef             int64    `json:"destroy_coef"`
	DevelopDuration         int64    `json:"develop_duration"`
	DisplayAssistAreaCoef   int64    `json:"display_assist_area_coef"`
	DisplayAssistDamageArea int64    `json:"display_assist_damage_area"`
	Dodge                   int64    `json:"dodge"`
	DormAi                  []int64  `json:"dorm_ai"`
	EnName                  string   `json:"en_name"`
	Hit                     int64    `json:"hit"`
	Hp                      int64    `json:"hp"`
	Id                      int64    `json:"id"`
	Introduce               []string `json:"introduce"`
	IsShow                  int64    `json:"is_show"`
	LaunchTime              string   `json:"launch_time"`
	Move                    int64    `json:"move"`
	MrePart                 int64    `json:"mre_part"`
	Name                    string   `json:"name"`
	NightVision             int64    `json:"night_vision"`
	NormalAttack            int64    `json:"normal_attack"`
	NormalAttackDescription int64    `json:"normal_attack_description"`
	OrgId                   int64    `json:"org_id"`
	PassiveSkill            []int64  `json:"passive_skill"`
	PerformanceSkill        int64    `json:"performance_skill"`
	PieceItemId             int64    `json:"piece_item_id"`
	Population              int64    `json:"population"`
	Skill1                  int64    `json:"skill1"`
	Skill2                  int64    `json:"skill2"`
	Skill3                  int64    `json:"skill3"`
	Type                    int64    `json:"type"`
}

type SquadList []Squad

func (s *Squad) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
