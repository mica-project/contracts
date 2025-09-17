package game

type CommanderClass struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	GroupId           string `json:"group_id"`
	Code              string `json:"code"`
	BonusPlastic      string `json:"bonus_plastic"`
	SkillId           int    `json:"skill_id"`
	Path              string `json:"path"`
	IsClass           int    `json:"is_class"`
	DesYear           string `json:"des_year"`
	Source            int    `json:"source"`
	SourceDescription string `json:"source_description"`
}
type CommanderEmoji struct {
	Id              int    `json:"id"`
	Content         string `json:"content"`
	IsShow          int    `json:"is_show"`
	RequiredComfort int    `json:"required_comfort"`
}

type CommanderRankingScores struct {
	Id          int     `json:"id"`
	Code        string  `json:"code"`
	TypeId      int     `json:"type_id"`
	BasicScores float64 `json:"basic_scores"`
	KSlopes     string  `json:"k_slopes"`
	XCounts     string  `json:"x_counts"`
}

type CommanderRankingTypes struct {
	Id      int    `json:"id"`
	ClassId int    `json:"class_id"`
	Title   string `json:"title"`
	Weight  int    `json:"weight"`
}

type CommanderUniform struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Code         string  `json:"code"`
	Icon         string  `json:"icon"`
	Type         int     `json:"type"`
	UniformClass string  `json:"uniform_class"`
	ColorNormal  string  `json:"color_normal"`
	Color0       string  `json:"color_0"`
	Color1       string  `json:"color_1"`
	Color2       string  `json:"color_2"`
	Color3       string  `json:"color_3"`
	Color4       string  `json:"color_4"`
	BoneName     string  `json:"bone_name"`
	Scale        float64 `json:"scale"`
	Position     string  `json:"position"`
	ColorIconId  string  `json:"color_icon_id"`
}
