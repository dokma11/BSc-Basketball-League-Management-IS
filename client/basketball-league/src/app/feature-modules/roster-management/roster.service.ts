import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/env/environment';
import { Team } from 'src/app/shared/model/team.model';
import { Player } from 'src/app/shared/model/player.model';
import { Pick } from 'src/app/shared/model/pick.model';
import { DraftRight } from 'src/app/shared/model/draftRight.model';

@Injectable({
  providedIn: 'root'
})
export class RosterService {

  constructor(private http: HttpClient) { }

  getAllPlayers(): Observable<Player> {
    return this.http.get<Player>(environment.host + 'player');
  }

  getAllPlayersByTeamId(teamId: number): Observable<Player> {
    return this.http.get<Player>(environment.host + 'player/team/' + teamId);
  }

  getAllTeams(): Observable<Team> {
    return this.http.get<Team>(environment.host + 'team');
  }

  getAllPicksByTeamId(teamId: number): Observable<Pick> {
    return this.http.get<Pick>(environment.host + 'pick/team/' + teamId);
  }

  getAllDraftRightsByTeamId(teamId: number): Observable<DraftRight> {
    return this.http.get<DraftRight>(environment.host + 'draftRight/team/' + teamId);
  }

}
