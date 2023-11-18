package wakfuType

var WItemData []WItem

type ItemDefinition struct {
	Item               ItemItem
	UseEffects         []Effect
	UseCriticalEffects []Effect
	EquipEffects       []Effect
}
type BaseParameters struct {
	ItemTypeId             int `int:"itemTypeId"`
	ItemSetId              int `int:"itemSetId"`
	Rarity                 int `int:"rarity"`
	BindType               int `int:"bindType"`
	MinimumShardSlotNumber int `int:"minimumShardSlotNumber"`
	MaximumShardSlotNumber int `int:"maximumShardSlotNumber"`
}

type UseParameters struct {
	UseCostMp           int  `int:"useCostMp"`
	UseCostWp           int  `int:"useCostWp"`
	UseCostAp           int  `int:"useCostAp"`
	UseRangeMin         int  `int:"useRangeMin"`
	UseRangeMax         int  `int:"useRangeMax"`
	UseTestFreeCell     bool `bool:"useTestFreeCell"`
	UseTestLos          bool `bool:"useTestLos"`
	UseTestOnlyLine     bool `bool:"useTestOnlyLine"`
	UseTestNoBorderCell bool `bool:"useTestNoBorderCell"`
	UseWorldTarget      int  `int:"useWorldTarget"`
}

type GraphicParameters struct {
	GfxId       int `int:"gfxId"`
	FemaleGfxId int `int:"femaleGfxId"`
}

type ItemItem struct {
	Id                int `int:"id"`
	Level             int `int:"level"`
	BaseParameters    BaseParameters
	UseParameters     UseParameters
	GraphicParameters GraphicParameters
	Properties        []int
}

type DefinitionDefinition struct {
	Id        int       `int:"id"`
	ActionId  int       `int:"actionId"`
	AreaShape int       `int:"areaShape"`
	AreaSize  []int     `int:"areaSize"`
	Params    []float64 `float64:"params"`
}

type EffectDefinition struct {
	Definition DefinitionDefinition
}

type Effect struct {
	Effect EffectDefinition
}

type WItem struct {
	Definition  ItemDefinition
	Title       Title
	Description Description
}
