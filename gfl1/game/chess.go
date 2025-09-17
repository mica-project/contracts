package game

type ChessBuff struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	Description  string `json:"description"`
	MaxTier      int    `json:"max_tier"`
	Type         int    `json:"type"`
	DurationType int    `json:"duration_type"`
	Duration     int    `json:"duration"`
	Parameter    string `json:"parameter"`
	CreationId   int    `json:"creation_id"`
}

type ChessCampType struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Des          string `json:"des"`
	CampSkill    string `json:"camp_skill"`
	PassiveSkill string `json:"passive_skill"`
	Color        string `json:"color"`
	Code         string `json:"code"`
}

type ChessChip struct {
	Id                int    `json:"id"`
	Code              string `json:"code"`
	ChipUpExp         int    `json:"chip_up_exp"`
	ChipGroup         int    `json:"chip_group"`
	Level             int    `json:"level"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Rank              int    `json:"rank"`
	Number            int    `json:"number"`
	SkillTargetFirst  string `json:"skill_target_first"`
	TargetSelectFirst int    `json:"target_select_first"`
	Type              int    `json:"type"`
	Price             string `json:"price"`
	SellPrice         int    `json:"sell_price"`
	Experience        int    `json:"experience"`
	InitTime          int    `json:"init_time"`
	FitGunType        string `json:"fit_gun_type"`
}

type ChessChipTargetSelect struct {
	Id           int    `json:"id"`
	TargetType   string `json:"target_type"`
	TargetNumber int    `json:"target_number"`
	IsSelect     int    `json:"is_select"`
	Range        string `json:"range"`
	SelectType   int    `json:"select_type"`
	SelectOrder  string `json:"select_order"`
	SelectLimit  int    `json:"select_limit"`
}

type ChessChoiceStage struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	Parameter   string `json:"parameter"`
	InitTime    string `json:"init_time"`
	Cd          int    `json:"cd"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Probability int    `json:"probability"`
}

type ChessCreationLogic struct {
	PerformCreationId string `json:"perform_creation_id"`
}

type ChessCreationPerform struct {
	Id                   int    `json:"id"`
	Code                 string `json:"code"`
	StartType            int    `json:"start_type"`
	DestinationType      int    `json:"destination_type"`
	RouteType            int    `json:"route_type"`
	RouteHight           int    `json:"route_hight"`
	Speed                int    `json:"speed"`
	SpinVelocity         int    `json:"spin_velocity"`
	Duration             string `json:"duration"`
	Scale                string `json:"scale"`
	TriggerCreation      string `json:"trigger_creation"`
	TriggerCreationDelay string `json:"trigger_creation_delay"`
}

type ChessEnemy struct {
	Id             int    `json:"id"`
	Code           string `json:"code"`
	Name           string `json:"name"`
	Power          int    `json:"power"`
	Range          int    `json:"range"`
	Hp             int    `json:"hp"`
	Ap             int    `json:"ap"`
	AttackTimes    int    `json:"attack_times"`
	RewardChipIds  string `json:"reward_chip_ids"`
	SpineDirection int    `json:"spine_direction"`
	AttackAngle    int    `json:"attack_angle"`
	LifebarOffset  string `json:"lifebar_offset"`
}

type ChessGameConfig struct {
	Id             int    `json:"id"`
	ParameterName  string `json:"parameter_name"`
	ParameterType  string `json:"parameter_type"`
	ParameterValue string `json:"parameter_value"`
}

type ChessGachaReward struct {
	Id         int    `json:"id"`
	Gift       string `json:"gift"`
	Type       int    `json:"type"`
	TicketsNum int    `json:"tickets_num"`
	PrizeId    int    `json:"prize_id"`
	ItemIds    string `json:"item_ids"`
}
type ChessGunType struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Power          int    `json:"power"`
	Range          int    `json:"range"`
	Hp             int    `json:"hp"`
	Ap             int    `json:"ap"`
	AttackTimes    int    `json:"attack_times"`
	GunTypeSkill   string `json:"gun_type_skill"`
	AttackAngle    int    `json:"attack_angle"`
	AiPassiveSkill string `json:"ai_passive_skill"`
}
type ChessMap struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	UnlockItemId int    `json:"unlock_item_id"`
}

type ChessMission struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	ChessSpotids      string `json:"chess_spotids"`
	Rotation          int    `json:"rotation"`
	CameraHeightRange string `json:"camera_height_range"`
	CameraAngleH      int    `json:"camera_angle_h"`
	CameraAngleL      int    `json:"camera_angle_l"`
	MapLimit          string `json:"map_limit"`
	MapId             string `json:"map_id"`
	GlobalLimit       string `json:"global_limit"`
	GlobalPos         string `json:"global_pos"`
	MissionIcon       string `json:"mission_icon"`
}
type ChessModel struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	TeamDes         string `json:"team_des"`
	PlayerNum       int    `json:"player_num"`
	TeamNum         int    `json:"team_num"`
	TeamPlayerlimit string `json:"team_playerlimit"`
	Prize           string `json:"prize"`
}
type ChessRandomEnemy struct {
	Id             int    `json:"id"`
	EnemyId        string `json:"enemy_id"`
	RandomSpotId   string `json:"random_spot_id"`
	LaunchTimeType int    `json:"launch_time_type"`
	Time           int    `json:"time"`
}

type ChessRandomSpot struct {
	Id         int    `json:"id"`
	SpotEffect string `json:"spot_effect"`
}
type ChessScoreLevel struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Code          string `json:"code"`
	ScoreFloor    int    `json:"score_floor"`
	ScoreCeilling int    `json:"score_ceilling"`
	ScoreGet      string `json:"score_get"`
	Prize         int    `json:"prize"`
}
type ChessSeasonEvent struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	MissionId int    `json:"mission_id"`
}

type ChessSelectFrame struct {
	Id             int    `json:"id"`
	SelectNum      int    `json:"select_num"`
	ItemlanguageId string `json:"itemlanguage_id"`
	Code           string `json:"code"`
	Des            string `json:"des"`
	Preview        string `json:"preview"`
}
type ChessSkill struct {
	Id                  int    `json:"id"`
	CdType              int    `json:"cd_type"`
	Duration            string `json:"duration"`
	Trigger             string `json:"trigger"`
	TargetBuff          string `json:"target_buff"`
	RandomBuff          string `json:"random_buff"`
	NegativeSkillTarget int    `json:"negative_skill_target"`
	ChipUiControl       int    `json:"chip_ui_control"`
}

type ChessSkillTrigger struct {
	Id         int    `json:"id"`
	Type       string `json:"type"`
	Target     int    `json:"target"`
	Parameter  int    `json:"parameter"`
	Parameter2 string `json:"parameter2"`
}

type ChessSpot struct {
	Id                int    `json:"id"`
	ChessMissionId    int    `json:"chess_mission_id"`
	Type              int    `json:"type"`
	Neighbor          string `json:"neighbor"`
	PlayerOrder       int    `json:"player_order"`
	AxialCoordinatorQ int    `json:"axial_coordinator_q"`
	AxialCoordinatorR int    `json:"axial_coordinator_r"`
	PositiveDirection string `json:"positive_direction"`
	NegativeDirection string `json:"negative_direction"`
}

type ChessVoice struct {
	Id        int    `json:"id"`
	Situation string `json:"situation"`
	IsShow    int    `json:"is_show"`
	Code      string `json:"code"`
}
