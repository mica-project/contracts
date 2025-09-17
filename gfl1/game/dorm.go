package game

type DormAction struct {
	Id                int    `json:"id"`
	SpineName         string `json:"spine_name"`
	InteractType      int    `json:"interact_type"`
	InteractPointType string `json:"interact_point_type"`
	ShadowExist       int    `json:"shadow_exist"`
}
type DormAi struct {
	Id         int    `json:"id"`
	Actions    string `json:"actions"`
	UpRate     string `json:"up_rate"`
	MinTime    string `json:"min_time"`
	TimeWeight string `json:"time_weight"`
}

type DormEmojiTextInfo struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}
