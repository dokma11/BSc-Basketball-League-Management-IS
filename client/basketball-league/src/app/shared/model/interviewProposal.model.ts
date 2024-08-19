export interface InterviewProposal {
  idPozInt?: number;
  mesOdrPozInt: string;
  datVrePozInt: Date;
  statusPozInt?: InterviewProposalStatus;
  razOdbPozInt?: string;
  idRegrut: number;
  idTrener: number;
}

export enum InterviewProposalStatus {
  WAITING = 0,
  AFFIRMED = 1,
  REJECTED = 2,
}
