export interface Contract {
  idUgo: number;
  datPotUgo: Date;  // Signing date
  datVazUgo: Date;  // Expiration date
  vredUgo: string;  // Value (in millions)
  opcUgo: ContractOption;
}

export enum ContractOption {
  PLAYER_OPTION = 0,
  TEAM_OPTION = 1,
  NO_OPTION = 2
}
