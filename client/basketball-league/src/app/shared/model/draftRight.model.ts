import { Pozicija } from "./player.model";

export interface DraftRight {
  idPrava: number;
  imeIgrPrava: string;        // Name of the player included in the draft rights
  prezimeIgrPrava: string;    // Surname of the player included in the draft rights
  pozicijaIgrPrava: Pozicija; // Position of the player included in the draft rights
}
