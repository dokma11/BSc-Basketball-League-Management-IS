import { Uloga } from "src/app/shared/model/player.model";

export interface User {
  id: number;
  email: string;
  ime?: string;  // Name
  prezime?: string; // Surname
  datRodj?: Date; // Date of birth
  lozinka?: string; // Password
  uloga?: Uloga; // Role
  teamId?: number;
}
