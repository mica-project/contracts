package vehicle

import "encoding/json/v2"

type Vehicle struct {
	COff                    []float64 `json:"C_off"`
	HArmorPiercing          int64     `json:"H_armor_piercing"`
	HAttackRang             int64     `json:"H_attack_rang"`
	LArmorPiercing          int64     `json:"L_armor_piercing"`
	LAttackRang             int64     `json:"L_attack_rang"`
	AmmoPart                int64     `json:"ammo_part"`
	AssistType              int64     `json:"assist_type"`
	AssitSkill              int64     `json:"assit_skill"`
	Baseammo                int64     `json:"baseammo"`
	Basemre                 int64     `json:"basemre"`
	BasicArmor              int64     `json:"basic_armor"`
	BasicAtkSpeed           int64     `json:"basic_atk_speed"`
	BasicDodge              int64     `json:"basic_dodge"`
	BasicEffect             int64     `json:"basic_effect"`
	BasicHeavyDamage        int64     `json:"basic_heavy_damage"`
	BasicHit                int64     `json:"basic_hit"`
	BasicHp                 int64     `json:"basic_hp"`
	BasicLightDamage        int64     `json:"basic_light_damage"`
	BasicPrecision          int64     `json:"basic_precision"`
	BasicReload             int64     `json:"basic_reload"`
	BattleAssistRange       []int64   `json:"battle_assist_range"`
	Code                    string    `json:"code"`
	CommonComponentNum      int64     `json:"common_component_num"`
	CrewIds                 []int64   `json:"crew_ids"`
	Crit                    int64     `json:"crit"`
	CritDmg                 int64     `json:"crit_dmg"`
	DefBreak                int64     `json:"def_break"`
	DisplayAssistDamageArea int64     `json:"display_assist_damage_area"`
	DormInteract            string    `json:"dorm_interact"`
	EnName                  string    `json:"en_name"`
	EnergyBakMax            int64     `json:"energy_bak_max"`
	EnergyCost              int64     `json:"energy_cost"`
	EnergyMax               int64     `json:"energy_max"`
	Id                      int64     `json:"id"`
	Introduce               []string  `json:"introduce"`
	LaunchTimes             string    `json:"launch_times"`
	LongAttackRange         []int64   `json:"long_attack_range"`
	Move                    int64     `json:"move"`
	MrePart                 int64     `json:"mre_part"`
	Name                    string    `json:"name"`
	NameComponent1          string    `json:"name_component1"`
	NameComponent2          string    `json:"name_component2"`
	NameComponent3          string    `json:"name_component3"`
	NameComponent4          string    `json:"name_component4"`
	NameComponent5          string    `json:"name_component5"`
	NightVision             int64     `json:"night_vision"`
	NormalAttack            int64     `json:"normal_attack"`
	OriginalComponent1      int64     `json:"original_component1"`
	OriginalComponent2      int64     `json:"original_component2"`
	PerformanceSkill        int64     `json:"performance_skill"`
	Population              int64     `json:"population"`
	TalentSkill             []int64   `json:"talent_skill"`
	Type                    int64     `json:"type"`
	TypeComponent1          string    `json:"type_component1"`
	TypeComponent2          []int64   `json:"type_component2"`
	TypeComponent3          []int64   `json:"type_component3"`
	TypeComponent4          string    `json:"type_component4"`
	TypeComponent5          string    `json:"type_component5"`
	UnlockMission           []string  `json:"unlock_mission"`
}

type VehicleList []Vehicle

func (s *Vehicle) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
