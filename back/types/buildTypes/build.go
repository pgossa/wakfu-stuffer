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
	Emblem       customTypes.CustomItem
	Pet          customTypes.CustomItem
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
	resList = append(resList, build.Emblem)
	resList = append(resList, build.Pet)
	return resList
}
