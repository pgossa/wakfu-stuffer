import Effect from "./Effect";
export default type Armor = {
  Name: string;
  Description: string;
  Definition: {
    level: number;
    rarity: number;
    isEpic: boolean;
    isRelic: boolean;
  };
  EquipEffects: Effect[];
};
