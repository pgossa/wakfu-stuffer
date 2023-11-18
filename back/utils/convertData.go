package utils

import (
	"errors"
	"log"
	"reflect"

	"github.com/pgossa/wakfu-stuffer/type/customType"
	"github.com/pgossa/wakfu-stuffer/type/wakfuType"
	"golang.org/x/exp/slices"
)

var multipleEffectsId = []int{1068, 1069, 1083, 1084, 832, 39, 40}

func ConvertItemListToCustomType() error {
	var resItemList []customType.CustomItem
	for _, item := range wakfuType.WItemData {
		resItem, err := convertItem(item)
		if err != nil {
			log.Println("Error while converting item: ", item.Definition.Item.Id)
			log.Println("Result is: ", resItem)
			return err
		} else {
			resItemList = append(resItemList, resItem)
		}
	}
	customType.CustomItemsData = resItemList
	return nil
}

func convertItem(wItem wakfuType.WItem) (customType.CustomItem, error) {
	resItem := customType.CustomItem{}
	convDefault(wItem, &resItem)
	convEquipmentPosition(wItem, &resItem)
	convTD(wItem, &resItem)
	convSecondarySpecs(wItem, &resItem)
	convUsage(wItem, &resItem)
	resItem.EquipEffects = convEffects(wItem.Definition.EquipEffects)
	return resItem, nil
}

func convDefault(wItem wakfuType.WItem, cItem *customType.CustomItem) error {
	if !reflect.DeepEqual(wItem.Definition.Item, wakfuType.ItemItem{}) {
		if wItem.Definition.Item.Id != 0 {
			cItem.Id = wItem.Definition.Item.Id
		}
		cItem.Level = wItem.Definition.Item.Level
		cItem.Rarity = wItem.Definition.Item.BaseParameters.Rarity
	} else {
		return errors.New("SUUU")
	}
	return nil
}

func convTD(wItem wakfuType.WItem, cItem *customType.CustomItem) {
	if (wItem.Title != wakfuType.Title{}) {
		cItem.Title.Fr = wItem.Title.Fr
		cItem.Title.En = wItem.Title.En
		cItem.Title.Es = wItem.Title.Es
		cItem.Title.Pt = wItem.Title.Pt
	}
	if (wItem.Description != wakfuType.Description{}) {
		cItem.Description.Fr = wItem.Description.Fr
		cItem.Description.En = wItem.Description.En
		cItem.Description.Es = wItem.Description.Es
		cItem.Description.Pt = wItem.Description.Pt
	}
}

func convSecondarySpecs(wItem wakfuType.WItem, cItem *customType.CustomItem) {
	if !reflect.DeepEqual(wItem.Definition.Item.BaseParameters, wakfuType.BaseParameters{}) {
		cItem.SecondarySpecs.AccountBindType = wItem.Definition.Item.BaseParameters.BindType
		cItem.SecondarySpecs.SetId = wItem.Definition.Item.BaseParameters.ItemSetId
		cItem.SecondarySpecs.MinimumShardSlotNumber = wItem.Definition.Item.BaseParameters.MinimumShardSlotNumber
		cItem.SecondarySpecs.MaximumShardSlotNumber = wItem.Definition.Item.BaseParameters.MaximumShardSlotNumber
	}
}

func convUsage(wItem wakfuType.WItem, cItem *customType.CustomItem) {
	convUsageUseParameters(wItem, cItem)
	cItem.Usage.UseEffects = convEffects(wItem.Definition.UseEffects)
	cItem.Usage.UseCriticalEffects = convEffects(wItem.Definition.UseCriticalEffects)
}

func convUsageUseParameters(wItem wakfuType.WItem, cItem *customType.CustomItem) {
	if !reflect.DeepEqual(wItem.Definition.Item.UseParameters, wakfuType.UseParameters{}) {
		cItem.Usage.UseParameters.AP = wItem.Definition.Item.UseParameters.UseCostAp
		cItem.Usage.UseParameters.MP = wItem.Definition.Item.UseParameters.UseCostMp
		cItem.Usage.UseParameters.WP = wItem.Definition.Item.UseParameters.UseCostWp
		cItem.Usage.UseParameters.MinRange = wItem.Definition.Item.UseParameters.UseRangeMin
		cItem.Usage.UseParameters.MaxRange = wItem.Definition.Item.UseParameters.UseRangeMax
		cItem.Usage.UseParameters.OnFreeCell = wItem.Definition.Item.UseParameters.UseTestFreeCell
		cItem.Usage.UseParameters.OnLos = wItem.Definition.Item.UseParameters.UseTestLos
		cItem.Usage.UseParameters.OnlyLine = wItem.Definition.Item.UseParameters.UseTestOnlyLine
		cItem.Usage.UseParameters.NoBorderCell = wItem.Definition.Item.UseParameters.UseTestNoBorderCell
		cItem.Usage.UseParameters.UseWorldTarget = wItem.Definition.Item.UseParameters.UseWorldTarget
	}
}

func convEffects(wEffects []wakfuType.Effect) []customType.Effect {
	var resEffects []customType.Effect
	for _, wEffect := range wEffects {
		// log.Println(wEffect)
		// log.Println(wEffect)
		resEffects = append(resEffects, convEffect(wEffect))

	}
	return resEffects
}

func convEffect(wEffect wakfuType.Effect) customType.Effect {
	resEffect := customType.Effect{
		Description: getEffectDescription(wEffect.Effect.Definition.ActionId),
		ActionId:    wEffect.Effect.Definition.ActionId,
	}
	if len(wEffect.Effect.Definition.Params) > 0 {
		resEffect.Quantity = wEffect.Effect.Definition.Params[0]
	}
	if slices.Contains(multipleEffectsId, wEffect.Effect.Definition.ActionId) {
		if len(wEffect.Effect.Definition.Params) > 2 {
			resEffect.Number = wEffect.Effect.Definition.Params[2]
		} else {
			resEffect.Number = wEffect.Effect.Definition.Params[1]
		}
	}
	resEffect.AreaSize = wEffect.Effect.Definition.AreaSize
	resEffect.AreaShape = wEffect.Effect.Definition.AreaShape
	return resEffect
}

func convEquipmentPosition(wItem wakfuType.WItem, cItem *customType.CustomItem) {
	cItem.ItemTypeId = wItem.Definition.Item.BaseParameters.ItemTypeId
	convEquipmentType(wItem, cItem)

}

func convEquipmentType(wItem wakfuType.WItem, cItem *customType.CustomItem) {
	for _, wType := range wakfuType.WItemTypesData {
		if wItem.Definition.Item.BaseParameters.ItemTypeId == wType.Definition.Id {
			cItem.EquipmentPosition.Title = getTypeTitle(wType)
			cItem.EquipmentPosition.Position, cItem.EquipmentPosition.PositionDisabled = getEquipmentPosition(wType)
		}
	}
}

func getTypeTitle(wType wakfuType.WItemTypes) customType.MultiLang {
	return customType.MultiLang(wType.Title)
}

func getEquipmentPosition(wType wakfuType.WItemTypes) ([]string, []string) {
	return wType.Definition.EquipmentPositions, wType.Definition.EquipmentDisabledPositions
}

func getEffectDescription(actionId int) customType.MultiLang {
	for _, action := range wakfuType.WActionData {
		if actionId == action.Definition.Id {
			return customType.MultiLang(action.Description)
		}
	}
	return customType.MultiLang{}
}
