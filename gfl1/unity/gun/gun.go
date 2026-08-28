package gun

import "encoding/json/v2"

type Gun struct {
	Ai                 int64    `json:"ai"`
	AmmoAddWithnumber  int64    `json:"ammo_add_withnumber"`
	ArmorPiercing      int64    `json:"armor_piercing"`
	AtkSpeedMax        int64    `json:"atk_speed_max"`
	Baseammo           int64    `json:"baseammo"`
	Basemre            int64    `json:"basemre"`
	Code               string   `json:"code"`
	Crit               int64    `json:"crit"`
	DevelopDuration    int64    `json:"develop_duration"`
	Dialogue           []string `json:"dialogue"`
	EatRatio           int64    `json:"eat_ratio"`
	EffectGridCenter   int64    `json:"effect_grid_center"`
	EffectGridEffect   []string `json:"effect_grid_effect"`
	EffectGridPos      []int64  `json:"effect_grid_pos"`
	EnIntroduce        []string `json:"en_introduce"`
	EnName             string   `json:"en_name"`
	ExploreTag         string   `json:"explore_tag"`
	Extra              []string `json:"extra"`
	GunDetailBg        string   `json:"gun_detail_bg"`
	GunTagIds          string   `json:"gun_tag_ids"`
	Id                 int64    `json:"id"`
	Introduce          []string `json:"introduce"`
	LaunchTime         string   `json:"launch_time"`
	MindupdateConsume  []string `json:"mindupdate_consume"`
	MreAddWithnumber   int64    `json:"mre_add_withnumber"`
	Name               string   `json:"name"`
	NormalAttack       int64    `json:"normal_attack"`
	ObtainIds          []int64  `json:"obtain_ids"`
	OrgId              int64    `json:"org_id"`
	PassiveSkill       string   `json:"passive_skill"`
	Rank               int64    `json:"rank"`
	RankDisplay        int64    `json:"rank_display"`
	RatioDodge         int64    `json:"ratio_dodge"`
	RatioHit           int64    `json:"ratio_hit"`
	RatioLife          int64    `json:"ratio_life"`
	RatioPow           int64    `json:"ratio_pow"`
	RatioRange         int64    `json:"ratio_range"`
	RatioRate          int64    `json:"ratio_rate"`
	RatioRec           int64    `json:"ratio_rec"`
	RatioSpeed         int64    `json:"ratio_speed"`
	RecommendedTeamIds []int64  `json:"recommended_team_ids"`
	RelatedStoryId     int64    `json:"related_story_id"`
	Retireammo         int64    `json:"retireammo"`
	Retiremp           int64    `json:"retiremp"`
	Retiremre          int64    `json:"retiremre"`
	Skill1             int64    `json:"skill1"`
	Type               int64    `json:"type"`
	TypeEquip1         []string `json:"type_equip1"`
	TypeEquip2         string   `json:"type_equip2"`
	TypeEquip3         string   `json:"type_equip3"`
}

type GunList []Gun

func (s *Gun) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
