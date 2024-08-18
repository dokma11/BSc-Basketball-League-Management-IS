export interface TrainingProposal {
  idPozTrng?: number;
  datVrePozTrng: Date;
  mesOdrPozTrng: string;
  statusPozTrng: TrainingProposalStatus;
  razOdbPozTrng?: string;
  idTrener: number;
}

export enum TrainingProposalStatus {
  PENDING = 0,
  APPROVED = 1,
  DISAPPROVED = 2 
}
