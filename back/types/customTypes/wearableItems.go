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
