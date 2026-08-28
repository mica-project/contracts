package misc

import "encoding/json/v2"

type FriendCosmetic struct {
	Code          string `json:"code"`
	DecomposeGift string `json:"decompose_gift"`
	FilterType    string `json:"filter_type"`
	Gem           int64  `json:"gem"`
	Id            int64  `json:"id"`
	InGasha       int64  `json:"in_gasha"`
	ItemId        int64  `json:"item_id"`
	ItemIds       string `json:"item_ids"`
	Onsale        int64  `json:"onsale"`
	Order         int64  `json:"order"`
	Rarity        int64  `json:"rarity"`
	SubType       int64  `json:"sub_type"`
}

type FriendCosmeticList []FriendCosmetic

func (s *FriendCosmetic) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
