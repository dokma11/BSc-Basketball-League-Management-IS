import { User } from "src/app/infrastructure/auth/model/user.model";
import { Pozicija } from "./player.model";

export interface Recruit extends User {
  koTelefonReg: string; // Phone number
  visReg: string; // Height
  tezReg: string; // Weight
  pozReg: Pozicija; // Position
  prosRankReg: string; // Average rank on scouting websites
  prosOcReg: string; // Average grade on scouting websites (primarily ESPN)
}
