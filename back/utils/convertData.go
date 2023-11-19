package utils

import (
	"errors"
	"log"
	"reflect"

	"github.com/pgossa/wakfu-stuffer/types/customTypes"
	"github.com/pgossa/wakfu-stuffer/types/wakfuTypes"
	"golang.org/x/exp/slices"
)

var multipleEffectsId = []int{1068, 1069, 1083, 1084, 832, 39, 40}

func ConvertItemListToCustomType() error {
	var resItemList []customTypes.CustomItem
	for _, item := range wakfuTypes.WItemData {
		resItem, err := convertItem(item)
		if err != nil {
			log.Println("Error while converting item: ", item.Definition.Item.Id)
			log.Println("Result is: ", resItem)
			return err
		} else {
			resItemList = append(resItemList, resItem)
		}
	}
	customTypes.CustomItemsData = resItemList
	return nil
}

func convertItem(wItem wakfuTypes.WItem) (customTypes.CustomItem, error) {
	resItem := customTypes.CustomItem{}
	convDefault(wItem, &resItem)
	convEquipmentPosition(wItem, &resItem)
	convTD(wItem, &resItem)
	convSecondarySpecs(wItem, &resItem)
	convUsage(wItem, &resItem)
	resItem.EquipEffects = convEffects(wItem.Definition.EquipEffects)
	return resItem, nil
}

func convDefault(wItem wakfuTypes.WItem, cItem *customTypes.CustomItem) error {
	if !reflect.DeepEqual(wItem.Definition.Item, wakfuTypes.ItemItem{}) {
		if wItem.Definition.Item.Id != 0 {
			cItem.Id = wItem.Definition.Item.Id
		}
		cItem.Level = wItem.Definition.Item.Level
		cItem.Rarity = wItem.Definition.Item.BaseParameters.Rarity
	} else {
		return errors.New("SUUU") //TODO: oops
	}
	return nil
}

func convTD(wItem wakfuTypes.WItem, cItem *customTypes.CustomItem) {
	if (wItem.Title != wakfuTypes.Title{}) {
		cItem.Title.Fr = wItem.Title.Fr
		cItem.Title.En = wItem.Title.En
		cItem.Title.Es = wItem.Title.Es
		cItem.Title.Pt = wItem.Title.Pt
	}
	if (wItem.Description != wakfuTypes.Description{}) {
		cItem.Description.Fr = wItem.Description.Fr
		cItem.Description.En = wItem.Description.En
		cItem.Description.Es = wItem.Description.Es
		cItem.Description.Pt = wItem.Description.Pt
	}
}

func convSecondarySpecs(wItem wakfuTypes.WItem, cItem *customTypes.CustomItem) {
	if !reflect.DeepEqual(wItem.Definition.Item.BaseParameters, wakfuTypes.BaseParameters{}) {
		cItem.SecondarySpecs.AccountBindType = wItem.Definition.Item.BaseParameters.BindType
		cItem.SecondarySpecs.SetId = wItem.Definition.Item.BaseParameters.ItemSetId
		cItem.SecondarySpecs.MinimumShardSlotNumber = wItem.Definition.Item.BaseParameters.MinimumShardSlotNumber
		cItem.SecondarySpecs.MaximumShardSlotNumber = wItem.Definition.Item.BaseParameters.MaximumShardSlotNumber
	}
}

func convUsage(wItem wakfuTypes.WItem, cItem *customTypes.CustomItem) {
	convUsageUseParameters(wItem, cItem)
	cItem.Usage.UseEffects = convEffects(wItem.Definition.UseEffects)
	cItem.Usage.UseCriticalEffects = convEffects(wItem.Definition.UseCriticalEffects)
}

func convUsageUseParameters(wItem wakfuTypes.WItem, cItem *customTypes.CustomItem) {
	if !reflect.DeepEqual(wItem.Definition.Item.UseParameters, wakfuTypes.UseParameters{}) {
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

func convEffects(wEffects []wakfuTypes.Effect) []customTypes.Effect {
	var resEffects []customTypes.Effect
	for _, wEffect := range wEffects {
		resEffects = append(resEffects, convEffect(wEffect))
	}
	return resEffects
}

func convEffect(wEffect wakfuTypes.Effect) customTypes.Effect {
	resEffect := customTypes.Effect{
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

func convEquipmentPosition(wItem wakfuTypes.WItem, cItem *customTypes.CustomItem) {
	cItem.ItemTypeId = wItem.Definition.Item.BaseParameters.ItemTypeId
	convEquipmentType(wItem, cItem)

}

func convEquipmentType(wItem wakfuTypes.WItem, cItem *customTypes.CustomItem) {
	for _, wType := range wakfuTypes.WItemTypesData {
		if wItem.Definition.Item.BaseParameters.ItemTypeId == wType.Definition.Id {
			cItem.EquipmentPosition.Title = getTypeTitle(wType)
			cItem.EquipmentPosition.Position, cItem.EquipmentPosition.PositionDisabled = getEquipmentPosition(wType)
		}
	}
}

func getTypeTitle(wType wakfuTypes.WItemTypes) customTypes.MultiLang {
	return customTypes.MultiLang(wType.Title)
}

func getEquipmentPosition(wType wakfuTypes.WItemTypes) ([]string, []string) {
	return wType.Definition.EquipmentPositions, wType.Definition.EquipmentDisabledPositions
}

func getEffectDescription(actionId int) customTypes.MultiLang {
	for _, action := range wakfuTypes.WActionData {
		if actionId == action.Definition.Id {
			return customTypes.MultiLang(action.Description)
		}
	}
	return customTypes.MultiLang{}
}
