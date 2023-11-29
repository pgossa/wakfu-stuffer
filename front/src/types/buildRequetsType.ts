export type BuildRequest = {
  level: number
  weightSpec: WeightSpec
  forbiddenItems: string[]
  mandatoryItems: string[]
  rarity: string[]
}

export type WeightSpec = {
  HP: number
  HPWeight: number
  AP: number
  APWeight: number
  MP: number
  MPWeight: number
  WP: number
  WPWeight: number
  PO: number
  POWeight: number
  ArmorGiven: number
  ArmorGivenWeight: number
  ArmorReceived: number
  ArmorReceivedWeight: number
  ElementaryDamages: Elementary
  ElementaryResistances: Elementary
  CriticalChance: number
  CriticalChanceWeight: number
  Block: number
  BlockWeight: number
  Initiative: number
  InitiativeWeight: number
  Dodge: number
  DodgeWeight: number
  Lock: number
  LockWeight: number
  Prospecting: number
  ProspectingWeight: number
  Wisdom: number
  WisdomWeight: number
  Control: number
  ControlWeight: number
  Will: number
  WillWeight: number
  MCritical: number
  MCriticalWeight: number
  RCritical: number
  RCriticalWeight: number
  MBack: number
  MBackWeight: number
  RBack: number
  RBackWeight: number
  MMelee: number
  MMeleeWeight: number
  MDistance: number
  MDistanceWeight: number
  MHeal: number
  MHealWeight: number
  MBerzerk: number
  MBerzerkWeight: number
}

export type Elementary = {
  Fire: number
  FireWeight: number
  Earth: number
  EarthWeight: number
  Water: number
  WaterWeight: number
  Air: number
  AirWeight: number
}
