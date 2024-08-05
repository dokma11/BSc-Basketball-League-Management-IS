import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, OnInit, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSelect } from '@angular/material/select';
import { MatSnackBar } from '@angular/material/snack-bar';
import { BehaviorSubject, map, Observable, startWith, Subject, take, takeUntil } from 'rxjs';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Player } from 'src/app/shared/model/player.model';
import { RosterService } from '../roster.service';
import { Team } from 'src/app/shared/model/team.model';
import { Pick } from 'src/app/shared/model/pick.model';
import { DraftRight } from 'src/app/shared/model/draftRight.model';

@Component({
  selector: 'app-roster',
  templateUrl: './roster.component.html',
  styleUrls: ['./roster.component.css'],
  animations: [
      trigger("fadeIn", [
        transition(":enter", [
            style({ opacity: 0, transform: "translateX(-40px)" }),
            animate(
                "0.5s ease",
                style({ opacity: 1, transform: "translateX(0)" }),
            ),
        ]),
      ]),
      trigger('buttonState', [
        state('clicked', style({
          transform: 'scale(0.9)',
          opacity: 0.5
        })),
        transition('* => clicked', [
          animate('200ms')
        ]),
        transition('clicked => idle', [
          animate('200ms')
        ])
      ]),
  ],
})
export class RosterComponent implements OnInit{
  user: User | undefined;
  backgroundSize: string = '100% 100%';
  players: Player[] = [];
  picks: Pick[] = [];
  draftRights: DraftRight[] = [];
  fullTeams: Team[] = [];
  public teams: string[] = [];
  public teamCtrl: FormControl<string | null> = new FormControl<string | null>('');
  public teamFilterCtrl: FormControl<string | null> = new FormControl<string | null>('');
  private teamsSubject: BehaviorSubject<string[]> = new BehaviorSubject<string[]>(this.teams);
  public filteredTeams: Observable<string[]> = this.teamsSubject.asObservable();
  initialTeamName: string = 'Team name';
  ownTeamChosen: boolean = true;

  @ViewChild('singleSelect', { static: true }) singleSelect: MatSelect | undefined;

  protected _onDestroy = new Subject<void>();

  constructor(private authService: AuthService,
              private rosterService: RosterService,
              private snackBar: MatSnackBar,) {
    this.authService.user$.subscribe((user) => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    this.getTeams();
    
    this.teamsSubject.next(this.teams);
    this.filteredTeams = this.teamFilterCtrl.valueChanges.pipe(
      startWith(''),
      map(value => this.filterTeams(value))
    );

    setTimeout(() => {
      this.getLoggedInUsersAssets();
    }, 50);
  }

  private filterTeams(value: any): string[] {
    const filterValue = value?.toLowerCase() || '';
    return this.teams.filter(team => team.toLowerCase().includes(filterValue));;
  }

  ngAfterViewInit() {
    this.setInitialValue();
  }

  ngOnDestroy() {
    this._onDestroy.next();
    this._onDestroy.complete();
  }

  protected setInitialValue() {
    this.filteredTeams
      .pipe(take(1), takeUntil(this._onDestroy))
      .subscribe(() => {
        this.singleSelect!.compareWith = (a: string, b: string) => a.toLowerCase() === b.toLowerCase();
      });
  }

  assetForm = new FormGroup({
    selectedAssetType: new FormControl('Players', [Validators.required]),
  });

  onAssetTypeChange(event: any) {
    this.showNotification('Selected asset type: ' + this.assetForm.value.selectedAssetType);
    this.getAssets();
  }

  onTeamSelected(event: any){
    this.getAssets();
  }
  
  getTeams() {
    this.rosterService.getAllTeams().subscribe({
      next: (result: Team[] | Team) => {
        if(Array.isArray(result)){
          this.fullTeams = result;
          this.fullTeams.forEach(team =>
            this.teams.push(team.nazTim)
          )
        }
      }
    })
  }

  getLoggedInUsersAssets() {
    this.fullTeams.forEach(team =>{
      if(team.idTim == this.user?.teamId){
        this.initialTeamName = team.nazTim;
        this.ownTeamChosen = true;
        if (this.assetForm.value.selectedAssetType === 'Players') {
          this.rosterService.getAllPlayersByTeamId(team.idTim).subscribe({
            next: (result: Player[] | Player) => {
              if(Array.isArray(result)){
                this.players = result;
                // Reset other unncessary lists
                this.picks = [];
                this.draftRights = [];
              }
            }
          })
        } 
        else if (this.assetForm.value.selectedAssetType === 'Picks') {
          this.rosterService.getAllPicksByTeamId(team.idTim).subscribe({
            next: (result: Pick[] | Pick) => {
              if(Array.isArray(result)){
                this.picks = result;
                // Reset other unncessary lists
                this.players = [];
                this.draftRights = [];
              }
            }
          })
        }  
        else if (this.assetForm.value.selectedAssetType === 'Draft rights') {
          this.rosterService.getAllDraftRightsByTeamId(team.idTim).subscribe({
            next: (result: DraftRight[] | DraftRight) => {
              if(Array.isArray(result)){
                this.draftRights = result;
                // Reset other unncessary lists
                this.picks = [];
                this.players = [];
              }
            }
          })
        }
      }
    });
  }

  getAssets() {
    var teamId = 0;
    if (this.teamCtrl.value != ''){
      this.fullTeams.forEach(team => {
          if(this.teamCtrl.value === team.nazTim){
            teamId = team.idTim;
            if (team.idTim == this.user?.teamId) {
              this.ownTeamChosen = true;
            } else {
              this.ownTeamChosen = false;
            }
          }
        }
      )

      if (this.assetForm.value.selectedAssetType === 'Players') {
        this.rosterService.getAllPlayersByTeamId(teamId).subscribe({
          next: (result: Player[] | Player) => {
            if(Array.isArray(result)){
              this.players = result;
              // Reset other unncessary lists
              this.picks = [];
              this.draftRights = [];
            }
          }
        })
      } 
      else if (this.assetForm.value.selectedAssetType === 'Picks') {
        this.rosterService.getAllPicksByTeamId(teamId).subscribe({
          next: (result: Pick[] | Pick) => {
            if(Array.isArray(result)){
              this.picks = result;
              // Reset other unncessary lists
              this.players = [];
              this.draftRights = [];
            }
          }
        })
      }  
      else if (this.assetForm.value.selectedAssetType === 'Draft rights') {
        this.rosterService.getAllDraftRightsByTeamId(teamId).subscribe({
          next: (result: DraftRight[] | DraftRight) => {
            if(Array.isArray(result)){
              this.draftRights = result;
              // Reset other unncessary lists
              this.picks = [];
              this.players = [];
            }
          }
        })
      }  
    } else{
      this.getLoggedInUsersAssets();
    }
  }

  handleDialogClosed(result: any) {
    this.getAssets();
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }
}
