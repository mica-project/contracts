package game

type BreakoutDrop struct {
	Id         int    `json:"id"`
	DropItem   string `json:"drop_item"`
	DropWeight string `json:"drop_weight"`
	DropTimes  int    `json:"drop_times"`
}

type BreakoutEnemy struct {
	Id       int    `json:"id"`
	Code     string `json:"code"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Hp       int    `json:"hp"`
	Atk      int    `json:"atk"`
	Rate     int    `json:"rate"`
	Dodge    int    `json:"dodge"`
	Range    int    `json:"range"`
	Speed    int    `json:"speed"`
	Skills   string `json:"skills"`
	SkillDec string `json:"skill_dec"`
	BtTree   string `json:"bt_tree"`
}

type BreakoutGun struct {
	Id           int    `json:"id"`
	Code         string `json:"code"`
	Icon         string `json:"icon"`
	Name         string `json:"name"`
	Type         int    `json:"type"`
	CostDeploy   int    `json:"cost_deploy"`
	CostUpgrade  string `json:"cost_upgrade"`
	CostMre      int    `json:"cost_mre"`
	Hp           string `json:"hp"`
	Atk          string `json:"atk"`
	Rate         string `json:"rate"`
	Dodge        string `json:"dodge"`
	Armor        string `json:"armor"`
	Piercing     string `json:"piercing"`
	Range        string `json:"range"`
	Speed        string `json:"speed"`
	Shield       string `json:"shield"`
	Skills       string `json:"skills"`
	SkillDec     string `json:"skill_dec"`
	DeploymentCd int    `json:"deployment_cd"`
	BtTree       string `json:"bt_tree"`
}

type BreakoutItem struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Dec      string `json:"dec"`
	Type     int    `json:"type"`
	SoldPart int    `json:"sold_part"`
}

type BreakoutPhase struct {
	Id                 int    `json:"id"`
	Mission            int    `json:"mission"`
	Order              int    `json:"order"`
	Name               string `json:"name"`
	BreakoutGunInshop  string `json:"breakout_gun_inshop"`
	BreakoutItemInshop string `json:"breakout_item_inshop"`
	BreakoutMreReg     int    `json:"breakout_mre_reg"`
	Goal               string `json:"goal"`
	EnemyShow          string `json:"enemy_show"`
	Main               int    `json:"main"`
	BreakoutAmmoIni    int    `json:"breakout_ammo_ini"`
	BreakoutAmmoMax    int    `json:"breakout_ammo_max"`
	BreakoutAmmoReg    int    `json:"breakout_ammo_reg"`
	BreakoutApIni      int    `json:"breakout_ap_ini"`
	BreakoutApMax      int    `json:"breakout_ap_max"`
	BreakoutApReg      int    `json:"breakout_ap_reg"`
	Trigger            string `json:"trigger"`
	AvgStartService    string `json:"avg_start_service"`
	AvgEndService      string `json:"avg_end_service"`
	AvgBattle          string `json:"avg_battle"`
	Background         string `json:"background"`
	NpcTalk            string `json:"npc_talk"`
}

type BreakoutSkill struct {
	Id         int    `json:"id"`
	Action     string `json:"action"`
	Node       string `json:"node"`
	Effect     string `json:"effect"`
	Percentage string `json:"percentage"`
	Delay      string `json:"delay"`
	Special    string `json:"special"`
}

type BreakoutTalk struct {
	Id      int    `json:"id"`
	Trigger int    `json:"trigger"`
	Weight  int    `json:"weight"`
	Line    string `json:"line"`
}

type BreakoutTrigger struct {
	Id     int    `json:"id"`
	BtTree string `json:"bt_tree"`
}
