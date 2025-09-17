package game

type BattleActionConfig struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	ActionOrder     string `json:"action_order"`
	ActionPlaySpeed string `json:"action_playspeed"`
	CreationOrder   string `json:"creation_order"`
}

type BattleBuff struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	ConflictType  int    `json:"conflict_type"`
	MaxTier       int    `json:"max_tier"`
	Type          int    `json:"type"`
	DurationType  int    `json:"duration_type"`
	Duration      string `json:"duration"`
	BossAvailable int    `json:"boss_available"`
}

type BattleCreation struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Code              string `json:"code"`
	StartType         string `json:"start_type"`
	DestinationType   string `json:"destination_type"`
	IsFormOffset      int    `json:"is_form_offset"`
	IsFormOffsetStart int    `json:"is_form_offset_start"`
	RouteType         int    `json:"route_type"`
	Speed             int    `json:"speed"`
	EffectArea        int    `json:"effect_area"`
	HurtID            string `json:"hurt_id"`
	IsFormPlay        int    `json:"is_form_play"`
	Scale             string `json:"scale"`
	SoundOrder        string `json:"sound_order"`
	IsOrbEffect       int    `json:"is_orb_effect"`
}

type BattleFormula struct {
	Id      int    `json:"id"`
	Formula string `json:"formula"`
}

type BattleHurtConfig struct {
	Id              int    `json:"id"`
	Description     string `json:"description"`
	DamageRatio     string `json:"damage_ratio"`
	IsCriticalHit   int    `json:"is_critical_hit"`
	CriticalHitRate string `json:"critical_hit_rate"`
	FloatWide       int    `json:"float_wide"`
	IsMiss          int    `json:"is_miss"`
	IsArmor         int    `json:"is_armor"`
	HitType         int    `json:"hit_type"`
}

type BattleMovementInfo struct {
	Id          string `json:"id"`
	Target      string `json:"target"`
	Offset      string `json:"offset"`
	Distance    string `json:"distance"`
	Duration    string `json:"duration"`
	Cd          string `json:"cd"`
	IsCharacter string `json:"is_character"`
}

type BattleSkillConfig struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	SkillGroupId       int    `json:"skill_group_id"`
	Level              int    `json:"level"`
	Type               int    `json:"type"`
	SkillPriority      int    `json:"skill_priority"`
	CdType             int    `json:"cd_type"`
	CdTime             int    `json:"cd_time"`
	StartCdTime        int    `json:"start_cd_time"`
	TriggerType        int    `json:"trigger_type"`
	TriggerTarget      int    `json:"trigger_target"`
	TargetSelectAi     int    `json:"target_select_ai"`
	ActionId           int    `json:"action_id"`
	IsFormAction       int    `json:"is_form_action"`
	SkinName           string `json:"skin_name"`
	BuffIdSelf         string `json:"buff_id_self"`
	Description        string `json:"description"`
	LvupDescription    string `json:"lvup_description"`
	DataPool1          string `json:"data_pool_1"`
	DataPool2          string `json:"data_pool_2"`
	Code               string `json:"code"`
	TrainCoinType      int    `json:"train_coin_type"`
	IsSwitch           int    `json:"is_switch"`
	IsCdr              int    `json:"is_cdr"`
	SkillLvCall        int    `json:"skill_lv_call"`
	SpecialBuffTrigger string `json:"special_buff_trigger"`
}

type BattleSkillTypeConfig struct {
	Id              int `json:"id"`
	ChargeTime      int `json:"charge_time"`
	ChargeTier      int `json:"charge_tier"`
	StartChargeTier int `json:"start_charge_tier"`
}

type BattleTargetSelectAi struct {
	Id              int    `json:"id"`
	Description     string `json:"description"`
	TargetType      int    `json:"target_type"`
	TargetNumber    int    `json:"target_number"`
	IsAscendOrder   string `json:"is_ascend_order"`
	ListRanking     int    `json:"list_ranking"`
	TargetParameter int    `json:"target_parameter"`
	BuffTriggerType int    `json:"buff_trigger_type"`
}

type BattleTrigger struct {
	Id        int    `json:"id"`
	Type      int    `json:"type"`
	Target    int    `json:"target"`
	Parameter string `json:"parameter"`
	BuffId    int    `json:"buff_id"`
}

type BattleWatch struct {
	Id                     int    `json:"id"`
	Description            string `json:"description"`
	RecordChara1           int    `json:"record_chara1"`
	WatchTriggerType       int    `json:"watch_trigger_type"`
	ActivationWatchTrigger int    `json:"activation_watch_trigger"`
	ActivationSkills       string `json:"activation_skills"`
	NewRecordChara1        int    `json:"new_record_chara1"`
}

type BattleWatchTrigger struct {
	Id                  int    `json:"id"`
	Description         string `json:"description"`
	WatchedSkillType    int    `json:"watched_skill_type"`
	WatchedSkillGroupId int    `json:"watched_skill_group_id"`
	HurtType            string `json:"hurt_type"`
	HurtJudge           int    `json:"hurt_judge"`
	BuffJudge           int    `json:"buff_judge"`
	IsHurtBody          int    `json:"is_hurt_body"`
}
