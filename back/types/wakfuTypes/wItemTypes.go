package wakfuTypes

var WItemTypesData []WItemTypes

type itemTypeDefinition struct {
	Id                         int      `int:"id"`
	ParentId                   int      `int:"parentid"`
	EquipmentPositions         []string `string:"equipmentPositions"`
	EquipmentDisabledPositions []string `string:"equipmentDisabledPositions"`
	IsRecyclable               bool     `bool:"isRecyclable"`
	IsVisibleInAnimation       bool     `bool:"isVisibleInAnimation"`
}

type Title struct {
	Fr string `string:"fr"`
	En string `string:"en"`
	Es string `string:"es"`
	Pt string `string:"pt"`
}

type WItemTypes struct {
	Definition itemTypeDefinition
	Title      Title
}
