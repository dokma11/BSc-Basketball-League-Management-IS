export interface TradeSubject {
  idPredTrg?: number;
  tipPredTrg: TradeSubjectType;
  idPrava?: number; // Draft Rights foreign key
  idIgrac?: number; // Player foreign key
  idZahTrg?: number; // Trade Proposal foreign key
  idPik?: number;   // Pick foreign key
  idTim?: number; // Team foreign key
}

export enum TradeSubjectType {
  IGRAC = 0,            // Player
  PIK = 1,              // Pick
  PRAVA_NA_IGRACA = 2   // Draft Rights
}
