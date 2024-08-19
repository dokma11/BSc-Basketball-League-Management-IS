import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/env/environment';
import { Team } from 'src/app/shared/model/team.model';
import { Player } from 'src/app/shared/model/player.model';
import { Pick } from 'src/app/shared/model/pick.model';
import { DraftRight } from 'src/app/shared/model/draftRight.model';
import { Recruit } from 'src/app/shared/model/recruit.model';
import { WishlistAsset } from 'src/app/shared/model/wishlistAsset.model';

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

  getAllAvailablePlayersByTeamId(teamId: number): Observable<Player> {
    return this.http.get<Player>(environment.host + 'player-available-team/' + teamId);
  }

  getAllTeams(): Observable<Team> {
    return this.http.get<Team>(environment.host + 'team');
  }

  getAllPicksByTeamId(teamId: number): Observable<Pick> {
    return this.http.get<Pick>(environment.host + 'pick/team/' + teamId);
  }

  getAllAvailablePicksByTeamId(teamId: number): Observable<Pick> {
    return this.http.get<Pick>(environment.host + 'pick-available-team/' + teamId);
  }

  getAllDraftRightsByTeamId(teamId: number): Observable<DraftRight> {
    return this.http.get<DraftRight>(environment.host + 'draftRight-team/' + teamId);
  }

  getAllAvailableDraftRightsByTeamId(teamId: number): Observable<DraftRight> {
    return this.http.get<DraftRight>(environment.host + 'draftRight-available-team/' + teamId);
  }

  updatePlayer(player: Player): Observable<Player> {
    return this.http.put<Player>(environment.host + 'player', player);
  }

  updatePick(pick: Pick): Observable<Pick> {
    return this.http.put<Pick>(environment.host + 'pick', pick);
  }

  updateDraftRights(draftRight: DraftRight): Observable<DraftRight> {
    return this.http.put<DraftRight>(environment.host + 'draftRight', draftRight);
  }

  getRecruitById(recruitId: number): Observable<Recruit> {
    return this.http.get<Recruit>(environment.host + 'recruit-id/' + recruitId);
  }

  addPlayerToWishlist(player: Player, teamId: number): Observable<Player> {
    return this.http.post<Player>(environment.host + 'player-wishlist/' + teamId, player);
  }

  addPickToWishlist(pick: Pick, teamId: number): Observable<Pick> {
    return this.http.post<Pick>(environment.host + 'pick-wishlist/' + teamId, pick);
  }

  addDraftRightsToWishlist(draftRight: DraftRight, teamId: number): Observable<DraftRight> {
    return this.http.post<DraftRight>(environment.host + 'draftRight-wishlist/' + teamId, draftRight);
  }

  addRecruitToWishlist(recruit: Recruit, teamId: number): Observable<Recruit> {
    return this.http.post<Recruit>(environment.host + 'recruit-wishlist/' + teamId, recruit);
  }

  removePlayerFromWishlist(player: Player, teamId: number): Observable<Player> {
    return this.http.post<Player>(environment.host + 'player-wishlist-remove/' + teamId, player);
  }

  removePickFromWishlist(pick: Pick, teamId: number): Observable<Pick> {
    return this.http.post<Pick>(environment.host + 'pick-wishlist-remove/' + teamId, pick);
  }

  removeDraftRightsFromWishlist(draftRight: DraftRight, teamId: number): Observable<DraftRight> {
    return this.http.post<DraftRight>(environment.host + 'draftRight-wishlist-remove/' + teamId, draftRight);
  }

  removeRecruitFromWishlist(recruit: Recruit, teamId: number): Observable<Recruit> {
    return this.http.post<Recruit>(environment.host + 'recruit-wishlist-remove/' + teamId, recruit);
  }
  
  getWishlistByTeamID(teamId: number): Observable<WishlistAsset> {
    return this.http.get<WishlistAsset>(environment.host + 'team-wishlist/' + teamId);
  }

}
