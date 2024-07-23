import { Employee } from "./employee.model";

export interface Player extends Employee {
  visIgr: number; // Height
  tezIgr: number; // Weight
  pozIgr: Pozicija; // Position
}

export enum Uloga {
  UloRegrut = 0, // Recruit role
  UloZaposleni = 1, // Employee role
}

export enum UlogaZaposlenog {
  UlogaMenadzer = 0, // Manager role
  UlogaIgrac = 1, // Player role
  UlogaTrener = 2, // Coach role
  UlogaSkaut = 3  // Scout role
}

export enum Pozicija {
  PG = 0, // Point Guard
  SG = 1, // Shooting Guard
  SF = 2, // Small Forward
  PF = 3, // Power Forward
  C = 4 // Center
}
