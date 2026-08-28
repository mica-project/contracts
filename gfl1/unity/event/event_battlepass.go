package event

import "encoding/json/v2"

type EventBattlepass struct {
	DiscountItemid        int64    `json:"discount_itemid"`
	EndTime               string   `json:"end_time"`
	ExhibitBonus          []int64  `json:"exhibit_bonus"`
	ExpMallid             int64    `json:"exp_mallid"`
	GiftExtra             int64    `json:"gift_extra"`
	GiftFree              []string `json:"gift_free"`
	GiftPaytounlock       []string `json:"gift_paytounlock"`
	GiftPlus              string   `json:"gift_plus"`
	Id                    int64    `json:"id"`
	Name                  string   `json:"name"`
	PaidProductid         []string `json:"paid_productid"`
	PaidProductidDiscount []string `json:"paid_productid_discount"`
	PtItemId              int64    `json:"pt_item_id"`
	SkinId                string   `json:"skin_id"`
	StartTime             string   `json:"start_time"`
	UnlockItemid          string   `json:"unlock_itemid"`
}

type EventBattlepassList []EventBattlepass

func (s *EventBattlepass) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
