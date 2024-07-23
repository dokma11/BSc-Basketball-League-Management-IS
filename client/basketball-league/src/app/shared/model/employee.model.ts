import { User } from "src/app/infrastructure/auth/model/user.model";
import { Uloga, UlogaZaposlenog } from "./player.model";

export interface Employee extends User {
  uloZap: UlogaZaposlenog; // Employee role
  mbrZap: string; // Identification number
}
