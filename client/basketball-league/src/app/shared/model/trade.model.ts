export interface Trade {
  idTrg: number;
  datTrg: Date;  // Date and time of trade occurrence
  tipTrg: TradeType;
  idZahTrg?: number; // Trade Proposal foreign key
}

export enum TradeType {
  PLAYER_PLAYER = 0,
  PLAYER_PICK = 1,
  PICK_PICK = 2
}
