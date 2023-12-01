package customTypes

var WearableItemsData WearableItems

type WearableItems struct {
	Equipments EquipmentByPosition
	Mounts     []CustomItem
	Pets       []CustomItem
}

type EquipmentByPosition struct {
	Head         []CustomItem
	Shoulder     []CustomItem
	Neck         []CustomItem
	Back         []CustomItem
	Chest        []CustomItem
	Hand         []CustomItem
	Belt         []CustomItem
	Legs         []CustomItem
	FirstWeapon  []CustomItem
	SecondWeapon []CustomItem
	Accessory    []CustomItem
}

func (wearable WearableItems) Len() int {
	return wearable.Equipments.Len() + len(wearable.Mounts) + len(wearable.Pets)
}

func (equipmentPosition EquipmentByPosition) Len() int {
	return len(equipmentPosition.Head) + len(equipmentPosition.Shoulder) + len(equipmentPosition.Neck) + len(equipmentPosition.Back) +
		len(equipmentPosition.Chest) + len(equipmentPosition.Hand) + len(equipmentPosition.Belt) + len(equipmentPosition.Legs) +
		len(equipmentPosition.FirstWeapon) + len(equipmentPosition.SecondWeapon) + len(equipmentPosition.Accessory)
}

func (wearable WearableItems) GetAll() []CustomItem {
	allItemsList := []CustomItem{}
	allItemsList = append(allItemsList, wearable.Equipments.GetAll()...)
	allItemsList = append(allItemsList, wearable.Mounts...)
	allItemsList = append(allItemsList, wearable.Pets...)
	return allItemsList
}

func (equipmentPosition EquipmentByPosition) GetAll() []CustomItem {
	allItemsList := []CustomItem{}
	allItemsList = append(allItemsList, equipmentPosition.Head...)
	allItemsList = append(allItemsList, equipmentPosition.Shoulder...)
	allItemsList = append(allItemsList, equipmentPosition.Neck...)
	allItemsList = append(allItemsList, equipmentPosition.Back...)
	allItemsList = append(allItemsList, equipmentPosition.Chest...)
	allItemsList = append(allItemsList, equipmentPosition.Hand...)
	allItemsList = append(allItemsList, equipmentPosition.Belt...)
	allItemsList = append(allItemsList, equipmentPosition.Legs...)
	allItemsList = append(allItemsList, equipmentPosition.FirstWeapon...)
	allItemsList = append(allItemsList, equipmentPosition.SecondWeapon...)
	allItemsList = append(allItemsList, equipmentPosition.Accessory...)
	return allItemsList
}
