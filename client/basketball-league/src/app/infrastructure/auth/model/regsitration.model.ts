export interface Registration {
  firstName: string,
  lastName: string,
  email: string,
  password: string,
  role: Role
}

export enum Role {
  Recruit = 0,
  Manager = 1,
  Player = 2,
  Coach = 3,
  Scout = 4,
}