package breakout

import "encoding/json/v2"

type BreakoutPhase struct {
	AvgBattle          string   `json:"avg_battle"`
	AvgEndService      string   `json:"avg_end_service"`
	AvgStartService    string   `json:"avg_start_service"`
	Background         string   `json:"background"`
	BreakoutAmmoMax    int64    `json:"breakout_ammo_max"`
	BreakoutAmmoReg    int64    `json:"breakout_ammo_reg"`
	BreakoutApIni      int64    `json:"breakout_ap_ini"`
	BreakoutApMax      int64    `json:"breakout_ap_max"`
	BreakoutApReg      int64    `json:"breakout_ap_reg"`
	BreakoutGunInshop  []int64  `json:"breakout_gun_inshop"`
	BreakoutItemInshop []string `json:"breakout_item_inshop"`
	BreakoutMreReg     int64    `json:"breakout_mre_reg"`
	DropMissionkey     string   `json:"drop_missionkey"`
	EnemyShow          []int64  `json:"enemy_show"`
	Goal               string   `json:"goal"`
	Id                 int64    `json:"id"`
	ItemDrop           string   `json:"item_drop"`
	Main               int64    `json:"main"`
	Mission            int64    `json:"mission"`
	Name               string   `json:"name"`
	NpcTalk            []int64  `json:"npc_talk"`
	Order              int64    `json:"order"`
	Trigger            string   `json:"trigger"`
}

type BreakoutPhaseList []BreakoutPhase

func (s *BreakoutPhase) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
