package main

type Resource string

const (
	Exp      Resource = "exp"
	Manpower Resource = "mp"
	Ammo     Resource = "ammo"
	MRE      Resource = "mre"
	Parts    Resource = "parts"
	Core     Resource = "core"
	Gem      Resource = "gem"
)

type Reward struct {
	ItemId int
	Amount int
}

type ResourceReward struct {
	Resource Resource
	Amount   int
}
