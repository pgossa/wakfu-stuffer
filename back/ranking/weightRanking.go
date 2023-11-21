package ranking

import (
	"github.com/pgossa/wakfu-stuffer/types"
	"github.com/pgossa/wakfu-stuffer/types/buildTypes"
	"github.com/pgossa/wakfu-stuffer/types/customTypes"
	"github.com/pgossa/wakfu-stuffer/utils"
)

// The idea for this algorithm is to take the top 10 non relic or epic items for each equipement position
// Then to pick the relics and epics
// And finally to try every possible combination
// There might be an issue with low number specs (such as AP, MP, ...)

func WeightRankBuild(request types.RequestRanking) any { //buildTypes.Build {
	itemList := utils.RemoveTooHighLevelItems(request.Level)
	itemList = getBetterItemsForPositions(request, itemList, 3) //TODO: Parameter 10
	return tryEveryCombination(request, itemList)
}

func getBetterItemsForPositions(request types.RequestRanking, itemList customTypes.WearableItems, itemNumber int) customTypes.WearableItems {
	itemList.Equipments.Head = utils.GetBetterHeuristicItems(request, itemList.Equipments.Head, itemNumber)
	itemList.Equipments.Shoulder = utils.GetBetterHeuristicItems(request, itemList.Equipments.Shoulder, itemNumber)
	itemList.Equipments.Neck = utils.GetBetterHeuristicItems(request, itemList.Equipments.Neck, itemNumber)
	itemList.Equipments.Back = utils.GetBetterHeuristicItems(request, itemList.Equipments.Back, itemNumber)
	itemList.Equipments.Chest = utils.GetBetterHeuristicItems(request, itemList.Equipments.Chest, itemNumber)
	itemList.Equipments.Hand = utils.GetBetterHeuristicItems(request, itemList.Equipments.Hand, itemNumber)
	itemList.Equipments.Belt = utils.GetBetterHeuristicItems(request, itemList.Equipments.Belt, itemNumber)
	itemList.Equipments.Legs = utils.GetBetterHeuristicItems(request, itemList.Equipments.Legs, itemNumber)
	itemList.Equipments.Accessory = utils.GetBetterHeuristicItems(request, itemList.Equipments.Accessory, itemNumber)
	itemList.Equipments.FirstWeapon = utils.GetBetterHeuristicItems(request, itemList.Equipments.FirstWeapon, itemNumber)
	itemList.Equipments.SecondWeapon = utils.GetBetterHeuristicItems(request, itemList.Equipments.SecondWeapon, itemNumber)
	itemList.Mounts = utils.GetBetterHeuristicItems(request, itemList.Mounts, itemNumber)
	itemList.Pets = utils.GetBetterHeuristicItems(request, itemList.Pets, itemNumber)
	return itemList
}

// Kind of scary and beautiful
func tryEveryCombination(request types.RequestRanking, itemList customTypes.WearableItems) buildTypes.Build {
	maxHeuristic := float64(0)
	bestBuild := buildTypes.Build{}
	for _, Head := range itemList.Equipments.Head {
		for _, Shoulder := range itemList.Equipments.Shoulder {
			for _, Neck := range itemList.Equipments.Neck {
				for _, Back := range itemList.Equipments.Back {
					for _, Chest := range itemList.Equipments.Chest {
						for i := 0; i < len(itemList.Equipments.Hand); i++ {
							for j := 1; j < len(itemList.Equipments.Hand); j++ {
								for _, Belt := range itemList.Equipments.Belt {
									for _, Legs := range itemList.Equipments.Legs {
										for _, Accessory := range itemList.Equipments.Accessory {
											for _, Mounts := range itemList.Mounts {
												for _, Pets := range itemList.Pets {
													for _, FirstWeapon := range itemList.Equipments.FirstWeapon {
														if len(FirstWeapon.EquipmentPosition.PositionDisabled) == 0 {
															for _, SecondWeapon := range itemList.Equipments.SecondWeapon {
																build := buildTypes.Build{
																	Helmet:       Head,
																	Epaulettes:   Shoulder,
																	Amulet:       Neck,
																	BreastPlate:  Chest,
																	Cloak:        Back,
																	LeftHand:     itemList.Equipments.Hand[i],
																	RightHand:    itemList.Equipments.Hand[j],
																	Belt:         Belt,
																	Boots:        Legs,
																	Accessory:    Accessory,
																	FirstWeapon:  FirstWeapon,
																	SecondWeapon: SecondWeapon,
																	Mount:        Mounts,
																	Pet:          Pets,
																}
																buildHeuristic := utils.CalculateBuildHeuristic(build, request)
																if buildHeuristic > maxHeuristic {
																	maxHeuristic = buildHeuristic
																	bestBuild = build
																}
															}
														} else {
															build := buildTypes.Build{
																Helmet:       Head,
																Epaulettes:   Shoulder,
																Amulet:       Neck,
																BreastPlate:  Chest,
																Cloak:        Back,
																LeftHand:     itemList.Equipments.Hand[i],
																RightHand:    itemList.Equipments.Hand[j],
																Belt:         Belt,
																Boots:        Legs,
																Accessory:    Accessory,
																FirstWeapon:  FirstWeapon,
																SecondWeapon: customTypes.CustomItem{},
																Mount:        Mounts,
																Pet:          Pets,
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
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return bestBuild
}
