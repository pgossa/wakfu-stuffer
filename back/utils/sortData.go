package utils

import "github.com/pgossa/wakfu-stuffer/type/customType"

func SortData() error {
	for _, item := range customType.CustomItemsData {
		sortItem(item)
	}
	return nil
}

func sortItem(item customType.CustomItem) {
	switch item.ItemTypeId {
	case 101:
		customType.Equipments.WeaponsAndShields.THWeapons.Axes = append(customType.Equipments.WeaponsAndShields.THWeapons.Axes, item)
	case 103:
		customType.Equipments.Armors.Rings = append(customType.Equipments.Armors.Rings, item)
	case 108:
		customType.Equipments.WeaponsAndShields.OHWeapons.Wands = append(customType.Equipments.WeaponsAndShields.OHWeapons.Wands, item)
	case 110:
		customType.Equipments.WeaponsAndShields.OHWeapons.Swords = append(customType.Equipments.WeaponsAndShields.OHWeapons.Swords, item)
	case 111:
		customType.Equipments.WeaponsAndShields.THWeapons.Shovels = append(customType.Equipments.WeaponsAndShields.THWeapons.Shovels, item)
	case 112:
		customType.Equipments.WeaponsAndShields.OffHWeapons.Daggers = append(customType.Equipments.WeaponsAndShields.OffHWeapons.Daggers, item)
	case 113:
		customType.Equipments.WeaponsAndShields.OHWeapons.Staves = append(customType.Equipments.WeaponsAndShields.OHWeapons.Staves, item)
	case 114:
		customType.Equipments.WeaponsAndShields.THWeapons.Hammers = append(customType.Equipments.WeaponsAndShields.THWeapons.Hammers, item)
	case 115:
		customType.Equipments.WeaponsAndShields.OHWeapons.ClockHands = append(customType.Equipments.WeaponsAndShields.OHWeapons.ClockHands, item)
	case 117:
		customType.Equipments.WeaponsAndShields.THWeapons.Bows = append(customType.Equipments.WeaponsAndShields.THWeapons.Bows, item)
	case 119:
		customType.Equipments.Armors.Boots = append(customType.Equipments.Armors.Boots, item)
	case 120:
		customType.Equipments.Armors.Amulets = append(customType.Equipments.Armors.Amulets, item)
	case 132:
		customType.Equipments.Armors.Cloaks = append(customType.Equipments.Armors.Cloaks, item)
	case 133:
		customType.Equipments.Armors.Belts = append(customType.Equipments.Armors.Belts, item)
	case 134:
		customType.Equipments.Armors.Helmets = append(customType.Equipments.Armors.Helmets, item)
	case 136:
		customType.Equipments.Armors.BreastPlates = append(customType.Equipments.Armors.BreastPlates, item)
	case 138:
		customType.Equipments.Armors.Epaulettes = append(customType.Equipments.Armors.Epaulettes, item)
	case 189:
		customType.Equipments.WeaponsAndShields.OffHWeapons.Shields = append(customType.Equipments.WeaponsAndShields.OffHWeapons.Shields, item)
	case 218:
		customType.Equipments.Bags = append(customType.Equipments.Bags, item)
	case 223:
		customType.Equipments.WeaponsAndShields.THWeapons.Swords = append(customType.Equipments.WeaponsAndShields.THWeapons.Swords, item)
	default:
		customType.NotSorted = append(customType.NotSorted, item)
	}
}
