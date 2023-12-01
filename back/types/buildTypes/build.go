package buildTypes

import "github.com/pgossa/wakfu-stuffer/types/customTypes"

type Build struct {
	Helmet       customTypes.CustomItem
	Amulet       customTypes.CustomItem
	BreastPlate  customTypes.CustomItem
	LeftHand     customTypes.CustomItem
	RightHand    customTypes.CustomItem
	Boots        customTypes.CustomItem
	Cloak        customTypes.CustomItem
	Epaulettes   customTypes.CustomItem
	Belt         customTypes.CustomItem
	Mount        customTypes.CustomItem
	FirstWeapon  customTypes.CustomItem
	SecondWeapon customTypes.CustomItem
	Accessory    customTypes.CustomItem
	Pet          customTypes.CustomItem
}

func (build Build) NewBuildChangeItemByPosition(customItem customTypes.CustomItem, position string) Build {
	newBuild := build
	switch position {
	case "HEAD":
		newBuild.Helmet = customItem
	case "NECK":
		newBuild.Amulet = customItem
	case "CHEST":
		newBuild.BreastPlate = customItem
	case "LEFT_HAND":
		newBuild.LeftHand = customItem
		if len(customItem.EquipmentPosition.PositionDisabled) > 0 {
			newBuild.RightHand = customTypes.CustomItem{}
		}
	case "RIGHT_HAND":
		if len(newBuild.FirstWeapon.EquipmentPosition.PositionDisabled) == 0 {
			newBuild.RightHand = customItem
		}
	case "BOOTS":
		newBuild.Boots = customItem
	case "BACK":
		newBuild.Cloak = customItem
	case "SHOULDERS":
		newBuild.Epaulettes = customItem
	case "BELT":
		newBuild.Belt = customItem
	case "MOUNT":
		newBuild.Mount = customItem
	case "FIRST_WEAPON":
		newBuild.FirstWeapon = customItem
	case "SECOND_WEAPON":
		newBuild.SecondWeapon = customItem
	case "ACCESSORY":
		newBuild.Accessory = customItem
	case "PET":
		newBuild.Pet = customItem
	}
	return newBuild
}

func (build Build) GetBuildItem() []customTypes.CustomItem {
	resList := []customTypes.CustomItem{}
	resList = append(resList, build.Helmet)
	resList = append(resList, build.Amulet)
	resList = append(resList, build.BreastPlate)
	resList = append(resList, build.LeftHand)
	resList = append(resList, build.RightHand)
	resList = append(resList, build.Boots)
	resList = append(resList, build.Cloak)
	resList = append(resList, build.Epaulettes)
	resList = append(resList, build.Belt)
	resList = append(resList, build.Mount)
	resList = append(resList, build.FirstWeapon)
	resList = append(resList, build.SecondWeapon)
	resList = append(resList, build.Accessory)
	resList = append(resList, build.Pet)
	return resList
}

func (build Build) ToString() string {
	resString := ""
	resString += build.Helmet.ToString() + " "
	resString += build.Amulet.ToString() + " "
	resString += build.BreastPlate.ToString() + " "
	resString += build.LeftHand.ToString() + " "
	resString += build.RightHand.ToString() + " "
	resString += build.Boots.ToString() + " "
	resString += build.Cloak.ToString() + " "
	resString += build.Epaulettes.ToString() + " "
	resString += build.Belt.ToString() + " "
	resString += build.FirstWeapon.ToString() + " "
	resString += build.SecondWeapon.ToString() + " "
	resString += build.Accessory.ToString() + " "
	resString += build.Pet.ToString() + " "
	resString += build.Mount.ToString()
	return resString
}
