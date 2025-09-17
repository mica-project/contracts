package game

type Daily struct {
	Identity        int    `json:"identity"`
	MissionType     int    `json:"mission_type"`
	Type            string `json:"type"`
	Count           int    `json:"count"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	PrizeId         int    `json:"prize_id"`
	ExchangePrizeId int    `json:"exchange_prize_id"`
}

type DailyEvent struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Hint    string `json:"hint"`
	Script  string `json:"script"`
	Options string `json:"options"`
	Notes   string `json:"notes"`
}
type DailyEventOption struct {
	Id                      int    `json:"id"`
	FightEnvironmentSkillId int    `json:"fight_environment_skill_id"`
	Description             string `json:"description"`
	BalanceOut              int    `json:"balance_out"`
	RoundDuration           int    `json:"round_duration"`
	BattleDuration          int    `json:"battle_duration"`
}

type DailyExmissionGroup struct {
	Id           int `json:"id"`
	MissionId    int `json:"mission_id"`
	StateType    int `json:"state_type"`
	VehicleId    int `json:"vehicle_id"`
	Effect       int `json:"effect"`
	MissionFloor int `json:"mission_floor"`
}

type DailyGunequipObtain struct {
	Id             int    `json:"id"`
	Type           int    `json:"type"`
	EquipOrGunId   int    `json:"equip_or_gun_id"`
	GetType        int    `json:"get_type"`
	DropDifficulty string `json:"drop_difficulty"`
}
type DailyMap struct {
	Id         int    `json:"id"`
	Coordinate string `json:"coordinate"`
	Type       string `json:"type"`
	Neighbor   string `json:"neighbor"`
	MapLabel   int    `json:"map_label"`
}
type DailyMapLight struct {
	Id                        int     `json:"id"`
	SeasonCode                int     `json:"season_code"`
	LightColor                string  `json:"light_color"`
	LightPos                  string  `json:"light_pos"`
	LightDir                  string  `json:"light_dir"`
	LightAngle                int     `json:"light_angle"`
	LightBrightness           float64 `json:"light_brightness"`
	BrightContrastRatio       int     `json:"bright_contrast_ratio"`
	DarkContrastRatio         float64 `json:"dark_contrast_ratio"`
	LightTexBrightness        float64 `json:"light_tex_brightness"`
	DirLightShadowColorLock   string  `json:"dir_light_shadow_colorLock"`
	ModelTexColor             string  `json:"model_tex_color"`
	StateCode                 int     `json:"state_code"`
	DirLightShadowColorUnLock string  `json:"dir_light_shadow_colorUnLock"`
}
type DailyMapModel struct {
	Id              int    `json:"id"`
	Type            int    `json:"type"`
	HeightRegion    string `json:"height_region"`
	DifficultyGroup string `json:"difficulty_group"`
	Code            string `json:"code"`
	IconDeviation   string `json:"icon_deviation"`
}
type DailyMapRoute struct {
	MapId    int    `json:"map_id"`
	Route    string `json:"route"`
	MapLabel int    `json:"map_label"`
}

type DailyMapSpot struct {
	Id          int    `json:"id"`
	MapId       int    `json:"map_id"`
	SpotId      int    `json:"spot_id"`
	SpotType    int    `json:"spot_type"`
	ModelCode   string `json:"model_code"`
	ModelHeight string `json:"model_height"`
}

type DailyMissionDifficultyGroup struct {
	Id                        int    `json:"id"`
	Name                      string `json:"name"`
	RecommendGunLevel         int    `json:"recommend_gun_level"`
	RecommendFairyLevel       string `json:"recommend_fairy_level"`
	RecommendFairyTip         string `json:"recommend_fairy_tip"`
	RecommendFairies          string `json:"recommend_fairies"`
	RecommendSquadLevel       string `json:"recommend_squad_level"`
	RecommendSquadTip         string `json:"recommend_squad_tip"`
	RecommendSquads           string `json:"recommend_squads"`
	RecommendedVehicles       string `json:"recommended_vehicles"`
	RecommendedVehicleTip     string `json:"recommended_vehicle_tip"`
	RecommendedVehiclesEffect string `json:"recommended_vehicles_effect"`
	DefaultEffects            string `json:"default_effects"`
	DefaultBg                 string `json:"default_bg"`
}

type DailyMissionGroup struct {
	Id              int    `json:"id"`
	Type            int    `json:"type"`
	Difficulty      int    `json:"difficulty"`
	DifficultyGroup int    `json:"difficulty_group"`
	StateType       int    `json:"state_type"`
	MissionGroup    string `json:"mission_group"`
	DailyPt         int    `json:"daily_pt"`
}

type DailyNote struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Code  string `json:"code"`
}
type DailyPass struct {
	Id             int    `json:"id"`
	Type           int    `json:"type"`
	Pt             string `json:"pt"`
	Gift           string `json:"gift"`
	Gift2          string `json:"gift2"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	GiftBackground string `json:"gift_background"`
	Item           int    `json:"item"`
}
