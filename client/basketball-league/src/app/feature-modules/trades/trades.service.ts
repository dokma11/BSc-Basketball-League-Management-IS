import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Employee } from 'src/app/shared/model/employee.model';
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

}
