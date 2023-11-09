
import Effect from "./Effect";
export default type Weapon = {
  Name: string;
  Description: string;
  Definition: {
    level: number;
    rarity: number;
    isEpic: boolean;
    isRelic: boolean;
  };
  EquipEffects: Effect[];
  Usage: {
    AP: number;
    MP: number;
    WP: number;
    MinRange: number;
    MaxRange: number;
    NormalEffect: UseEffect;
    CriticalEffect: UseEffect;
  };
};
