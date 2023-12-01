package utils

import (
	"sort"

	"github.com/pgossa/wakfu-stuffer/types"
	"github.com/pgossa/wakfu-stuffer/types/buildTypes"
	"github.com/pgossa/wakfu-stuffer/types/customTypes"
	"golang.org/x/exp/slices"
)

type ItemHeuristic struct {
	Heuristic float64
	Item      customTypes.CustomItem
}

func FilterEquipmentsByRarity(items []customTypes.CustomItem, parameter []string) []customTypes.CustomItem {
	allowedRaritiesId := getRaritiesIdByNames(parameter)
	resList := []customTypes.CustomItem{}
	for _, equipment := range items {
		if slices.Contains(allowedRaritiesId, equipment.Rarity) {
			resList = append(resList, equipment)
		}
	}
	return resList
}

func RemoveForbiddenItemByRarity(items customTypes.WearableItems, parameter []string) customTypes.WearableItems {
	items = iterateOnWearableEquipment(items, func(ci []customTypes.CustomItem, a any) []customTypes.CustomItem {
		return FilterEquipmentsByRarity(ci, parameter)
	}, parameter)
	return items
}

func GetBetterAndRemoveRelicAndEpicItems(request types.RequestRanking, items customTypes.WearableItems) (customTypes.WearableItems, customTypes.WearableItems, customTypes.WearableItems) {
	relicItems := customTypes.WearableItems{}
	epicItems := customTypes.WearableItems{}
	if slices.Contains(request.Rarity, "Relic") {
		items, relicItems = getBetterRelicItems(request, items)
	}
	if slices.Contains(request.Rarity, "Epic") {
		items, epicItems = getBetterEpicItems(request, items)
	}
	return items, relicItems, epicItems
}

func getBetterRelicItems(request types.RequestRanking, items customTypes.WearableItems) (customTypes.WearableItems, customTypes.WearableItems) {
	items, relicItems := iterateOnWearableEquipmentTwoReturn(items, func(ci []customTypes.CustomItem, a any) ([]customTypes.CustomItem, []customTypes.CustomItem) {
		return getAndRemoveRarityItems(ci, types.Rarity.Relic)
	}, types.Rarity.Relic)
	return items, relicItems
}

func getBetterEpicItems(request types.RequestRanking, items customTypes.WearableItems) (customTypes.WearableItems, customTypes.WearableItems) {
	items, epicItems := iterateOnWearableEquipmentTwoReturn(items, func(ci []customTypes.CustomItem, a any) ([]customTypes.CustomItem, []customTypes.CustomItem) {
		return getAndRemoveRarityItems(ci, types.Rarity.Epic)
	}, types.Rarity.Epic)
	return items, epicItems
}

func getRaritiesIdByNames(rarity []string) []int {
	resList := []int{}
	for _, r := range rarity {
		resList = append(resList, types.Rarity.GetRarityByName(r))
	}
	return resList
}

func iterateOnWearableEquipment(items customTypes.WearableItems, fn func([]customTypes.CustomItem, any) []customTypes.CustomItem, parameter any) customTypes.WearableItems {
	items.Equipments.Accessory = fn(items.Equipments.Accessory, parameter)
	items.Equipments.Back = fn(items.Equipments.Back, parameter)
	items.Equipments.Belt = fn(items.Equipments.Belt, parameter)
	items.Equipments.Chest = fn(items.Equipments.Chest, parameter)
	items.Equipments.Hand = fn(items.Equipments.Hand, parameter)
	items.Equipments.Head = fn(items.Equipments.Head, parameter)
	items.Equipments.Legs = fn(items.Equipments.Legs, parameter)
	items.Equipments.Neck = fn(items.Equipments.Neck, parameter)
	items.Equipments.Shoulder = fn(items.Equipments.Shoulder, parameter)
	items.Equipments.FirstWeapon = fn(items.Equipments.FirstWeapon, parameter)
	items.Equipments.SecondWeapon = fn(items.Equipments.SecondWeapon, parameter)
	items.Mounts = fn(items.Mounts, parameter)
	items.Pets = fn(items.Pets, parameter)
	return items
}

func iterateOnWearableEquipmentTwoReturn(items customTypes.WearableItems, fn func([]customTypes.CustomItem, any) ([]customTypes.CustomItem, []customTypes.CustomItem), parameter any) (customTypes.WearableItems, customTypes.WearableItems) {
	secondReturn := customTypes.WearableItems{}
	items.Equipments.Accessory, secondReturn.Equipments.Accessory = fn(items.Equipments.Accessory, parameter)
	items.Equipments.Back, secondReturn.Equipments.Back = fn(items.Equipments.Back, parameter)
	items.Equipments.Belt, secondReturn.Equipments.Belt = fn(items.Equipments.Belt, parameter)
	items.Equipments.Chest, secondReturn.Equipments.Chest = fn(items.Equipments.Chest, parameter)
	items.Equipments.Hand, secondReturn.Equipments.Hand = fn(items.Equipments.Hand, parameter)
	items.Equipments.Head, secondReturn.Equipments.Head = fn(items.Equipments.Head, parameter)
	items.Equipments.Legs, secondReturn.Equipments.Legs = fn(items.Equipments.Legs, parameter)
	items.Equipments.Neck, secondReturn.Equipments.Neck = fn(items.Equipments.Neck, parameter)
	items.Equipments.Shoulder, secondReturn.Equipments.Shoulder = fn(items.Equipments.Shoulder, parameter)
	items.Equipments.FirstWeapon, secondReturn.Equipments.FirstWeapon = fn(items.Equipments.FirstWeapon, parameter)
	items.Equipments.SecondWeapon, secondReturn.Equipments.SecondWeapon = fn(items.Equipments.SecondWeapon, parameter)
	items.Mounts, secondReturn.Mounts = fn(items.Mounts, parameter)
	items.Pets, secondReturn.Pets = fn(items.Pets, parameter)
	return items, secondReturn
}

func RemoveTooHighLevelItems(items customTypes.WearableItems, level int) customTypes.WearableItems {
	items = iterateOnWearableEquipment(items, func(ci []customTypes.CustomItem, a any) []customTypes.CustomItem {
		return removeTooHighLevelItemsFromList(ci, level)
	}, level)
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

func RemoveItemByNames(items customTypes.WearableItems, names []string) customTypes.WearableItems {
	items = iterateOnWearableEquipment(items, func(ci []customTypes.CustomItem, a any) []customTypes.CustomItem {
		return removeItemsFromListByName(ci, names)
	}, names)
	return items
}

func removeItemsFromListByName(itemList []customTypes.CustomItem, names []string) []customTypes.CustomItem {
	resList := []customTypes.CustomItem{}
	for _, item := range itemList {
		if !slices.Contains(names, item.Title.En) {
			resList = append(resList, item)
		}
	}
	return resList
}

func getAndRemoveRarityItems(itemList []customTypes.CustomItem, rarity int) ([]customTypes.CustomItem, []customTypes.CustomItem) {
	resList := []customTypes.CustomItem{}
	rarityItemList := []customTypes.CustomItem{}
	for _, item := range itemList {
		if item.Rarity == rarity {
			rarityItemList = append(rarityItemList, item)
		} else {
			resList = append(resList, item)
		}
	}
	return resList, rarityItemList
}

func CalculateBuildHeuristic(build buildTypes.Build, request types.RequestRanking) float64 {
	heuristicSum := float64(0)
	for _, item := range build.GetBuildItem() {
		heuristicSum += calculateHeuristicForBetterItem(item, request)
	}
	return heuristicSum
}

func GetBetterHeuristicItems(request types.RequestRanking, itemList []customTypes.CustomItem, itemNumber int) []customTypes.CustomItem {
	loopItemList := []ItemHeuristic{}
	for _, item := range itemList {
		if slices.Contains(request.MandatoryItems, item.Title.En) {
			return []customTypes.CustomItem{item}
		}
		loopItemList = append(loopItemList, ItemHeuristic{Item: item, Heuristic: calculateHeuristicForBetterItem(item, request)})
	}
	sort.SliceStable(loopItemList, func(i, j int) bool {
		return loopItemList[i].Heuristic > loopItemList[j].Heuristic
	})
	if itemNumber > len(loopItemList) {
		itemNumber = len(loopItemList)
	}
	itemList = []customTypes.CustomItem{}
	if len(loopItemList) > 0 {
		if len(loopItemList[0].Item.EquipmentPosition.Position) > 0 &&
			loopItemList[0].Item.EquipmentPosition.Position[0] == "FIRST_WEAPON" {
			oneHand := []customTypes.CustomItem{}
			twoHand := []customTypes.CustomItem{}
			for i := 0; i < len(loopItemList); i++ {
				if len(oneHand) < itemNumber/2 &&
					len(loopItemList[i].Item.EquipmentPosition.PositionDisabled) == 0 {
					oneHand = append(oneHand, loopItemList[i].Item)
				} else if len(twoHand) < itemNumber/2 &&
					len(loopItemList[i].Item.EquipmentPosition.PositionDisabled) > 0 &&
					loopItemList[i].Item.EquipmentPosition.PositionDisabled[0] == "SECOND_WEAPON" {
					twoHand = append(twoHand, loopItemList[i].Item)
				}
			}
			itemList = append(itemList, oneHand...)
			itemList = append(itemList, twoHand...)
		} else {
			for i := 0; i < itemNumber; i++ {
				itemList = append(itemList, loopItemList[i].Item)
			}
		}
	} else {
		return nil
	}
	return itemList
}

func calculateHeuristicForBetterItem(item customTypes.CustomItem, request types.RequestRanking) float64 {
	heuristicSum := float64(0)
	for _, effect := range item.EquipEffects {
		heuristicSum += calculateEffectHeuristic(effect, request)
	}
	return heuristicSum
}

// TODO: Manage weight correctly
func calculateEffectHeuristic(effect customTypes.Effect, request types.RequestRanking) float64 {
	switch effect.ActionId {
	case 20:
		if request.WeightSpec.HPWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.HPWeight)
		}
	case 21:
		if request.WeightSpec.HPWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.HPWeight)
		}
	case 26:
		if request.WeightSpec.MHealWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.MHealWeight)
		}
	case 31:
		if request.WeightSpec.APWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.APWeight)
		}
	case 41:
		if request.WeightSpec.MPWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.MPWeight)
		}
	case 56:
		if request.WeightSpec.APWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.APWeight)
		}
	case 57:
		if request.WeightSpec.MPWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.MPWeight)
		}
	case 71:
		if request.WeightSpec.RBackWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.RBackWeight)
		}
	case 80:
		if request.WeightSpec.ElementaryResistances.EarthWeight > 0 ||
			request.WeightSpec.ElementaryResistances.FireWeight > 0 ||
			request.WeightSpec.ElementaryResistances.AirWeight > 0 ||
			request.WeightSpec.ElementaryResistances.WaterWeight > 0 {
			return effect.Quantity * float64((request.WeightSpec.ElementaryResistances.AirWeight + request.WeightSpec.ElementaryResistances.WaterWeight + request.WeightSpec.ElementaryResistances.EarthWeight + request.WeightSpec.ElementaryResistances.FireWeight))
		}
	case 82:
		if request.WeightSpec.ElementaryResistances.FireWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.ElementaryResistances.FireWeight)
		}
	case 83:
		if request.WeightSpec.ElementaryResistances.WaterWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.ElementaryResistances.WaterWeight)
		}
	case 84:
		if request.WeightSpec.ElementaryResistances.EarthWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.ElementaryResistances.EarthWeight)
		}
	case 85:
		if request.WeightSpec.ElementaryResistances.AirWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.ElementaryResistances.AirWeight)
		}
	case 90:
		if request.WeightSpec.ElementaryResistances.EarthWeight > 0 ||
			request.WeightSpec.ElementaryResistances.FireWeight > 0 ||
			request.WeightSpec.ElementaryResistances.AirWeight > 0 ||
			request.WeightSpec.ElementaryResistances.WaterWeight > 0 {
			return -effect.Quantity * float64((request.WeightSpec.ElementaryResistances.AirWeight + request.WeightSpec.ElementaryResistances.WaterWeight + request.WeightSpec.ElementaryResistances.EarthWeight + request.WeightSpec.ElementaryResistances.FireWeight))
		}
	case 96:
		if request.WeightSpec.ElementaryResistances.EarthWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.ElementaryResistances.EarthWeight)
		}
	case 97:
		if request.WeightSpec.ElementaryResistances.FireWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.ElementaryResistances.FireWeight)
		}
	case 98:
		if request.WeightSpec.ElementaryResistances.WaterWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.ElementaryResistances.WaterWeight)
		}
	case 100:
		if request.WeightSpec.ElementaryResistances.EarthWeight > 0 ||
			request.WeightSpec.ElementaryResistances.FireWeight > 0 ||
			request.WeightSpec.ElementaryResistances.AirWeight > 0 ||
			request.WeightSpec.ElementaryResistances.WaterWeight > 0 {
			return -effect.Quantity * float64((request.WeightSpec.ElementaryResistances.AirWeight + request.WeightSpec.ElementaryResistances.WaterWeight + request.WeightSpec.ElementaryResistances.EarthWeight + request.WeightSpec.ElementaryResistances.FireWeight))
		}
	case 120:
		if request.WeightSpec.ElementaryDamages.EarthWeight > 0 ||
			request.WeightSpec.ElementaryDamages.FireWeight > 0 ||
			request.WeightSpec.ElementaryDamages.AirWeight > 0 ||
			request.WeightSpec.ElementaryDamages.WaterWeight > 0 {
			return effect.Quantity * float64((request.WeightSpec.ElementaryDamages.AirWeight + request.WeightSpec.ElementaryDamages.WaterWeight + request.WeightSpec.ElementaryDamages.EarthWeight + request.WeightSpec.ElementaryDamages.FireWeight))
		}

	case 122:
		if request.WeightSpec.ElementaryDamages.FireWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.ElementaryDamages.FireWeight)
		}
	case 123:
		if request.WeightSpec.ElementaryDamages.EarthWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.ElementaryDamages.EarthWeight)
		}
	case 124:
		if request.WeightSpec.ElementaryDamages.WaterWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.ElementaryDamages.WaterWeight)
		}
	case 125:
		if request.WeightSpec.ElementaryDamages.AirWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.ElementaryDamages.AirWeight)
		}
	case 130:
		if request.WeightSpec.ElementaryDamages.EarthWeight > 0 ||
			request.WeightSpec.ElementaryDamages.FireWeight > 0 ||
			request.WeightSpec.ElementaryDamages.AirWeight > 0 ||
			request.WeightSpec.ElementaryDamages.WaterWeight > 0 {
			return -effect.Quantity * float64((request.WeightSpec.ElementaryDamages.AirWeight + request.WeightSpec.ElementaryDamages.WaterWeight + request.WeightSpec.ElementaryDamages.EarthWeight + request.WeightSpec.ElementaryDamages.FireWeight))
		}
	case 132:
		if request.WeightSpec.ElementaryDamages.FireWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.ElementaryDamages.WaterWeight)
		}
	case 149:
		if request.WeightSpec.MCriticalWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.MCriticalWeight)
		}
	case 150:
		if request.WeightSpec.CriticalChanceWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.CriticalChanceWeight)
		}
	case 160:
		if request.WeightSpec.POWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.POWeight)
		}
	case 161:
		if request.WeightSpec.POWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.POWeight)
		}
	case 162:
		if request.WeightSpec.ProspectingWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.ProspectingWeight)
		}
	case 166:
		if request.WeightSpec.WisdomWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.WisdomWeight)
		}
	case 168:
		if request.WeightSpec.CriticalChanceWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.CriticalChanceWeight)
		}
	case 171:
		if request.WeightSpec.InitiativeWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.InitiativeWeight)
		}
	case 172:
		if request.WeightSpec.InitiativeWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.InitiativeWeight)
		}
	case 173:
		if request.WeightSpec.LockWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.LockWeight)
		}
	case 174:
		if request.WeightSpec.LockWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.LockWeight)
		}
	case 175:
		if request.WeightSpec.DodgeWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.DodgeWeight)
		}
	case 176:
		if request.WeightSpec.DodgeWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.DodgeWeight)
		}
	case 177:
		if request.WeightSpec.WillWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.WillWeight)
		}
	case 180:
		if request.WeightSpec.MBackWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.MBackWeight)
		}
	case 181:
		if request.WeightSpec.MBackWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.MBackWeight)
		}
	case 184:
		if request.WeightSpec.ControlWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.ControlWeight)
		}
	case 191:
		if request.WeightSpec.WPWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.WPWeight)
		}
	case 192:
		if request.WeightSpec.WPWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.WPWeight)
		}
	case 194:
		if request.WeightSpec.WPWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.WPWeight)
		}
	case 400: // NullEffect
		return 0
	case 875:
		if request.WeightSpec.BlockWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.BlockWeight)
		}
	case 876:
		if request.WeightSpec.BlockWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.BlockWeight)
		}
	case 988:
		if request.WeightSpec.RCriticalWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.RCriticalWeight)
		}
	case 1052:
		if request.WeightSpec.MMeleeWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.MMeleeWeight)
		}
	case 1053:
		if request.WeightSpec.MDistanceWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.MDistanceWeight)
		}
	case 1055:
		if request.WeightSpec.MBerzerkWeight > 0 {
			return effect.Quantity * float64(request.WeightSpec.MBerzerkWeight)
		}
	case 1056:
		if request.WeightSpec.MBerzerkWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.MBerzerkWeight)
		}
	case 1059:
		if request.WeightSpec.MMeleeWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.MMeleeWeight)
		}
	case 1060:
		if request.WeightSpec.MDistanceWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.MDistanceWeight)
		}
	case 1061:
		if request.WeightSpec.MBerzerkWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.MBerzerkWeight)
		}
	case 1062:
		if request.WeightSpec.RCriticalWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.RCriticalWeight)
		}
	case 1063:
		if request.WeightSpec.RBackWeight > 0 {
			return -effect.Quantity * float64(request.WeightSpec.RBackWeight)
		}
	case 1068: // Gain of number in ElementaryMastery
		// TODO: Better way to do this
		sumHeuristic := float64(0)
		elemDoneList := []string{}
		for i := 0; i < int(effect.Number); i++ {
			maxElemNumber := 0
			maxElemName := ""
			if !slices.Contains(elemDoneList, "Fire") {
				if maxElemNumber < request.WeightSpec.ElementaryDamages.FireWeight {
					maxElemNumber = request.WeightSpec.ElementaryDamages.FireWeight
					maxElemName = "Fire"
				}
			}
			if !slices.Contains(elemDoneList, "Earth") {
				if maxElemNumber < request.WeightSpec.ElementaryDamages.EarthWeight {
					maxElemNumber = request.WeightSpec.ElementaryDamages.EarthWeight
					maxElemName = "Earth"
				}
			}
			if !slices.Contains(elemDoneList, "Water") {
				if maxElemNumber < request.WeightSpec.ElementaryDamages.WaterWeight {
					maxElemNumber = request.WeightSpec.ElementaryDamages.WaterWeight
					maxElemName = "Water"
				}
			}
			if !slices.Contains(elemDoneList, "Air") {
				if maxElemNumber < request.WeightSpec.ElementaryDamages.AirWeight {
					maxElemNumber = request.WeightSpec.ElementaryDamages.AirWeight
					maxElemName = "Air"
				}
			}
			elemDoneList = append(elemDoneList, maxElemName)
			sumHeuristic += calculateEffectHeuristic(customTypes.Effect{
				ActionId: elemMaitriseNameToActionId(maxElemName),
				Quantity: effect.Quantity,
			}, request)
		}
		return sumHeuristic
	case 1069: // Gain of number in ElementaryResistance
		// TODO: Better way to do this
		sumHeuristic := float64(0)
		elemDoneList := []string{}
		for i := 0; i < int(effect.Number); i++ {
			maxElemNumber := 0
			maxElemName := ""
			if !slices.Contains(elemDoneList, "Fire") {
				if maxElemNumber < request.WeightSpec.ElementaryResistances.FireWeight {
					maxElemNumber = request.WeightSpec.ElementaryResistances.FireWeight
					maxElemName = "Fire"
				}
			}
			if !slices.Contains(elemDoneList, "Earth") {
				if maxElemNumber < request.WeightSpec.ElementaryResistances.EarthWeight {
					maxElemNumber = request.WeightSpec.ElementaryResistances.EarthWeight
					maxElemName = "Earth"
				}
			}
			if !slices.Contains(elemDoneList, "Water") {
				if maxElemNumber < request.WeightSpec.ElementaryResistances.WaterWeight {
					maxElemNumber = request.WeightSpec.ElementaryResistances.WaterWeight
					maxElemName = "Water"
				}
			}
			if !slices.Contains(elemDoneList, "Air") {
				if maxElemNumber < request.WeightSpec.ElementaryResistances.AirWeight {
					maxElemNumber = request.WeightSpec.ElementaryResistances.AirWeight
					maxElemName = "Air"
				}
			}
			elemDoneList = append(elemDoneList, maxElemName)
			sumHeuristic += calculateEffectHeuristic(customTypes.Effect{
				ActionId: elemResiNameToActionId(maxElemName),
				Quantity: effect.Quantity,
			}, request)
		}
		return sumHeuristic
	case 2001: // Resources Quantity
		return 0
	default:
		// log.Println("Effect not found")
		// log.Println(effect.ActionId)
		return 0
	}
	return 0
}

func elemMaitriseNameToActionId(elemMaitriseName string) int {
	switch elemMaitriseName {
	case "Fire":
		return 122
	case "Earth":
		return 123
	case "Water":
		return 124
	case "Air":
		return 125
	default:
		// log.Println("Elem mastery name not found")
		return 0
	}
}

func elemResiNameToActionId(elemMaitriseName string) int {
	switch elemMaitriseName {
	case "Fire":
		return 82
	case "Earth":
		return 83
	case "Water":
		return 84
	case "Air":
		return 85
	default:
		// log.Println("Elem mastery name not found")
		return 0
	}
}

func calculateHeuristicForBetterBuild(effectName string) string {
	return ""
}
