export type Item = {
  Title: MultiLang;
  Description: MultiLang;
  Id: number;
  Level: number;
  Rarity: number;
  ItemTypeId: number;
  EquipmentPosition: EquipmentPosition;
  SecondarySpecs: SecondarySpecs;
  Usage: Usage
  EquipEffects: Effect[]
}

export type MultiLang = {
  Fr: string;
  En: string;
  Es: string;
  Pt: string;
}

export type EquipmentPosition = {
  Title: MultiLang;
  Position: string[]
  PositionDisabled: string[]
}

export type SecondarySpecs = {
  AccountBindType: number;
  SetId: number;
  MinimumShardSlotNumber: number;
  MaximumShardSlotNumber: number;
}

export type Usage = {
  UseParameters: UseParameters;
  UseEffects: Effect[]
  UseCriticalEffects: boolean;
}

export type UseParameters = {
  AP: number;
  MP: number;
  WP: number;
  MinRange: number;
  MaxRange: number;
  OnFreeCell: boolean;
  OnLos: boolean;
  OnlyLine: boolean;
  NoBorderCell: boolean;
  UseWorldTarget: boolean;
}

export type Effect = {
  Description: MultiLang;
  ActionId: number;
  Quantity: number;
  Number: number;
  AreaShape: number;
  AreaSize: number[];
}