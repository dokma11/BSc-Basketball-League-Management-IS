import { TradeType } from "./trade.model";

export interface TradeProposal {
  idZahTrg: number;
  datZahTrg: Date;  // Date and time of proposalcreation 
  tipZahTrg: TradeType;
  statusZahTrg: TradeProposalStatus;
  razlogOdbij: string;  // Denial reason
}

export enum TradeProposalStatus {
  IN_PROGRESS = 0,
  ACCEPTED = 1,
  DECLINED = 2,
  CANCELLED = 3
}
