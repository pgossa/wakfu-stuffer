package ranking

import (
	"log"

	"github.com/pgossa/wakfu-stuffer/types"
	"github.com/pgossa/wakfu-stuffer/types/buildTypes"
	"github.com/pgossa/wakfu-stuffer/types/customTypes"
	"github.com/pgossa/wakfu-stuffer/utils"
)

// The idea for this algorithm is to take the top 10 non relic or epic items for each equipement position
// Then to pick the relics and epics
// And finally to try every possible combination
// There might be an issue with low number specs (such as AP, MP, ...)

func WeightRankBuild(request types.RequestRanking) buildTypes.Build {
	itemList := customTypes.WearableItemsData
	itemList = utils.RemoveTooHighLevelItems(itemList, request.Level)
	log.Println(itemList.Len())
	itemList = utils.RemoveForbiddenItemByRarity(itemList, request.Rarity)
	log.Println(itemList.Len())
	itemList = utils.RemoveItemByNames(itemList, request.ForbiddenItems)
	itemList, relicItemList, epicItemList := utils.GetBetterAndRemoveRelicAndEpicItems(request, itemList)
	relicItemList = getBetterWeightItemsForPositions(request, relicItemList)
	epicItemList = getBetterWeightItemsForPositions(request, epicItemList)
	doNotCrash(relicItemList)
	doNotCrash(epicItemList)
	itemList = getBetterWeightItemsForPositions(request, itemList)
	log.Println(itemList.Len())
	// return tryEveryCombinationWithWeight(request, itemList)
	return getBetterCombinationWithRelicAndEpic(request, itemList, relicItemList, epicItemList)
}

func doNotCrash(lol customTypes.WearableItems) {
}

func getBetterWeightItemsForPositions(request types.RequestRanking, itemList customTypes.WearableItems) customTypes.WearableItems {
	itemList.Equipments.Head = utils.GetBetterHeuristicItems(request, itemList.Equipments.Head, 1)
	itemList.Equipments.Shoulder = utils.GetBetterHeuristicItems(request, itemList.Equipments.Shoulder, 1)
	itemList.Equipments.Neck = utils.GetBetterHeuristicItems(request, itemList.Equipments.Neck, 1)
	itemList.Equipments.Back = utils.GetBetterHeuristicItems(request, itemList.Equipments.Back, 1)
	itemList.Equipments.Chest = utils.GetBetterHeuristicItems(request, itemList.Equipments.Chest, 1)
	itemList.Equipments.Hand = utils.GetBetterHeuristicItems(request, itemList.Equipments.Hand, 2) // Need at least 2 rings
	itemList.Equipments.Belt = utils.GetBetterHeuristicItems(request, itemList.Equipments.Belt, 1)
	itemList.Equipments.Legs = utils.GetBetterHeuristicItems(request, itemList.Equipments.Legs, 1)
	itemList.Equipments.Accessory = utils.GetBetterHeuristicItems(request, itemList.Equipments.Accessory, 1)
	itemList.Equipments.FirstWeapon = utils.GetBetterHeuristicItems(request, itemList.Equipments.FirstWeapon, 2) // Need at least 1 1-handed and 1 2-handed
	itemList.Equipments.SecondWeapon = utils.GetBetterHeuristicItems(request, itemList.Equipments.SecondWeapon, 1)
	itemList.Mounts = utils.GetBetterHeuristicItems(request, itemList.Mounts, 1) // Does not make sense to test with more mounts
	itemList.Pets = utils.GetBetterHeuristicItems(request, itemList.Pets, 1)     // Does not make sense to test with more pets
	return itemList
}

func getBetterCombinationWithRelicAndEpic(request types.RequestRanking, itemList customTypes.WearableItems, relicItemList customTypes.WearableItems, epicItemList customTypes.WearableItems) buildTypes.Build {
	counter := 0
	bestBuild := buildTypes.Build{}
	basicBuild := buildTypes.Build{}
	loopBuild := buildTypes.Build{}
	maxHeuristic := 0.0
	loopHeuristic := 0.0

	mount := customTypes.CustomItem{}
	pet := customTypes.CustomItem{}
	accessory := customTypes.CustomItem{}

	if len(itemList.Mounts) > 0 {
		mount = itemList.Mounts[0]
	}
	if len(itemList.Pets) > 0 {
		pet = itemList.Pets[0]
	}
	if len(itemList.Equipments.Accessory) > 0 {
		accessory = itemList.Equipments.Accessory[0]
	}
	for _, item := range itemList.Equipments.FirstWeapon {
		if len(item.EquipmentPosition.PositionDisabled) == 0 {
			loopBuild = buildTypes.Build{
				Helmet:       itemList.Equipments.Head[0],
				Epaulettes:   itemList.Equipments.Shoulder[0],
				Amulet:       itemList.Equipments.Neck[0],
				BreastPlate:  itemList.Equipments.Chest[0],
				Cloak:        itemList.Equipments.Back[0],
				LeftHand:     itemList.Equipments.Hand[0],
				RightHand:    itemList.Equipments.Hand[1],
				Belt:         itemList.Equipments.Belt[0],
				Boots:        itemList.Equipments.Legs[0],
				FirstWeapon:  item,
				SecondWeapon: itemList.Equipments.SecondWeapon[0],
				Accessory:    accessory,
				Mount:        mount,
				Pet:          pet,
			}
		} else {
			loopBuild = buildTypes.Build{
				Helmet:       itemList.Equipments.Head[0],
				Epaulettes:   itemList.Equipments.Shoulder[0],
				Amulet:       itemList.Equipments.Neck[0],
				BreastPlate:  itemList.Equipments.Chest[0],
				Cloak:        itemList.Equipments.Back[0],
				LeftHand:     itemList.Equipments.Hand[0],
				RightHand:    itemList.Equipments.Hand[1],
				Belt:         itemList.Equipments.Belt[0],
				Boots:        itemList.Equipments.Legs[0],
				FirstWeapon:  item,
				SecondWeapon: customTypes.CustomItem{},
				Accessory:    accessory,
				Mount:        mount,
				Pet:          pet,
			}
		}
		counter++
		loopHeuristic = utils.CalculateBuildHeuristic(loopBuild, request)
		if maxHeuristic <= loopHeuristic {
			maxHeuristic = loopHeuristic
			basicBuild = loopBuild
		}
	}

	log.Println("SOUCE")
	log.Println(basicBuild.ToString())
	// TODO: Much better way to do this
	// TODO: Also mandatory items get overwritten by epic / relic (ref: utils/build.go:162)
	//
	if len(relicItemList.GetAll()) <= 0 && len(epicItemList.GetAll()) <= 0 {
		return basicBuild
	}

	if len(relicItemList.GetAll()) <= 0 {
		for _, loopEpic := range epicItemList.GetAll() {
			fullBuild := buildTypes.Build{}
			fullBuild = basicBuild.NewBuildChangeItemByPosition(loopEpic, loopEpic.EquipmentPosition.Position[0])
			counter++
			loopHeuristic = utils.CalculateBuildHeuristic(fullBuild, request)
			if maxHeuristic <= loopHeuristic {
				maxHeuristic = loopHeuristic
				bestBuild = fullBuild
			}
		}
	}

	if len(epicItemList.GetAll()) <= 0 {
		for _, loopRelic := range relicItemList.GetAll() {
			fullBuild := buildTypes.Build{}
			fullBuild = basicBuild.NewBuildChangeItemByPosition(loopRelic, loopRelic.EquipmentPosition.Position[0])
			counter++
			loopHeuristic = utils.CalculateBuildHeuristic(fullBuild, request)
			if maxHeuristic <= loopHeuristic {
				maxHeuristic = loopHeuristic
				bestBuild = fullBuild
			}
		}
	}

	for _, loopRelic := range relicItemList.GetAll() {
		relicBuild := basicBuild.NewBuildChangeItemByPosition(loopRelic, loopRelic.EquipmentPosition.Position[0])
		for _, loopEpic := range epicItemList.GetAll() {
			fullBuild := buildTypes.Build{}
			if loopRelic.EquipmentPosition.Position[0] == loopEpic.EquipmentPosition.Position[0] {
				if loopRelic.EquipmentPosition.Position[0] != "LEFT_HAND" {
					break
				} else {
					fullBuild = relicBuild.NewBuildChangeItemByPosition(loopEpic, loopEpic.EquipmentPosition.Position[1])
				}
			} else {
				fullBuild = relicBuild.NewBuildChangeItemByPosition(loopEpic, loopEpic.EquipmentPosition.Position[0])
			}
			counter++
			loopHeuristic = utils.CalculateBuildHeuristic(fullBuild, request)
			if maxHeuristic <= loopHeuristic {
				maxHeuristic = loopHeuristic
				bestBuild = fullBuild
			}
		}
	}

	return bestBuild
}

// Kind of scary and beautiful
func tryEveryCombinationWithWeight(request types.RequestRanking, itemList customTypes.WearableItems) buildTypes.Build {
	maxHeuristic := float64(0)
	bestBuild := buildTypes.Build{}
	for i := 0; i < len(itemList.Equipments.Hand); i++ {
		for j := 1; j < len(itemList.Equipments.Hand); j++ {
			if i == j { // Cannot equip the same ring
				break
			}
			for _, FirstWeapon := range itemList.Equipments.FirstWeapon {
				if len(FirstWeapon.EquipmentPosition.PositionDisabled) == 0 {
					for _, SecondWeapon := range itemList.Equipments.SecondWeapon {
						build := buildTypes.Build{
							Helmet:       itemList.Equipments.Head[0],
							Epaulettes:   itemList.Equipments.Shoulder[0],
							Amulet:       itemList.Equipments.Neck[0],
							BreastPlate:  itemList.Equipments.Chest[0],
							Cloak:        itemList.Equipments.Back[0],
							LeftHand:     itemList.Equipments.Hand[i],
							RightHand:    itemList.Equipments.Hand[j],
							Belt:         itemList.Equipments.Belt[0],
							Boots:        itemList.Equipments.Legs[0],
							Accessory:    itemList.Equipments.Accessory[0],
							FirstWeapon:  FirstWeapon,
							SecondWeapon: SecondWeapon,
							Mount:        itemList.Mounts[0],
							Pet:          itemList.Pets[0],
						}
						buildHeuristic := utils.CalculateBuildHeuristic(build, request)
						if buildHeuristic >= maxHeuristic {
							maxHeuristic = buildHeuristic
							bestBuild = build
						}
					}
				} else {
					build := buildTypes.Build{
						Helmet:       itemList.Equipments.Head[0],
						Epaulettes:   itemList.Equipments.Shoulder[0],
						Amulet:       itemList.Equipments.Neck[0],
						BreastPlate:  itemList.Equipments.Chest[0],
						Cloak:        itemList.Equipments.Back[0],
						LeftHand:     itemList.Equipments.Hand[i],
						RightHand:    itemList.Equipments.Hand[j],
						Belt:         itemList.Equipments.Belt[0],
						Boots:        itemList.Equipments.Legs[0],
						Accessory:    itemList.Equipments.Accessory[0],
						FirstWeapon:  FirstWeapon,
						SecondWeapon: customTypes.CustomItem{},
						Mount:        itemList.Mounts[0],
						Pet:          itemList.Pets[0],
					}
					buildHeuristic := utils.CalculateBuildHeuristic(build, request)
					if buildHeuristic > maxHeuristic {
						maxHeuristic = buildHeuristic
						bestBuild = build
					}
				}
			}
		}
	}
	// log.Println(bestBuild)
	return bestBuild
}
