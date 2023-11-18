package customType

var Equipments Equipment
var Consumables Consumable
var Resources Resource
var HavenBags HavenBag
var Miscellanies Miscellaneous
var Cosmetics Cosmetic
var QuestItems []CustomItem
var NotSorted []CustomItem

type Consumable struct {
	Id            int //106
	Food          []CustomItem
	Teleportation []CustomItem
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
	OHWeapons   OHWeapon
	THWeapons   THWeapon
	OffHWeapons OffHWeapon
}

type OHWeapon struct {
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
	id      int //521
	Tools   []CustomItem
	Emblems []CustomItem
}

type Resource struct {
	id                     int //226
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
	id                 int //602
	Transmutations     []CustomItem
	Enchantments       []CustomItem
	SublimationScrolls []CustomItem
}

type Crop struct {
	Id                 int // 758
	MinerHarvests      []CustomItem
	TrapperHarvests    []CustomItem
	LumberjackHarvests []CustomItem
	FishermanHarvests  []CustomItem
	FarmerHarvests     []CustomItem
	HerbalistHarvests  []CustomItem
	VariedCrops        []CustomItem
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
	Furnitures []CustomItem
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
