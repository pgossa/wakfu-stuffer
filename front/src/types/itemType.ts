export type Item = {
  title: MultiLang;
  description: MultiLang;
  id: number;
  level: number;
  rarity: number;
  itemTypeId: number;
  equipmentPosition: EquipmentPosition;
  secondarySpecs: SecondarySpecs;
  usage: Usage
  equipEffects: Effect[]
}

export type MultiLang = {
  fr: string;
  en: string;
  es: string;
  pt: string;
}

export type EquipmentPosition = {
  title: MultiLang;
  position: string[]
  positionDisabled: string[]
}

export type SecondarySpecs = {
  accountBindType: number;
  setId: number;
  minimumShardSlotNumber: number;
  maximumShardSlotNumber: number;
}

export type Usage = {
  useParameters: UseParameters;
  useEffects: Effect[]
  useCriticalEffects: boolean;
}

export type UseParameters = {
  ap: number;
  mp: number;
  wp: number;
  minRange: number;
  maxRange: number;
  onFreeCell: boolean;
  onLos: boolean;
  onlyLine: boolean;
  noBorderCell: boolean;
  useWorldTarget: boolean;
}

export type Effect = {
  description: MultiLang;
  actionId: number;
  quantity: number;
  number: number;
  areaShape: number;
  areaSize: number[];
}