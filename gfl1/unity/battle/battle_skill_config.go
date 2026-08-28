package battle

import "encoding/json/v2"

type BattleSkillConfig struct {
	ActionId           int64    `json:"action_id"`
	BuffIdSelf         []string `json:"buff_id_self"`
	CdTime             int64    `json:"cd_time"`
	CdType             int64    `json:"cd_type"`
	Code               string   `json:"code"`
	DataPool1          string   `json:"data_pool_1"`
	DataPool2          string   `json:"data_pool_2"`
	Description        string   `json:"description"`
	Id                 int64    `json:"id"`
	IsCdr              int64    `json:"is_cdr"`
	IsFormAction       int64    `json:"is_form_action"`
	IsSwitch           int64    `json:"is_switch"`
	Level              int64    `json:"level"`
	LvupDescription    []string `json:"lvup_description"`
	Name               string   `json:"name"`
	SkillGroupId       int64    `json:"skill_group_id"`
	SkillLvCall        int64    `json:"skill_lv_call"`
	SkillPriority      int64    `json:"skill_priority"`
	SkinName           string   `json:"skin_name"`
	SpecialBuffTrigger string   `json:"special_buff_trigger"`
	StartCdTime        int64    `json:"start_cd_time"`
	TargetSelectAi     int64    `json:"target_select_ai"`
	TrainCoinType      int64    `json:"train_coin_type"`
	TriggerTarget      int64    `json:"trigger_target"`
	TriggerType        int64    `json:"trigger_type"`
	Type               int64    `json:"type"`
}

type BattleSkillConfigList []BattleSkillConfig

func (s *BattleSkillConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
