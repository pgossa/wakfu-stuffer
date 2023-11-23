import type { Item } from '@/types/itemType'

export default class itemUtils {
  static getImageUrl(item: Item): string {
    if (item != undefined){
      return "https://vertylo.github.io/wakassets/items/" + item.ItemTypeId + item.Id + ".png"
    }
    return ""
  }

  static getColorByRarity(item: Item): string {
    if (item != undefined){
      switch (item.Rarity) {
        case 0:
          return "border-gray-500"
        case 1:
          return "border-gray-100"
        case 2:
          return "border-green-500"
        case 3:
          return "border-orange-600"
        case 4:
          return "border-amber-400"
        case 5:
          return "border-pink-600"
        case 6:
          return "border-sky-300"
        case 7:
          return "border-purple-500"
        default:
          return "border-gray-950"
      }
    }
    return "border-gray-950"
  }

  static getIconByEffectId(effectId: number): string {
    switch (effectId) {
      case 20: // HP
        return "https://vertylo.github.io/wakassets/aptitudes/50.png"
      case 26: // MHeal
        return "https://vertylo.github.io/wakassets/aptitudes/136.png"
      case 184: // Control
        return "https://vertylo.github.io/wakassets/aptitudes/10.png"
      case 191: // WP
        return "https://vertylo.github.io/wakassets/aptitudes/105.png"
      case 150: // CritChance
        return "https://vertylo.github.io/wakassets/aptitudes/109.png"
      case 80: // Resi elem
        return "https://vertylo.github.io/wakassets/aptitudes/116.png"
      case 180: // MBack
        return "https://vertylo.github.io/wakassets/aptitudes/13.png"
      case 175: // Dodge
        return "https://vertylo.github.io/wakassets/aptitudes/16.png"
      case 160: // PO
        return "https://vertylo.github.io/wakassets/aptitudes/17.png"
      case 177: // Will
        return "https://vertylo.github.io/wakassets/aptitudes/19.png"
      case 171: // Init
        return "https://vertylo.github.io/wakassets/aptitudes/19.png"
      case 1052: // MMelee
        return "https://vertylo.github.io/wakassets/aptitudes/226.png"
      case 1053: // MDistance
        return "https://vertylo.github.io/wakassets/aptitudes/115.png"
      case 875: // Block
        return "https://vertylo.github.io/wakassets/aptitudes/230.png"
      case 173: // Lock
        return "https://vertylo.github.io/wakassets/aptitudes/121.png"
      case 41: // MP
        return "https://vertylo.github.io/wakassets/aptitudes/4.png"
      case 26: // DMGInf
        return "https://vertylo.github.io/wakassets/aptitudes/52.png"
      case 31: // AP
        return "https://vertylo.github.io/wakassets/aptitudes/8.png"
      default:
        return "https://vertylo.github.io/wakassets/aptitudes/1.png"
    }
  }
}