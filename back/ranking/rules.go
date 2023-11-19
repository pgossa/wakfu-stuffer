package ranking

import (
	"reflect"

	"github.com/pgossa/wakfu-stuffer/types/buildTypes"
	"github.com/pgossa/wakfu-stuffer/types/customTypes"
)

func IsBuildLegal(build buildTypes.Build) bool {
	// if IsWeaponLegal(build) &&
	if isRelicLegal(build) &&
		isEpicLegal(build) &&
		isRingLegal(build) {
		return true
	}
	return false
}

func isRingLegal(build buildTypes.Build) bool {
	return !isSameRing(build)
}

// func IsWeaponLegal(build buildTypes.Build) bool {
// 	if isWeaponEquipped(build.THand) && (isWeaponEquipped(build.MHand) || isWeaponEquipped(build.OffHand)) {
// 		return false
// 	}
// 	return true
// }

func isWeaponEquipped(item customTypes.CustomItem) bool {
	return !reflect.DeepEqual(item, customTypes.CustomItem{})
}

func isRelicLegal(build buildTypes.Build) bool {
	relicCounter := 0
	for _, item := range build.GetBuildItem() {
		if item.Rarity == 5 {
			relicCounter++
		}
	}
	return relicCounter <= 1
}

func isEpicLegal(build buildTypes.Build) bool {
	epicCounter := 0
	for _, item := range build.GetBuildItem() {
		if item.Rarity == 7 {
			epicCounter++
		}
	}
	return epicCounter <= 1
}

func isSameRing(build buildTypes.Build) bool {
	return reflect.DeepEqual(build.LeftHand, build.RightHand)
}
