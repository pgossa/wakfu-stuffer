package wakfuType

var WActionData []WAction

type actionDefinition struct {
	Id     int    `int:"id"`
	Effect string `string:"effect"`
}

type Description struct {
	Fr string `string:"fr"`
	En string `string:"en"`
	Es string `string:"es"`
	Pt string `string:"pt"`
}

type WAction struct {
	Definition  actionDefinition
	Description Description
}
