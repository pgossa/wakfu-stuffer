import type { BuildRequest } from '@/types/buildRequetsType'

export default class buildUtils {
  static createEasyBuildRequestBody(
    level: number,
    stats: string[],
    mandatoryItems: string,
    forbiddenItems: string,
    rarity: string[]
  ): BuildRequest {
    console.log(stats)
    const request: BuildRequest = {
      level: level,
      mandatoryItems: this.parseItems(mandatoryItems),
      forbiddenItems: this.parseItems(forbiddenItems),
      rarity: rarity,
      weightSpec: {
        HP: 0,
        HPWeight: 0,
        AP: 0,
        APWeight: 0,
        MP: 0,
        MPWeight: 0,
        WP: 0,
        WPWeight: 0,
        PO: 0,
        POWeight: 0,
        ArmorGiven: 0,
        ArmorGivenWeight: 0,
        ArmorReceived: 0,
        ArmorReceivedWeight: 0,
        ElementaryDamages: {
          Fire: 0,
          FireWeight: 0,
          Earth: 0,
          EarthWeight: 0,
          Water: 0,
          WaterWeight: 0,
          Air: 0,
          AirWeight: 0
        },
        ElementaryResistances: {
          Fire: 0,
          FireWeight: 0,
          Earth: 0,
          EarthWeight: 0,
          Water: 0,
          WaterWeight: 0,
          Air: 0,
          AirWeight: 0
        },
        CriticalChance: 0,
        CriticalChanceWeight: 0,
        Block: 0,
        BlockWeight: 0,
        Initiative: 0,
        InitiativeWeight: 0,
        Dodge: 0,
        DodgeWeight: 0,
        Lock: 0,
        LockWeight: 0,
        Prospecting: 0,
        ProspectingWeight: 0,
        Wisdom: 0,
        WisdomWeight: 0,
        Control: 0,
        ControlWeight: 0,
        Will: 0,
        WillWeight: 0,
        MCritical: 0,
        MCriticalWeight: 0,
        RCritical: 0,
        RCriticalWeight: 0,
        MBack: 0,
        MBackWeight: 0,
        RBack: 0,
        RBackWeight: 0,
        MMelee: 0,
        MMeleeWeight: 0,
        MDistance: 0,
        MDistanceWeight: 0,
        MHeal: 0,
        MHealWeight: 0,
        MBerzerk: 0,
        MBerzerkWeight: 0
      }
    }
    for (const stat of stats) {
      switch (stat) {
        case 'Action Point':
          request.weightSpec.APWeight = 1
          break
        case 'Movement Point':
          request.weightSpec.MPWeight = 1
          break
        case 'Wakfu Point':
          request.weightSpec.WPWeight = 1
          break
        case 'Elemental masteries':
          request.weightSpec.ElementaryDamages = {
            FireWeight: 1,
            EarthWeight: 1,
            WaterWeight: 1,
            AirWeight: 1,
            Fire: 0,
            Earth: 0,
            Water: 0,
            Air: 0
          }
          break
        case 'Melee masteries':
          request.weightSpec.MMeleeWeight = 1
          break
        case 'Distance masteries':
          request.weightSpec.MDistanceWeight = 1
          break
        case 'Rear masteries':
          request.weightSpec.RBackWeight = 1
          break
        case 'Critical':
          request.weightSpec.CriticalChanceWeight = 1
          request.weightSpec.MCriticalWeight = 1
          break
        case 'Heal masteries':
          request.weightSpec.MHealWeight = 1
          break
        case 'Health Points':
          request.weightSpec.HPWeight = 1
          break
        case 'Elemental resistances':
          request.weightSpec.ElementaryResistances = {
            FireWeight: 1,
            EarthWeight: 1,
            WaterWeight: 1,
            AirWeight: 1,
            Fire: 0,
            Earth: 0,
            Water: 0,
            Air: 0
          }
          break
        case 'Block':
          request.weightSpec.BlockWeight = 1
          break
        case 'Dodge':
          request.weightSpec.DodgeWeight = 1
          break
        case 'Lock':
          request.weightSpec.LockWeight = 1
          break
        case 'Initiative':
          request.weightSpec.InitiativeWeight = 1
          break
        case 'Range':
          request.weightSpec.POWeight = 1
          break
        case 'Control':
          request.weightSpec.ControlWeight = 1
          break
      }
    }
    return request
  }

  static parseItems(items: string): string[] {
    return items.split(/[\r\n]+/)
  }
}
