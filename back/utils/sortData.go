package utils

import (
	"github.com/pgossa/wakfu-stuffer/types/customTypes"
)

func SortData() error {
	for _, item := range customTypes.CustomItemsData {
		sortItem(item)
	}
	sortWearableItems()
	return nil
}

// Disclaimer: Done this with chatGPT so there may be issues (not on Wearable stuff for sure)
// TODO: Issue with haven bag decoration (six instance dk why: gg ankama)
// TODO: Same for consumable
func sortItem(item customTypes.CustomItem) { // May be good to do it with a map instead but it seems less readable
	switch item.ItemTypeId {
	case 101:
		customTypes.Equipments.WeaponsAndShields.THWeapons.Axes = append(customTypes.Equipments.WeaponsAndShields.THWeapons.Axes, item)
	case 103:
		customTypes.Equipments.Armors.Rings = append(customTypes.Equipments.Armors.Rings, item)
	case 108:
		customTypes.Equipments.WeaponsAndShields.MHWeapons.Wands = append(customTypes.Equipments.WeaponsAndShields.MHWeapons.Wands, item)
	case 110:
		customTypes.Equipments.WeaponsAndShields.MHWeapons.Swords = append(customTypes.Equipments.WeaponsAndShields.MHWeapons.Swords, item)
	case 111:
		customTypes.Equipments.WeaponsAndShields.THWeapons.Shovels = append(customTypes.Equipments.WeaponsAndShields.THWeapons.Shovels, item)
	case 112:
		customTypes.Equipments.WeaponsAndShields.OffHWeapons.Daggers = append(customTypes.Equipments.WeaponsAndShields.OffHWeapons.Daggers, item)
	case 113:
		customTypes.Equipments.WeaponsAndShields.MHWeapons.Staves = append(customTypes.Equipments.WeaponsAndShields.MHWeapons.Staves, item)
	case 114:
		customTypes.Equipments.WeaponsAndShields.THWeapons.Hammers = append(customTypes.Equipments.WeaponsAndShields.THWeapons.Hammers, item)
	case 115:
		customTypes.Equipments.WeaponsAndShields.MHWeapons.ClockHands = append(customTypes.Equipments.WeaponsAndShields.MHWeapons.ClockHands, item)
	case 117:
		customTypes.Equipments.WeaponsAndShields.THWeapons.Bows = append(customTypes.Equipments.WeaponsAndShields.THWeapons.Bows, item)
	case 119:
		customTypes.Equipments.Armors.Boots = append(customTypes.Equipments.Armors.Boots, item)
	case 120:
		customTypes.Equipments.Armors.Amulets = append(customTypes.Equipments.Armors.Amulets, item)
	case 132:
		customTypes.Equipments.Armors.Cloaks = append(customTypes.Equipments.Armors.Cloaks, item)
	case 133:
		customTypes.Equipments.Armors.Belts = append(customTypes.Equipments.Armors.Belts, item)
	case 134:
		customTypes.Equipments.Armors.Helmets = append(customTypes.Equipments.Armors.Helmets, item)
	case 136:
		customTypes.Equipments.Armors.BreastPlates = append(customTypes.Equipments.Armors.BreastPlates, item)
	case 138:
		customTypes.Equipments.Armors.Epaulettes = append(customTypes.Equipments.Armors.Epaulettes, item)
	case 189:
		customTypes.Equipments.WeaponsAndShields.OffHWeapons.Shields = append(customTypes.Equipments.WeaponsAndShields.OffHWeapons.Shields, item)
	case 218:
		customTypes.Equipments.Bags = append(customTypes.Equipments.Bags, item)
	case 223:
		customTypes.Equipments.WeaponsAndShields.THWeapons.Swords = append(customTypes.Equipments.WeaponsAndShields.THWeapons.Swords, item)
	case 253:
		customTypes.Equipments.WeaponsAndShields.THWeapons.Staves = append(customTypes.Equipments.WeaponsAndShields.THWeapons.Staves, item)
	case 254:
		customTypes.Equipments.WeaponsAndShields.MHWeapons.Cards = append(customTypes.Equipments.WeaponsAndShields.MHWeapons.Cards, item)
	case 281:
		customTypes.Resources.Crops.MinerHarvests = append(customTypes.Resources.Crops.MinerHarvests, item)
	case 294:
		customTypes.HavenBags.HavenGems = append(customTypes.HavenBags.HavenGems, item)
	case 295:
		customTypes.HavenBags.HavenWorlds = append(customTypes.HavenBags.HavenWorlds, item)
	case 306:
		customTypes.Resources.Crops.LumberjackHarvests = append(customTypes.Resources.Crops.LumberjackHarvests, item)
	case 308:
		customTypes.Resources.Crops.FishermanHarvests = append(customTypes.Resources.Crops.FishermanHarvests, item)
	case 309:
		customTypes.Resources.Crops.FarmerHarvests = append(customTypes.Resources.Crops.FarmerHarvests, item)
	case 313:
		customTypes.Resources.Crops.HerbalistHarvests = append(customTypes.Resources.Crops.HerbalistHarvests, item)
	case 317:
		customTypes.Miscellanies.Keys = append(customTypes.Miscellanies.Keys, item)
	case 327:
		customTypes.Resources.RefinementsComponents.LumberjackRefinements = append(customTypes.Resources.RefinementsComponents.LumberjackRefinements, item)
	case 393:
		customTypes.Resources.RefinementsComponents.FarmerRefinements = append(customTypes.Resources.RefinementsComponents.FarmerRefinements, item)
	case 415:
		customTypes.HavenBags.HavenBagDecorations.Furnitures.HavenBagFurnitures = append(customTypes.HavenBags.HavenBagDecorations.Furnitures.HavenBagFurnitures, item)
	case 416:
		customTypes.HavenBags.HavenBagDecorations.Furnitures.HavenBagDecorations = append(customTypes.HavenBags.HavenBagDecorations.Furnitures.HavenBagDecorations, item)
	case 419:
		customTypes.Resources.Crops.TrapperHarvests.Seeds = append(customTypes.Resources.Crops.TrapperHarvests.Seeds, item)
	case 447:
		customTypes.HavenBags.DisplayWindowsWorkshops.CraftingMachines = append(customTypes.HavenBags.DisplayWindowsWorkshops.CraftingMachines, item)
	case 449:
		customTypes.HavenBags.HavenBagDecorations.Furnitures.TradeRoomDecorations = append(customTypes.HavenBags.HavenBagDecorations.Furnitures.TradeRoomDecorations, item)
	case 463:
		customTypes.Resources.RefinementsComponents.HerbalistRefinements = append(customTypes.Resources.RefinementsComponents.HerbalistRefinements, item)
	case 514:
		customTypes.Resources.RefinementsComponents.MinerRefinements = append(customTypes.Resources.RefinementsComponents.MinerRefinements, item)
	case 515:
		customTypes.HavenBags.HavenBagDecorations.Furnitures.GardenRoomDecorations = append(customTypes.HavenBags.HavenBagDecorations.Furnitures.GardenRoomDecorations, item)
	case 531:
		customTypes.Resources.MonsterResources = append(customTypes.Resources.MonsterResources, item)
	case 537:
		customTypes.Equipments.Accessories.Tools = append(customTypes.Equipments.Accessories.Tools, item)
	case 546:
		customTypes.HavenBags.HavenBagDecorations.Furnitures.HavenBagDecorations = append(customTypes.HavenBags.HavenBagDecorations.Furnitures.HavenBagDecorations, item)
	case 551:
		customTypes.QuestItems = append(customTypes.QuestItems, item)
	case 567:
		customTypes.Resources.RefinementsComponents.FishermanRefinements = append(customTypes.Resources.RefinementsComponents.FishermanRefinements, item)
	case 568:
		customTypes.Resources.RefinementsComponents.TrapperRefinements = append(customTypes.Resources.RefinementsComponents.TrapperRefinements, item)
	case 569:
		customTypes.Resources.RefinementsComponents.ChefComponents = append(customTypes.Resources.RefinementsComponents.ChefComponents, item)
	case 570:
		customTypes.Resources.RefinementsComponents.TailorComponents = append(customTypes.Resources.RefinementsComponents.TailorComponents, item)
	case 571:
		customTypes.Resources.RefinementsComponents.WeaponsMasterComponents = append(customTypes.Resources.RefinementsComponents.WeaponsMasterComponents, item)
	case 574:
		customTypes.Resources.RefinementsComponents.LeatherDealerComponents = append(customTypes.Resources.RefinementsComponents.LeatherDealerComponents, item)
	case 575:
		customTypes.Resources.RefinementsComponents.ArmorerComponents = append(customTypes.Resources.RefinementsComponents.ArmorerComponents, item)
	case 576:
		customTypes.Resources.RefinementsComponents.JewelerComponents = append(customTypes.Resources.RefinementsComponents.JewelerComponents, item)
	case 577:
		customTypes.Resources.RefinementsComponents.BakerComponents = append(customTypes.Resources.RefinementsComponents.BakerComponents, item)
	case 578:
		customTypes.Resources.RefinementsComponents.HandymanComponents = append(customTypes.Resources.RefinementsComponents.HandymanComponents, item)
	case 582:
		customTypes.Pets = append(customTypes.Pets, item)
	case 604:
		customTypes.Equipments.Sets = append(customTypes.Equipments.Sets, item)
	case 611:
		customTypes.Mounts = append(customTypes.Mounts, item)
	case 614:
		customTypes.Resources.Kamas = append(customTypes.Resources.Kamas, item)
	case 630:
		customTypes.Consumables.Teleportations = append(customTypes.Consumables.Teleportations, item)
	case 646:
		customTypes.Equipments.Accessories.Emblems = append(customTypes.Equipments.Accessories.Emblems, item)
	case 652:
		customTypes.HavenBags.HavenWorlds = append(customTypes.HavenBags.HavenWorlds, item)
	case 687:
		customTypes.Resources.RelicFragments = append(customTypes.Resources.RelicFragments, item)
	case 709:
		customTypes.Resources.Improvements.Transmutations = append(customTypes.Resources.Improvements.Transmutations, item)
	case 719:
		customTypes.Resources.Recipes = append(customTypes.Resources.Recipes, item)
	case 738:
		customTypes.Cosmetics.Events.Transformations = append(customTypes.Cosmetics.Events.Transformations, item)
	case 739:
		customTypes.Resources.SouperGlous = append(customTypes.Resources.SouperGlous, item)
	case 745:
		customTypes.Consumables.Foods.Foods = append(customTypes.Consumables.Foods.Foods, item)
	case 809:
		customTypes.Resources.Crops.VariedCrops = append(customTypes.Resources.Crops.VariedCrops, item)
	case 811:
		customTypes.Resources.Improvements.Enchantments = append(customTypes.Resources.Improvements.Enchantments, item)
	case 812:
		customTypes.Resources.Improvements.SublimationScrolls = append(customTypes.Resources.Improvements.SublimationScrolls, item)
	case 822:
		customTypes.Resources.MiscellaneousResources = append(customTypes.Resources.MiscellaneousResources, item)
	default:
		customTypes.NotSorted = append(customTypes.NotSorted, item)
	}
}

func sortWearableItems() {
	wearableItems := customTypes.WearableItems{
		Equipments: customTypes.EquipmentByPosition{
			Head:         customTypes.Equipments.Armors.Helmets,
			Shoulder:     customTypes.Equipments.Armors.Epaulettes,
			Neck:         customTypes.Equipments.Armors.Amulets,
			Back:         customTypes.Equipments.Armors.Cloaks,
			Chest:        customTypes.Equipments.Armors.BreastPlates,
			Hand:         customTypes.Equipments.Armors.Rings,
			Belt:         customTypes.Equipments.Armors.Belts,
			Legs:         customTypes.Equipments.Armors.Boots,
			Accessory:    appendToolsAndAccessories(),
			FirstWeapon:  appendMHWeaponsAndTHWeapons(),
			SecondWeapon: getOHWeapons(),
		},
		Mounts: customTypes.Mounts,
		Pets:   customTypes.Pets,
	}
	customTypes.WearableItemsData = wearableItems
}

func appendMHWeaponsAndTHWeapons() []customTypes.CustomItem {
	resList := []customTypes.CustomItem{}
	resList = append(resList, getMHWeapons()...)
	resList = append(resList, getTHWeapons()...)
	return resList
}

func appendToolsAndAccessories() []customTypes.CustomItem {
	resList := []customTypes.CustomItem{}
	resList = append(resList, customTypes.Equipments.Accessories.Tools...)
	resList = append(resList, customTypes.Equipments.Accessories.Emblems...)
	return resList
}

func getMHWeapons() []customTypes.CustomItem {
	resList := []customTypes.CustomItem{}
	resList = append(resList, customTypes.Equipments.WeaponsAndShields.MHWeapons.Staves...)
	resList = append(resList, customTypes.Equipments.WeaponsAndShields.MHWeapons.Swords...)
	resList = append(resList, customTypes.Equipments.WeaponsAndShields.MHWeapons.ClockHands...)
	resList = append(resList, customTypes.Equipments.WeaponsAndShields.MHWeapons.Cards...)
	resList = append(resList, customTypes.Equipments.WeaponsAndShields.MHWeapons.Wands...)
	return resList
}

func getOHWeapons() []customTypes.CustomItem {
	resList := []customTypes.CustomItem{}
	resList = append(resList, customTypes.Equipments.WeaponsAndShields.OffHWeapons.Daggers...)
	resList = append(resList, customTypes.Equipments.WeaponsAndShields.OffHWeapons.Shields...)
	return resList
}

func getTHWeapons() []customTypes.CustomItem {
	resList := []customTypes.CustomItem{}
	resList = append(resList, customTypes.Equipments.WeaponsAndShields.THWeapons.Axes...)
	resList = append(resList, customTypes.Equipments.WeaponsAndShields.THWeapons.Shovels...)
	resList = append(resList, customTypes.Equipments.WeaponsAndShields.THWeapons.Hammers...)
	resList = append(resList, customTypes.Equipments.WeaponsAndShields.THWeapons.Bows...)
	resList = append(resList, customTypes.Equipments.WeaponsAndShields.THWeapons.Swords...)
	resList = append(resList, customTypes.Equipments.WeaponsAndShields.THWeapons.Staves...)
	return resList
}
