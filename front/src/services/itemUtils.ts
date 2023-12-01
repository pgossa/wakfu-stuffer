import type { Effect, Item } from '@/types/itemType'
import ap from '../assets/wakfu/ap.png'
import block from '../assets/wakfu/block.png'
import cc from '../assets/wakfu/cc.png'
import control from '../assets/wakfu/control.png'
import dodge from '../assets/wakfu/dodge.png'
import hp from '../assets/wakfu/hp.png'
import ini from '../assets/wakfu/ini.png'
import lock from '../assets/wakfu/lock.png'
import mb from '../assets/wakfu/mb.png'
import mback from '../assets/wakfu/mback.png'
import mc from '../assets/wakfu/mc.png'
import md from '../assets/wakfu/md.png'
import me from '../assets/wakfu/me.png'
import mh from '../assets/wakfu/mh.png'
import mm from '../assets/wakfu/mm.png'
import mp from '../assets/wakfu/mp.png'
import po from '../assets/wakfu/po.png'
import rback from '../assets/wakfu/rback.png'
import re from '../assets/wakfu/re.png'
import wp from '../assets/wakfu/wp.png'
import rea from '../assets/wakfu/rea.png'
import ree from '../assets/wakfu/ree.png'
import ref from '../assets/wakfu/ref.png'
import rew from '../assets/wakfu/rew.png'

export default class itemUtils {
  static getImageUrl(item: Item): string {
    if (item != undefined) {
      return 'https://vertylo.github.io/wakassets/items/' + item.ItemTypeId + item.Id + '.png'
    }
    return ''
  }

  static getColorByRarity(item: Item): string {
    if (item != undefined) {
      switch (item.Rarity) {
        case 0:
          return 'border-gray-500'
        case 1:
          return 'border-gray-100'
        case 2:
          return 'border-green-500'
        case 3:
          return 'border-orange-500'
        case 4:
          return 'border-amber-400'
        case 5:
          return 'border-pink-600'
        case 6:
          return 'border-sky-300'
        case 7:
          return 'border-purple-500'
        default:
          return 'border-gray-950'
      }
    }
    return 'border-gray-950'
  }

  static getIconByEffect(effect: Effect): string {
    switch (effect.ActionId) {
      case 20: // HP
        return hp
      case 21:
        return hp
      case 26: // MHeal
        return mh
      case 184: // Control
        return control
      case 191: // WP
        return wp
      case 150: // CritChance
        return cc
      case 80: // Resi elem
        return re
      case 82: //re fire
        return ref
      case 83: //re water
        return rew
      case 84: //re earth
        return ree
      case 85: //re air
        return rea
      case 180: // MBack
        return mback
      case 175: // Dodge
        return dodge
      case 176:
        return dodge
      case 160: // PO
        return po
      case 192:
        return wp
      case 177: // Will
        return 'https://vertylo.github.io/wakassets/aptitudes/19.png'
      case 171: // Init
        return ini
      case 1052: // MMelee
        return mm
      case 1053: // MDistance
        return md
      case 875: // Block
        return block
      case 173: // Lock
        return lock
      case 174:
        return lock
      case 41: // MP
        return mp
      case 31: // AP
        return ap
      case 71: // Rback
        return rback
      case 181: // Mback
        return mback
      case 149: // MCrit
        return mc
      case 1061: // mberzerk
        return mb
      case 1068: // Me multiple eleme
        return me
      case 1069:
        return re
      case 120: // Me
        return me
      case 400:
        return null
      default:
        console.log(effect.ActionId)
        return 'https://vertylo.github.io/wakassets/aptitudes/1.png'
    }
  }
  // TODO: Fix '-' which is getting after the number
  static getDescByEffect(effect: Effect): any {
    switch (effect.ActionId) {
      case 20: // HP
        return 'HP'
      case 21:
        return '- HP'
      case 26: // MHeal
        return 'Heal masteries'
      case 184: // Control
        return 'Control'
      case 191: // WP
        return 'WP'
      case 192:
        return '- WP'
      case 150: // CritChance
        return '% Critical Chance'
      case 80: // Resi elem
        return 'Elemental resistance'
      case 82: //re fire
        return 'Fire resistance'
      case 83: //re water
        return 'Water resistance'
      case 84: //re earth
        return 'Earth resistance'
      case 85: //re air
        return 'Air resistance'
      case 90:
        return '- Elemental resistance'
      case 96:
        return '- Earth resistance'
      case 97:
        return '- Fire resistance'
      case 98:
        return '- Water resistance'
      case 100:
        return '- Elemental resistance'
      case 180: // MBack
        return 'Back masteries'
      case 175: // Dodge
        return 'Dodge'
      case 176:
        return '- Dodge'
      case 160: // PO
        return 'Range'
      case 177: // Will
        return 'Force of Will'
      case 171: // Init
        return 'Initiative'
      case 172:
        return '- Initiative'
      case 1052: // MMelee
        return 'Melee masteries'
      case 1053: // MDistance
        return 'Distance masteries'
      case 875: // Block
        return 'Block'
      case 876:
        return '- Block'
      case 173: // Lock
        return 'Lock'
      case 174:
        return '- Lock'
      case 41: // MP
        return 'MP'
      case 31: // AP
        return 'AP'
      case 56:
        return '- AP'
      case 57:
        return '- MP'
      case 71: // Rback
        return 'Back Resistance'
      case 181: // Mback
        return 'Lose Back Masteries'
      case 149: // MCrit
        return 'Critical Masteries'
      case 1061: // mberzerk
        return 'Berzerk masteries'
      case 1068: // Me multiple eleme
        return ' masteries in ' + effect.Number + ' elements'
      case 1069:
        return ' resistance in ' + effect.Number + ' elements'
      case 120: // Me
        return 'Elemental masteries'
      case 122:
        return 'Fire Masteries'
      case 123:
        return 'Earth Masteries'
      case 124:
        return 'Water Masteries'
      case 125:
        return 'Air Masteries'
      case 130:
        return '- Elemental Masteries'
      case 132:
        return '- Fire Masteries'
      case 988:
        return 'Critical Resistance'
      case 1055:
        return 'Berzerk masteries'
      case 1056:
        return '- Critical Masteries'
      case 1059:
        return '- Melee Masteries'
      case 1060:
        return '- Distance Masteries'
      case 1062:
        return '- Berzerk Masteries'
      case 1063:
        return '- Rear Resistance'
      default:
        console.log(effect.ActionId)
        return 'Special'
    }
  }
}
