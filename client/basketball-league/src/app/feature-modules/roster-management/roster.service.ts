import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Igrac } from './model/igrac.model';
import { Observable } from 'rxjs';
import { environment } from 'src/env/environment';
import { Team } from 'src/app/shared/model/team.model';

@Injectable({
  providedIn: 'root'
})
export class RosterService {

  constructor(private http: HttpClient) { }

  getAllPlayers(): Observable<Igrac> {
    return this.http.get<Igrac>(environment.host + 'igrac');
  }

  getAllTeams(): Observable<Team> {
    return this.http.get<Team>(environment.host + 'tim');
  }

}
