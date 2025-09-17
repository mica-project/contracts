package game

type Achievement struct {
	Identity int    `json:"identity"`
	Type     string `json:"type"`
	Count    int    `json:"count"`
	UserExp  int    `json:"user_exp"`
	Mp       int    `json:"mp"`
	Ammo     int    `json:"ammo"`
	Mre      int    `json:"mre"`
	Part     int    `json:"part"`
	Core     int    `json:"core"`
	ItemIds  string `json:"item_ids"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	TypeSort int    `json:"type_sort"`
	Sort     int    `json:"sort"`
	IconCode string `json:"icon_code"`
}

type AdjutantSkin struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	ObtainTxt string `json:"obtain_txt"`
	PosInfo   string `json:"pos_info"`
}

type AllyTeam struct {
	ID             int    `json:"id"`
	Code           string `json:"code"`
	UIImageIcon    string `json:"ui_image_icon"`
	Name           string `json:"name"`
	Guns           string `json:"guns"`
	EnemyTeamID    int    `json:"enemy_team_id"`
	InitialType    int    `json:"initial_type"`
	Ai             string `json:"ai"`
	AiContent      string `json:"ai_content"`
	EnemyPanelType int    `json:"enemy_panel_type"`
	NoBattleDamage int    `json:"no_battle_damage"`
}

type AttendanceInfo struct {
	Day            string `json:"day"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	AttendanceType string `json:"attendance_type"`
	Mp             string `json:"mp"`
}

type BondageInfo struct {
	Id        int    `json:"id"`
	MainId    string `json:"main_id"`
	SubId     string `json:"sub_id"`
	LinesCode string `json:"lines_code"`
	LinesTxt  string `json:"lines_txt"`
}

type Building struct {
	Id                    int    `json:"id"`
	Defender              int    `json:"defender"`
	DefenderUpper         int    `json:"defender_upper"`
	Belong                int    `json:"belong"`
	HoldBelong            int    `json:"hold_belong"`
	MissionSkill          string `json:"mission_skill"`
	BattleSkill           string `json:"battle_skill"`
	Condition             string `json:"condition"`
	IsDestroy             string `json:"is_destroy"`
	Name                  string `json:"name"`
	Code                  string `json:"code"`
	ShiftingSpot          string `json:"shifting_spot"`
	ShiftingTeam          string `json:"shifting_team"`
	BattleAssistRange     string `json:"battle_assist_range"`
	PerformanceSkill      string `json:"performance_skill"`
	ShowInfo              int    `json:"show_info"`
	Type                  int    `json:"type"`
	BelongColor           string `json:"belong_color"`
	WorkingSpotActivation string `json:"working_spot_activation"`
	ActiveBuildingInfo    string `json:"active_building_info"`
	Des                   string `json:"des"`
	ConfrontDes           string `json:"confront_des"`
}
