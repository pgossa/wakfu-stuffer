package utils

import (
	"github.com/pgossa/wakfu-stuffer/types"
	"github.com/pgossa/wakfu-stuffer/types/customTypes"
)

func FilterEquipementsByRarity(equipmentList []customTypes.CustomItem, rarity string) []customTypes.CustomItem {
	resList := []customTypes.CustomItem{}
	for _, equipment := range equipmentList {
		if equipment.Rarity == types.Rarity.GetRarityByName(rarity) {
			resList = append(resList, equipment)
		}
	}
	return resList
}

func RemoveTooHighLevelItems(level int) customTypes.EquipableItems {
	items := customTypes.EquipableItemsData
	items.Equipments.Accessory = removeTooHighLevelItemsFromList(items.Equipments.Accessory, level)
	items.Equipments.Back = removeTooHighLevelItemsFromList(items.Equipments.Back, level)
	items.Equipments.Belt = removeTooHighLevelItemsFromList(items.Equipments.Belt, level)
	items.Equipments.Chest = removeTooHighLevelItemsFromList(items.Equipments.Chest, level)
	items.Equipments.Hand = removeTooHighLevelItemsFromList(items.Equipments.Hand, level)
	items.Equipments.Head = removeTooHighLevelItemsFromList(items.Equipments.Head, level)
	items.Equipments.Legs = removeTooHighLevelItemsFromList(items.Equipments.Legs, level)
	items.Equipments.Neck = removeTooHighLevelItemsFromList(items.Equipments.Neck, level)
	items.Equipments.Shoulder = removeTooHighLevelItemsFromList(items.Equipments.Shoulder, level)
	items.Equipments.FirstWeapon = removeTooHighLevelItemsFromList(items.Equipments.FirstWeapon, level)
	items.Equipments.SecondWeapon = removeTooHighLevelItemsFromList(items.Equipments.SecondWeapon, level)
	items.Mounts = removeTooHighLevelItemsFromList(items.Mounts, level)
	items.Pets = removeTooHighLevelItemsFromList(items.Pets, level)
	return items
}

func removeTooHighLevelItemsFromList(itemList []customTypes.CustomItem, level int) []customTypes.CustomItem {
	resList := []customTypes.CustomItem{}
	for _, item := range itemList {
		if item.Level <= level {
			resList = append(resList, item)
		}
	}
	return resList
}
