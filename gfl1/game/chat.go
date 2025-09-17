package game

type ChatChannel struct {
	Id             int `json:"id"`
	IfInput        int `json:"if_input"`
	IsFixedPhrases int `json:"is_fixed_phrases"`
}

type ChatFixPhrases struct {
	Id      int    `json:"id"`
	Group   int    `json:"group"`
	Content string `json:"content"`
	Type    int    `json:"type"`
}
