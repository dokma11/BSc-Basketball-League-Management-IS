import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Trade } from 'src/app/shared/model/trade.model';
import { TradeProposal } from 'src/app/shared/model/tradeProposal.model';
import { environment } from 'src/env/environment';

@Injectable({
  providedIn: 'root'
})
export class NewsService {

  constructor(private http: HttpClient) { }

  getAllTrades(): Observable<Trade> {
    return this.http.get<Trade>(environment.host + 'trade');
  }

  getTradeProposalByID(tradeProposalID: number): Observable<TradeProposal> {
    return this.http.get<TradeProposal>(environment.host + 'tradeProposal/' + tradeProposalID);
  }

}
