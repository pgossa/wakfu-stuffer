package customTypes

//TODO: Init Id or delte them
var Equipments Equipment
var Consumables Consumable
var Resources Resource
var HavenBags HavenBag
var Miscellanies Miscellaneous
var Cosmetics Cosmetic
var Mounts []CustomItem
var Pets []CustomItem
var QuestItems []CustomItem
var NotSorted []CustomItem

type Consumable struct {
	Id             int //106
	Foods          Food
	Teleportations []CustomItem
}

type Food struct {
	Id    int // 745
	Foods []CustomItem
}

type Equipment struct {
	Id                int //109
	WeaponsAndShields WeaponsAndShield
	Armors            Armor
	Bags              []CustomItem
	Accessories       Accessory
	Sets              []CustomItem
}

type WeaponsAndShield struct {
	Id          int // 100
	MHWeapons   MHWeapon
	THWeapons   THWeapon
	OffHWeapons OffHWeapon
}

type MHWeapon struct {
	Id         int //518
	Wands      []CustomItem
	Swords     []CustomItem
	Staves     []CustomItem
	ClockHands []CustomItem
	Cards      []CustomItem
}

type THWeapon struct {
	Id      int //519
	Axes    []CustomItem
	Shovels []CustomItem
	Hammers []CustomItem
	Bows    []CustomItem
	Swords  []CustomItem
	Staves  []CustomItem
}

type OffHWeapon struct {
	Id      int // 520
	Daggers []CustomItem
	Shields []CustomItem
}

type Armor struct {
	Id           int //118
	Rings        []CustomItem
	Boots        []CustomItem
	Amulets      []CustomItem
	Cloaks       []CustomItem
	Belts        []CustomItem
	Helmets      []CustomItem
	BreastPlates []CustomItem
	Epaulettes   []CustomItem
}

type Accessory struct {
	Id      int //521
	Tools   []CustomItem
	Emblems []CustomItem
}

type Resource struct {
	Id                     int //226
	MonsterResources       []CustomItem
	Improvements           Improvement
	Kamas                  []CustomItem
	RelicFragments         []CustomItem
	Recipes                []CustomItem
	SouperGlous            []CustomItem
	Crops                  Crop
	RefinementsComponents  RefinementsComponent
	MiscellaneousResources []CustomItem
}

type Improvement struct {
	Id                 int //602
	Transmutations     []CustomItem
	Enchantments       []CustomItem
	SublimationScrolls []CustomItem
}

type Crop struct {
	Id                 int // 758
	MinerHarvests      []CustomItem
	TrapperHarvests    TrapperHarvest
	LumberjackHarvests []CustomItem
	FishermanHarvests  []CustomItem
	FarmerHarvests     []CustomItem
	HerbalistHarvests  []CustomItem
	VariedCrops        []CustomItem
}

type TrapperHarvest struct {
	Ìd    int //282
	Seeds []CustomItem
}

type RefinementsComponent struct {
	Id                      int //761
	LumberjackRefinements   []CustomItem
	FarmerRefinements       []CustomItem
	HerbalistRefinements    []CustomItem
	MinerRefinements        []CustomItem
	FishermanRefinements    []CustomItem
	TrapperRefinements      []CustomItem
	ChefComponents          []CustomItem
	TailorComponents        []CustomItem
	WeaponsMasterComponents []CustomItem
	LeatherDealerComponents []CustomItem
	ArmorerComponents       []CustomItem
	JewelerComponents       []CustomItem
	BakerComponents         []CustomItem
	HandymanComponents      []CustomItem
}

type HavenBag struct {
	Id                      int //295
	HavenGems               []CustomItem
	DisplayWindowsWorkshops DisplayWindowsWorkshop
	HavenBagDecorations     HavenBagDecoration
	HavenWorlds             []CustomItem
}

type DisplayWindowsWorkshop struct {
	Id               int // 534
	DisplayWindows   []CustomItem
	CraftingMachines []CustomItem
}

type HavenBagDecoration struct {
	Id         int //535
	Furnitures Furniture
}

type Furniture struct {
	Id                    int //297
	HavenBagFurnitures    []CustomItem
	HavenBagDecorations   []CustomItem
	TradeRoomDecorations  []CustomItem
	GardenRoomDecorations []CustomItem
}

type Miscellaneous struct {
	Id   int // 385
	Keys []CustomItem
}

type Cosmetic struct {
	Id     int // 525
	Events Event
}

type Event struct {
	Id              int //756
	Transformations []CustomItem
}
