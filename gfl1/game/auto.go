package game

type AutoFormation struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	RecoLocation string `json:"reco_location"`
	RecoGun1     string `json:"reco_gun_1"`
	RecoGun2     string `json:"reco_gun_2"`
	RecoGun5     string `json:"reco_gun_5"`
	RecoGun3     string `json:"reco_gun_3"`
	RecoGun6     string `json:"reco_gun_6"`
	TeamTagIds   string `json:"team_tag_ids"`
}

type AutoFormationSangvis struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TeamInfo    string `json:"team_info"`
	FrontTagIds string `json:"front_tag_ids"`
	BackTagIds  string `json:"back_tag_ids"`
	TeamTagIds  string `json:"team_tag_ids"`
}

type AutoMission struct {
	MissionID      int    `json:"mission_id"`
	TeamEffect     int    `json:"team_effect"`
	MonthTeamCount int    `json:"month_team_count"`
	TeamCount      int    `json:"team_count"`
	Mp             int    `json:"mp"`
	Ammo           int    `json:"ammo"`
	Mre            int    `json:"mre"`
	Part           int    `json:"part"`
	Duration       int    `json:"duration"`
	Experience     int    `json:"experience"`
	ExpectGunLevel int    `json:"expect_gun_level"`
	GetGunNum      int    `json:"get_gun_num"`
	GunNPool       string `json:"gun_n_pool"`
	Gun1Pool       string `json:"gun_1_pool"`
}
