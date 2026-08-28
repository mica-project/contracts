package fairy

import "encoding/json/v2"

type Fairy struct {
	Ai                int64     `json:"ai"`
	Armor             int64     `json:"armor"`
	AvatarOffset      []float64 `json:"avatar_offset"`
	AvatarScale       []float64 `json:"avatar_scale"`
	Category          int64     `json:"category"`
	Code              string    `json:"code"`
	Description       []string  `json:"description"`
	DevelopDuration   int64     `json:"develop_duration"`
	Dodge             int64     `json:"dodge"`
	Grow              int64     `json:"grow"`
	Hit               int64     `json:"hit"`
	Id                int64     `json:"id"`
	Introduce         string    `json:"introduce"`
	ModExp            int64     `json:"mod_exp"`
	Name              string    `json:"name"`
	ObtainIds         string    `json:"obtain_ids"`
	OrgId             int64     `json:"org_id"`
	PictureOffset     []int64   `json:"picture_offset"`
	PictureScale      []float64 `json:"picture_scale"`
	Pow               int64     `json:"pow"`
	PowerupAmmo       int64     `json:"powerup_ammo"`
	PowerupMp         int64     `json:"powerup_mp"`
	PowerupMre        int64     `json:"powerup_mre"`
	PowerupPart       int64     `json:"powerup_part"`
	Proportion        []string  `json:"proportion"`
	QualityExp        int64     `json:"quality_exp"`
	QualityNeedNumber []string  `json:"quality_need_number"`
	Retireammo        int64     `json:"retireammo"`
	Retiremp          int64     `json:"retiremp"`
	Retiremre         int64     `json:"retiremre"`
	Retirepart        int64     `json:"retirepart"`
	SkillId           string    `json:"skill_id"`
	Type              int64     `json:"type"`
}

type FairyList []Fairy

func (s *Fairy) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
