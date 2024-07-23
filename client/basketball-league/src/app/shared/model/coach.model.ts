import { Employee } from "./employee.model";

export interface Coach extends Employee {
  GodIskTrener: string;
  SpecTrener: CoachSpecialization;
}

export enum CoachSpecialization {
  OFFENSE = 0,
  DEFENSE = 1,
  PLAYER_MANAGEMENT = 2
}
