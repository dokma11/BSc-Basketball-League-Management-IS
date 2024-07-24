export interface ContractType {
  idTipUgo: number;
  nazTipUgo: ContractTypeName;
}

export enum ContractTypeName {
  PLAYER_CONTRACT = 0,
	MANAGER_CONTRACT = 1,
	COACH_CONTRACT = 2,
	SCOUT_CONTRACT = 3
}
