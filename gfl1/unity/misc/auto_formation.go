package misc

import "encoding/json/v2"

type AutoFormation struct {
	Description  []string `json:"description"`
	Id           int64    `json:"id"`
	Name         string   `json:"name"`
	RecoEquip1   []string `json:"reco_equip_1"`
	RecoEquip2   []string `json:"reco_equip_2"`
	RecoEquip3   []string `json:"reco_equip_3"`
	RecoEquip5   []string `json:"reco_equip_5"`
	RecoEquip6   []string `json:"reco_equip_6"`
	RecoGun1     string   `json:"reco_gun_1"`
	RecoGun2     string   `json:"reco_gun_2"`
	RecoGun3     string   `json:"reco_gun_3"`
	RecoGun5     string   `json:"reco_gun_5"`
	RecoGun6     string   `json:"reco_gun_6"`
	RecoLocation []string `json:"reco_location"`
	TeamTagIds   []int64  `json:"team_tag_ids"`
}

type AutoFormationList []AutoFormation

func (s *AutoFormation) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
