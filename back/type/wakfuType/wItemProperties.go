package wakfuType

var WItemPropertiesData []WItemProperties

type WItemProperties struct {
	Id          int    `int:"id"`
	Name        string `string:"name"`
	Description string `string:"description"`
}
