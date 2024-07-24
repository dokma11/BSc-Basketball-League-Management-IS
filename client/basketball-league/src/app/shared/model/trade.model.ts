export interface Trade {
  idTrg: number;
  datVreTrg: Date;  // Date and time of trade occurrence
  tipTrg: TradeType;
}

export enum TradeType {
  PLAYER_PLAYER = 0,
  PLAYER_PICK = 1,
  PICK_PICK = 2
}
