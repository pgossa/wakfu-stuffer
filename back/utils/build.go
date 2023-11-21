package utils

import (
	"log"
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

func FilterEquipementsByRarity(equipmentList []customTypes.CustomItem, rarity string) []customTypes.CustomItem {
	resList := []customTypes.CustomItem{}
	for _, equipment := range equipmentList {
		if equipment.Rarity == types.Rarity.GetRarityByName(rarity) {
			resList = append(resList, equipment)
		}
	}
	return resList
}

func RemoveTooHighLevelItems(level int) customTypes.WearableItems {
	items := customTypes.WearableItemsData
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
		loopItemList = append(loopItemList, ItemHeuristic{Item: item, Heuristic: calculateHeuristicForBetterItem(item, request)})
	}
	sort.SliceStable(loopItemList, func(i, j int) bool {
		return loopItemList[i].Heuristic > loopItemList[j].Heuristic
	})
	if len(loopItemList) < itemNumber {
		itemNumber = len(loopItemList)
	}
	itemList = []customTypes.CustomItem{}
	if loopItemList != nil &&
		loopItemList[0].Item.EquipmentPosition.Position != nil &&
		loopItemList[0].Item.EquipmentPosition.Position[0] == "FIRST_WEAPON" {
		for i := 0; i < itemNumber/2; i++ {
			if loopItemList[i].Item.EquipmentPosition.PositionDisabled == nil {
				itemList = append(itemList, loopItemList[i].Item)
			}
		}
		for i := 0; i < itemNumber/2; i++ {
			itemList = append(itemList, loopItemList[i].Item)
			if loopItemList[i].Item.EquipmentPosition.PositionDisabled != nil &&
				loopItemList[i].Item.EquipmentPosition.PositionDisabled[0] == "SECOND_WEAPON" {
				itemList = append(itemList, loopItemList[i].Item)
			}
		}
	} else {
		for i := 0; i < itemNumber; i++ {
			itemList = append(itemList, loopItemList[i].Item)
		}
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
	case 400: //NullEffect
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
		//TODO: Better way to do this
		sumHeuristic := float64(0)
		elemDoneList := []string{}
		for i := 0; i < int(effect.Number); i++ {
			maxElemNumber := 0
			maxElemName := ""
			if !slices.Contains(elemDoneList, "Fire") {
				if maxElemNumber < request.WeightSpec.ElementaryDamages.Fire {
					maxElemNumber = request.WeightSpec.ElementaryDamages.Fire
					maxElemName = "Fire"
				}
			}
			if !slices.Contains(elemDoneList, "Earth") {
				if maxElemNumber < request.WeightSpec.ElementaryDamages.Earth {
					maxElemNumber = request.WeightSpec.ElementaryDamages.Earth
					maxElemName = "Earth"
				}
			}
			if !slices.Contains(elemDoneList, "Water") {
				if maxElemNumber < request.WeightSpec.ElementaryDamages.Water {
					maxElemNumber = request.WeightSpec.ElementaryDamages.Water
					maxElemName = "Water"
				}
			}
			if !slices.Contains(elemDoneList, "Air") {
				if maxElemNumber < request.WeightSpec.ElementaryDamages.Air {
					maxElemNumber = request.WeightSpec.ElementaryDamages.Air
					maxElemName = "Air"
				}
			}
			elemDoneList = append(elemDoneList, maxElemName)
			sumHeuristic += calculateEffectHeuristic(customTypes.Effect{
				ActionId: elemMaitriseNameToActionId(maxElemName),
				Quantity: float64(maxElemNumber),
			}, request)
		}
		return sumHeuristic
	case 1069: // Gain of number in ElementaryResistance
		//TODO: Better way to do this
		sumHeuristic := float64(0)
		elemDoneList := []string{}
		for i := 0; i < int(effect.Number); i++ {
			maxElemNumber := 0
			maxElemName := ""
			if !slices.Contains(elemDoneList, "Fire") {
				if maxElemNumber < request.WeightSpec.ElementaryResistances.Fire {
					maxElemNumber = request.WeightSpec.ElementaryResistances.Fire
					maxElemName = "Fire"
				}
			}
			if !slices.Contains(elemDoneList, "Earth") {
				if maxElemNumber < request.WeightSpec.ElementaryResistances.Earth {
					maxElemNumber = request.WeightSpec.ElementaryResistances.Earth
					maxElemName = "Earth"
				}
			}
			if !slices.Contains(elemDoneList, "Water") {
				if maxElemNumber < request.WeightSpec.ElementaryResistances.Water {
					maxElemNumber = request.WeightSpec.ElementaryResistances.Water
					maxElemName = "Water"
				}
			}
			if !slices.Contains(elemDoneList, "Air") {
				if maxElemNumber < request.WeightSpec.ElementaryResistances.Air {
					maxElemNumber = request.WeightSpec.ElementaryResistances.Air
					maxElemName = "Air"
				}
			}
			elemDoneList = append(elemDoneList, maxElemName)
			sumHeuristic += calculateEffectHeuristic(customTypes.Effect{
				ActionId: elemResiNameToActionId(maxElemName),
				Quantity: float64(maxElemNumber),
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
		log.Println("Elem maitrise name not found")
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
		log.Println("Elem maitrise name not found")
		return 0
	}
}

func calculateHeuristicForBetterBuild(effectName string) string {
	return ""
}
