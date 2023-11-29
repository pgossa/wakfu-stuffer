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
	itemList = utils.RemoveForbiddenItemByRarity(itemList, request.Rarity)
	log.Println(itemList)
	itemList = getBetterWeightItemsForPositions(request, itemList, 1) // No need to put higher than one
	return tryEveryCombinationWithWeight(request, itemList)
}

func getBetterWeightItemsForPositions(request types.RequestRanking, itemList customTypes.WearableItems, itemNumber int) customTypes.WearableItems {
	itemList.Equipments.Head = utils.GetBetterHeuristicItems(request, itemList.Equipments.Head, itemNumber)
	itemList.Equipments.Shoulder = utils.GetBetterHeuristicItems(request, itemList.Equipments.Shoulder, itemNumber)
	itemList.Equipments.Neck = utils.GetBetterHeuristicItems(request, itemList.Equipments.Neck, itemNumber)
	itemList.Equipments.Back = utils.GetBetterHeuristicItems(request, itemList.Equipments.Back, itemNumber)
	itemList.Equipments.Chest = utils.GetBetterHeuristicItems(request, itemList.Equipments.Chest, itemNumber)
	itemList.Equipments.Hand = utils.GetBetterHeuristicItems(request, itemList.Equipments.Hand, 2) // Need at least 2 rings
	itemList.Equipments.Belt = utils.GetBetterHeuristicItems(request, itemList.Equipments.Belt, itemNumber)
	itemList.Equipments.Legs = utils.GetBetterHeuristicItems(request, itemList.Equipments.Legs, itemNumber)
	itemList.Equipments.Accessory = utils.GetBetterHeuristicItems(request, itemList.Equipments.Accessory, itemNumber)
	itemList.Equipments.FirstWeapon = utils.GetBetterHeuristicItems(request, itemList.Equipments.FirstWeapon, 2) // Need at least 1 1-handed and 1 2-handed
	itemList.Equipments.SecondWeapon = utils.GetBetterHeuristicItems(request, itemList.Equipments.SecondWeapon, itemNumber)
	itemList.Mounts = utils.GetBetterHeuristicItems(request, itemList.Mounts, 1) // Does not make sense to test with more mounts
	itemList.Pets = utils.GetBetterHeuristicItems(request, itemList.Pets, 1)     // Does not make sense to test with more pets
	return itemList
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
						if buildHeuristic > maxHeuristic {
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
	return bestBuild
}
