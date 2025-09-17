package game

type CoffeeshopComic struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cost        string `json:"cost"`
	PrizeId     string `json:"prize_id"`
}

type CoffeeshopPv struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cost        string `json:"cost"`
	PrizeId     string `json:"prize_id"`
}
