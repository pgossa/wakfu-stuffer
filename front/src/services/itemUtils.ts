import type { Item } from '@/types/itemType'

export default class itemUtils {
  static getImageUrl(item: Item): string {
    if (item != undefined){
      return "https://vertylo.github.io/wakassets/items/" + item.ItemTypeId + item.Id + ".png"
    }
    return ""
  }
}