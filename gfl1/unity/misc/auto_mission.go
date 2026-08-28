package misc

import "encoding/json/v2"

type AutoMission struct {
	Ammo           int64   `json:"ammo"`
	Duration       int64   `json:"duration"`
	ExpectGunLevel int64   `json:"expect_gun_level"`
	Experience     int64   `json:"experience"`
	GetGunNum      int64   `json:"get_gun_num"`
	Gun1Pool       []int64 `json:"gun_1_pool"`
	GunNPool       []int64 `json:"gun_n_pool"`
	MissionId      int64   `json:"mission_id"`
	MonthTeamCount int64   `json:"month_team_count"`
	Mp             int64   `json:"mp"`
	Mre            int64   `json:"mre"`
	Part           int64   `json:"part"`
	TeamCount      int64   `json:"team_count"`
	TeamEffect     int64   `json:"team_effect"`
}

type AutoMissionList []AutoMission

func (s *AutoMission) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
