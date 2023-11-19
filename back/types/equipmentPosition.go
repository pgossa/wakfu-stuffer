package types

import "github.com/pgossa/wakfu-stuffer/types/customTypes"

func GetEquipmentByPosition(position string) []customTypes.CustomItem {
	resList := []customTypes.CustomItem{}
	for _, item := range customTypes.CustomItemsData {
		for _, loopPosition := range item.EquipmentPosition.Position {
			if loopPosition == position {
				resList = append(resList, item)
			}
		}
	}
	return resList
}
