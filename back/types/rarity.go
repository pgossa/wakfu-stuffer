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
	case "Common":
		return Rarity.Common
	case "Unusual":
		return Rarity.Unusual
	case "Rare":
		return Rarity.Rare
	case "Mythical":
		return Rarity.Mythical
	case "Legendary":
		return Rarity.Legendary
	case "Relic":
		return Rarity.Relic
	case "Souvenir":
		return Rarity.Souvenir
	case "Epic":
		return Rarity.Epic
	default:
		return -1
	}
}
