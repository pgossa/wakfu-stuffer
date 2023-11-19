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
	AP                    int
	MP                    int
	WP                    int
	PO                    int
	ArmorGiven            int
	ArmorReceived         int
	ElementaryDamages     Elementary
	ElementaryResistances Elementary
	CriticalChance        int
	Block                 int
	Initiative            int
	Dodge                 int
	Lock                  int
	Prospecting           int
	Wisdom                int
	Control               int
	Will                  int
	MCritical             int
	RCritical             int
	MBack                 int
	RBack                 int
	MMeler                int
	MDistance             int
	MHeal                 int
	MBerzerk              int
}

type Elementary struct {
	Fire  int
	Earth int
	Water int
	Air   int
}
