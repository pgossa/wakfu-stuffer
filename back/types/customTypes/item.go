package customTypes

import "fmt"

var CustomItemsData []CustomItem

type SecondarySpecs struct {
	AccountBindType        int `int:"accountbindtype"`
	SetId                  int `int:"setid"`
	MinimumShardSlotNumber int `int:"maximumShardSlotNumber"`
	MaximumShardSlotNumber int `int:"maximumShardSlotNumber"`
}

type Usage struct {
	UseParameters      UseParameters
	UseEffects         []Effect
	UseCriticalEffects []Effect
}

type UseParameters struct {
	AP             int  `int:"AP"`
	MP             int  `int:"MP"`
	WP             int  `int:"WP"`
	MinRange       int  `int:"minRange"`
	MaxRange       int  `int:"maxRange"`
	OnFreeCell     bool `bool:"onFreeCell"` // Not used ATM (always false)
	OnLos          bool `bool:"onLos"`      // DK but used
	OnlyLine       bool `bool:"onlyLine"`
	NoBorderCell   bool `bool:"noBorderCell"`  // DK but used
	UseWorldTarget int  `int:"useWorldTarget"` // Not used ATM (always 0)
}

type Effect struct {
	Description MultiLang
	ActionId    int     `int:"actionId"`
	Quantity    float64 `float64:"quantity"`
	Number      float64 `float64:"omitempty number"` // Used for: Give Quantity in Number of Description
	AreaShape   int     `int:"areaShape"`
	AreaSize    []int   `int:"areaSize"`
}

type MultiLang struct {
	Fr string `string:"fr"`
	En string `string:"en"`
	Es string `string:"es"`
	Pt string `string:"pt"`
}

type EquipmentPosition struct {
	Title            MultiLang
	Position         []string `string:"position"`
	PositionDisabled []string `string:"positionDisabled"`
}

type CustomItem struct {
	Title             MultiLang
	Description       MultiLang
	Id                int `int:"id"`
	Level             int `int:"level"`
	Rarity            int `int:"rarity"`
	ItemTypeId        int `int:"itemTypeId"`
	EquipmentPosition EquipmentPosition
	SecondarySpecs    SecondarySpecs
	Usage             Usage
	EquipEffects      []Effect
}

func (item CustomItem) ToString() string {
	resString := ""
	resString += item.Title.En + " "
	resString += fmt.Sprint(item.Level) + " "
	resString += fmt.Sprint(item.Rarity) + " "
	resString += item.EquipmentPosition.Title.En
	return resString
}
