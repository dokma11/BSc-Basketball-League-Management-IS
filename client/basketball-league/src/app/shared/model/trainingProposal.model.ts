export interface TrainingProposal {
  idPozTrng?: number;
  datVrePozTrng: Date;
  mesOdrPozTrng: string;
  statusPozTrng: TrainingProposalStatus;
  razOdbPozTrng?: string;
  idTrener: number;
  trajTrng: string;
  nazTipTrng: string;
  idRegrut?: number;
}

export enum TrainingProposalStatus {
  PENDING = 0,
  APPROVED = 1,
  DISAPPROVED = 2 
}
