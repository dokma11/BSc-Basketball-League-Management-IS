import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { DraftRight } from 'src/app/shared/model/draftRight.model';
import { Employee } from 'src/app/shared/model/employee.model';
import { Pick } from 'src/app/shared/model/pick.model';
import { Player } from 'src/app/shared/model/player.model';
import { Team } from 'src/app/shared/model/team.model';
import { TradeProposal } from 'src/app/shared/model/tradeProposal.model';
import { TradeSubject } from 'src/app/shared/model/tradeSubject.model';
import { environment } from 'src/env/environment';

@Injectable({
  providedIn: 'root'
})
export class TradesService {

  constructor(private http: HttpClient) { }

  createTradeProposal(tradeProposal: TradeProposal): Observable<TradeProposal> {
    return this.http.post<TradeProposal>(environment.host + 'tradeProposal', tradeProposal);
  }

  createTradeSubject(tradeSubject: TradeSubject): Observable<TradeSubject> {
    return this.http.post<TradeSubject>(environment.host + 'tradeSubject', tradeSubject);
  }

  getAllReceivedTradeProposalsByManagerID(managerId: number): Observable<TradeProposal> {
    return this.http.get<TradeProposal>(environment.host + 'tradeProposal-received/' + managerId);
  }

  getAllSentTradeProposalsByManagerID(managerId: number): Observable<TradeProposal> {
    return this.http.get<TradeProposal>(environment.host + 'tradeProposal-sent/' + managerId);
  }

  getTeamByManagerID(managerId: number): Observable<Team> {
    return this.http.get<Team>(environment.host + 'team-user/' + managerId);
  }

  getManagerByID(managerId: number): Observable<Employee> {
    return this.http.get<Employee>(environment.host + 'user/' + managerId);
  }

  updateTradeProposal(tradeProposal: TradeProposal): Observable<TradeProposal> {
    return this.http.put<TradeProposal>(environment.host + 'tradeProposal', tradeProposal);
  }

  commitTrade(tradeProposal: TradeProposal): Observable<TradeProposal> {
    return this.http.post<TradeProposal>(environment.host + 'tradeSubject-commit-trade', tradeProposal);
  }

  getAllTradeSubjectsByTradeProposalID(tradeProposalId: number): Observable<TradeSubject> {
    return this.http.get<TradeSubject>(environment.host + 'tradeSubject-trade/' + tradeProposalId);
  }

  getPlayerByID(playerId: number): Observable<Player> {
    return this.http.get<Player>(environment.host + 'player/' + playerId);
  }

  getPickByID(pickId: number): Observable<Pick> {
    return this.http.get<Pick>(environment.host + 'pick/' + pickId);
  }

  getDraftRightsByID(draftRightsId: number): Observable<DraftRight> {
    return this.http.get<DraftRight>(environment.host + 'draftRights/' + draftRightsId);
  }

  getTradeProposalDetailsByID(tradeProposalId: number): Observable<TradeSubject> {
    return this.http.get<TradeSubject>(environment.host + 'tradeSubject-details/' + tradeProposalId);
  }

}
