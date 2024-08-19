export interface Registration {
  ime: string,
  prezime: string,
  email: string,
  datRodj: Date,
  lozinka: string,
  role?: Role
}

export enum Role {
  Recruit = 0,
  Manager = 1,
  Player = 2,
  Coach = 3,
  Scout = 4,
}