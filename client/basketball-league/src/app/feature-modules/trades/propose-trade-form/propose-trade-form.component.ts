import { trigger, transition, style, animate, state } from '@angular/animations';
import { AfterViewInit, Component, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatDialogRef, MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faMinus, faPlus } from '@fortawesome/free-solid-svg-icons';
import { AssetChoosingFormComponent } from '../asset-choosing-form/asset-choosing-form.component';
import { Observable, map, Subject, startWith, takeUntil, take, BehaviorSubject } from 'rxjs';
import { MatSelect } from '@angular/material/select';
import { Team } from 'src/app/shared/model/team.model';
import { RosterService } from '../../roster-management/roster.service';
import { Pick } from 'src/app/shared/model/pick.model';
import { TradeProposal } from 'src/app/shared/model/tradeProposal.model';
import { Player } from 'src/app/shared/model/player.model';
import { DraftRight } from 'src/app/shared/model/draftRight.model';
import { TradeType } from 'src/app/shared/model/trade.model';
import { TradesService } from '../trades.service';
import { TradeSubject, TradeSubjectType } from 'src/app/shared/model/tradeSubject.model';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';

@Component({
  selector: 'app-propose-trade-form',
  templateUrl: './propose-trade-form.component.html',
  styleUrls: ['./propose-trade-form.component.css'],
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
export class ProposeTradeFormComponent implements OnInit, AfterViewInit, OnDestroy{
  buttonState: string = 'idle';
  removeAssetButtonState: string = 'idle';
  addPartnersAssetButtonState: string = 'idle';
  addYoursAssetButtonState: string = 'idle';
  focused: string = '';
  fullTeams: Team[] = [];
  public teams: string[] = [];
  public teamCtrl: FormControl<string | null> = new FormControl<string | null>('');
  public teamFilterCtrl: FormControl<string | null> = new FormControl<string | null>('');
  private teamsSubject: BehaviorSubject<string[]> = new BehaviorSubject<string[]>(this.teams);
  public filteredTeams: Observable<string[]> = this.teamsSubject.asObservable();
  chosenPartnerPicks: Pick[] = [];
  chosenPartnerPlayers: Player[] = [];
  chosenPartnerDraftRights: DraftRight[] = [];
  chosenOwnPicks: Pick[] = [];
  chosenOwnPlayers: Player[] = [];
  chosenOwnDraftRights: DraftRight[] = [];
  user: User | undefined;

  @ViewChild('singleSelect', { static: true }) singleSelect: MatSelect | undefined;

  protected _onDestroy = new Subject<void>();

  constructor(private snackBar: MatSnackBar,
              private dialogRef: MatDialogRef<ProposeTradeFormComponent>,
              private dialogRefAsset: MatDialogRef<AssetChoosingFormComponent>,
              private dialog: MatDialog,
              private rosterService: RosterService, 
              private tradesService: TradesService,
              private authService: AuthService) {
    this.authService.user$.subscribe((user) => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    this.getTeams();

    this.teamCtrl.setValue('');
    this.teamsSubject.next(this.teams);

    this.filteredTeams = this.teamFilterCtrl.valueChanges.pipe(
      startWith(''),
      map(value => this.filterTeams(value))
    );
  }

  private filterTeams(value: any): string[] {
    const filterValue = value?.toLowerCase() || '';
    return this.teams.filter(team => team.toLowerCase().includes(filterValue));
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

  submitTradeProposalButtonClicked() {
    let teamToSend;
    this.fullTeams.forEach(team => {
      if(team.nazTim == this.teamCtrl.value){
        teamToSend = team;
      }
    })

    let tradeType: TradeType;
    if(this.chosenOwnPicks.length == 0 && this.chosenPartnerPicks.length == 0){
      tradeType = TradeType.PLAYER_PLAYER;
    } else if ( (this.chosenOwnPicks.length != 0 || this.chosenPartnerPicks.length != 0) && (this.chosenOwnPlayers.length != 0 || 
      this.chosenPartnerPlayers.length != 0 || this.chosenPartnerDraftRights.length != 0 || this.chosenOwnDraftRights.length != 0)){
      tradeType = TradeType.PLAYER_PICK;
    } else if (this.chosenOwnPicks.length != 0 && this.chosenPartnerPicks.length != 0 && this.chosenOwnDraftRights.length == 0 && 
      this.chosenPartnerDraftRights.length == 0 && this.chosenOwnPlayers.length == 0 && this.chosenPartnerPlayers.length == 0){
      tradeType = TradeType.PICK_PICK;
    } else{
      tradeType = TradeType.PLAYER_PICK;
    }

    const tradeProposal : TradeProposal = {
      datZahTrg: new Date(),
      tipZahTrg: tradeType,
      idMenadzerPos: this.user!.id, 
      //idMenadzerPrim: 1, // Umesto da unesem menadzera unecu samo tim
      idMenadzerPrimTim: teamToSend!.idTim
    };

    this.tradesService.createTradeProposal(tradeProposal).subscribe({
      next: (result: TradeProposal) => {
        console.log(result);
        this.showNotification('Trade proposal successfully sent!');
        // Id zahteva cu dobaviti na bekendu
        if(this.chosenOwnPlayers){
          this.chosenOwnPlayers.forEach(player => {
            const tradeSubject: TradeSubject = {
              tipPredTrg: TradeSubjectType.IGRAC,
              idIgrac: player.id,
            }
            this.tradesService.createTradeSubject(tradeSubject).subscribe({
              next: (result: any) => {}
            });
          })
        }
        if(this.chosenPartnerPlayers){
          this.chosenPartnerPlayers.forEach(player => {
            const tradeSubject: TradeSubject = {
              tipPredTrg: TradeSubjectType.IGRAC,
              idIgrac: player.id,
            }
            this.tradesService.createTradeSubject(tradeSubject).subscribe({
              next: (result: any) => {}
            });
          })
        }

        if(this.chosenOwnPicks){
          this.chosenOwnPicks.forEach(pick => {
            const tradeSubject: TradeSubject = {
              tipPredTrg: TradeSubjectType.PIK,
              idPik: pick.idPik,
            }
            this.tradesService.createTradeSubject(tradeSubject).subscribe({
              next: (result: any) => {}
            });
          })
        }
        if(this.chosenPartnerPicks){
          this.chosenPartnerPicks.forEach(pick => {
            const tradeSubject: TradeSubject = {
              tipPredTrg: TradeSubjectType.PIK,
              idPik: pick.idPik,
            }
            this.tradesService.createTradeSubject(tradeSubject).subscribe({
              next: (result: any) => {}
            });
          })
        }
        
        if(this.chosenOwnDraftRights){
          this.chosenOwnDraftRights.forEach(draftRight => {
            const tradeSubject: TradeSubject = {
              tipPredTrg: TradeSubjectType.PRAVA_NA_IGRACA,
              idPrava: draftRight.idPrava,            }
            this.tradesService.createTradeSubject(tradeSubject).subscribe({
              next: (result: any) => {}
            });
          })
        }
        if(this.chosenPartnerDraftRights){
          this.chosenPartnerDraftRights.forEach(draftRight => {
            const tradeSubject: TradeSubject = {
              tipPredTrg: TradeSubjectType.PRAVA_NA_IGRACA,
              idPrava: draftRight.idPrava,
            }
            this.tradesService.createTradeSubject(tradeSubject).subscribe({
              next: (result: any) => {}
            });
          })
        }
        this.dialogRef.close();
      }
    })
  }

  addPartnersAssetButtonClicked(): void {
    this.addPartnersAssetButtonState = 'clicked';
    setTimeout(() => { this.addPartnersAssetButtonState = 'idle'; }, 200);
    
    let teamToSend;
    this.fullTeams.forEach(team => {
      if(team.nazTim == this.teamCtrl.value){
        teamToSend = team;
      }
    })

    this.dialogRefAsset = this.dialog.open(AssetChoosingFormComponent, {
        data: {
          team: teamToSend,
          chosenPlayers: this.chosenPartnerPlayers,
          chosenPicks: this.chosenPartnerPicks,
          chosenDraftRights: this.chosenPartnerDraftRights
        }
    });

    if (this.dialogRefAsset) {
      this.dialogRefAsset.afterClosed().subscribe((result: any) => {
      });
    }
  }

  addYoursAssetButtonClicked(): void {
    this.addYoursAssetButtonState = 'clicked';
    setTimeout(() => { this.addYoursAssetButtonState = 'idle'; }, 200);
    
    let teamToSend;
    this.fullTeams.forEach(team => {
      if(team.idTim == this.user?.teamId){
        teamToSend = team;
      }
    })

    this.dialogRefAsset = this.dialog.open(AssetChoosingFormComponent, {
        data: {
          team: teamToSend,
          chosenPlayers: this.chosenOwnPlayers,
          chosenPicks: this.chosenOwnPicks,
          chosenDraftRights: this.chosenOwnDraftRights
        }
    });

    if (this.dialogRefAsset) {
      this.dialogRefAsset.afterClosed().subscribe((result: any) => {
      });
    }
  }

  removeAssetButtonClicked(): void {
    this.removeAssetButtonState = 'clicked';
    setTimeout(() => { this.removeAssetButtonState = 'idle'; }, 200);
  }

  onTeamSelected(event: any){
    this.chosenPartnerPicks = [];
    this.chosenPartnerPlayers = [];
    this.chosenPartnerDraftRights = [];
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

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }

  overviewClicked(){
    this.dialogRef.close();
  }

  faPlus = faPlus;
  faMinus = faMinus;
}
