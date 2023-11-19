package types

var Rarity = rarity{
	Common:    0, // not existing in json provided by wakfu
	Unusual:   1,
	Rare:      2,
	Mythical:  3,
	Legendary: 4,
	Relic:     5,
	Souvenir:  6,
	Epic:      7,
}

type rarity struct {
	Common    int
	Unusual   int
	Rare      int
	Mythical  int
	Legendary int
	Relic     int
	Souvenir  int
	Epic      int
}

func (r *rarity) GetRarityByName(rarityName string) int {
	switch rarityName {
	case "common":
		return Rarity.Common
	case "unusual":
		return Rarity.Unusual
	case "rare":
		return Rarity.Rare
	case "mythical":
		return Rarity.Mythical
	case "legendary":
		return Rarity.Legendary
	case "relic":
		return Rarity.Relic
	case "souvenir":
		return Rarity.Souvenir
	case "epic":
		return Rarity.Epic
	default:
		return -1
	}
}
