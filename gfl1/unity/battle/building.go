package battle

import "encoding/json/v2"

type Building struct {
	ActiveBuildingInfo    string   `json:"active_building_info"`
	BattleAssistRange     []int64  `json:"battle_assist_range"`
	BattleSkill           string   `json:"battle_skill"`
	Belong                int64    `json:"belong"`
	BelongColor           []string `json:"belong_color"`
	Code                  string   `json:"code"`
	Condition             []string `json:"condition"`
	ConfrontDes           string   `json:"confront_des"`
	Defender              int64    `json:"defender"`
	DefenderUpper         int64    `json:"defender_upper"`
	Des                   []string `json:"des"`
	HoldBelong            int64    `json:"hold_belong"`
	Id                    int64    `json:"id"`
	IsDestroy             []int64  `json:"is_destroy"`
	MissionSkill          string   `json:"mission_skill"`
	Name                  string   `json:"name"`
	PerformanceSkill      []string `json:"performance_skill"`
	ShiftingSpot          []int64  `json:"shifting_spot"`
	ShiftingTeam          []int64  `json:"shifting_team"`
	ShowInfo              int64    `json:"show_info"`
	Type                  int64    `json:"type"`
	WorkingSpotActivation string   `json:"working_spot_activation"`
}

type BuildingList []Building

func (s *Building) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
