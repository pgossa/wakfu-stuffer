package types

import "github.com/pgossa/wakfu-stuffer/types/customTypes"

type RequestRanking struct {
	Level          int
	WeightSpec     WeightSpec
	ForbiddenItems []customTypes.CustomItem
	MandatoryItems []customTypes.CustomItem
}

type WeightSpec struct {
	HP                    int
	HPWeight              int
	AP                    int
	APWeight              int
	MP                    int
	MPWeight              int
	WP                    int
	WPWeight              int
	PO                    int
	POWeight              int
	ArmorGiven            int
	ArmorGivenWeight      int
	ArmorReceived         int
	ArmorReceivedWeight   int
	ElementaryDamages     Elementary
	ElementaryResistances Elementary
	CriticalChance        int
	CriticalChanceWeight  int
	Block                 int
	BlockWeight           int
	Initiative            int
	InitiativeWeight      int
	Dodge                 int
	DodgeWeight           int
	Lock                  int
	LockWeight            int
	Prospecting           int
	ProspectingWeight     int
	Wisdom                int
	WisdomWeight          int
	Control               int
	ControlWeight         int
	Will                  int
	WillWeight            int
	MCritical             int
	MCriticalWeight       int
	RCritical             int
	RCriticalWeight       int
	MBack                 int
	MBackWeight           int
	RBack                 int
	RBackWeight           int
	MMelee                int
	MMeleeWeight          int
	MDistance             int
	MDistanceWeight       int
	MHeal                 int
	MHealWeight           int
	MBerzerk              int
	MBerzerkWeight        int
}

type Elementary struct {
	Fire        int
	FireWeight  int
	Earth       int
	EarthWeight int
	Water       int
	WaterWeight int
	Air         int
	AirWeight   int
}
