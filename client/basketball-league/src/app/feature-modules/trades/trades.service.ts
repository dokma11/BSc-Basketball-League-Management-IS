import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { TradeProposal } from 'src/app/shared/model/tradeProposal.model';
import { TradeSubject } from 'src/app/shared/model/tradeSubject.model';
import { environment } from 'src/env/environment';

@Injectable({
  providedIn: 'root'
})
export class TradesService {

  constructor(private http: HttpClient) { }

  createTradeProposal(tradeProposal: TradeProposal): Observable<TradeProposal> {
    return this.http.post<TradeProposal>(environment.host + "tradeProposal", tradeProposal);
  }

  createTradeSubject(tradeSubject: TradeSubject): Observable<TradeSubject> {
    return this.http.post<TradeSubject>(environment.host + "tradeSubject", tradeSubject);
  }

}
