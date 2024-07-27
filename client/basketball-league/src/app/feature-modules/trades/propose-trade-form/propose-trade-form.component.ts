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
import { TradeProposal, TradeProposalStatus } from 'src/app/shared/model/tradeProposal.model';
import { Player } from 'src/app/shared/model/player.model';
import { DraftRight } from 'src/app/shared/model/draftRight.model';
import { Trade, TradeType } from 'src/app/shared/model/trade.model';
import { TradesService } from '../trades.service';
import { TradeSubject, TradeSubjectType } from 'src/app/shared/model/tradeSubject.model';

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

  @ViewChild('singleSelect', { static: true }) singleSelect: MatSelect | undefined;

  protected _onDestroy = new Subject<void>();

  constructor(private snackBar: MatSnackBar,
              private dialogRef: MatDialogRef<ProposeTradeFormComponent>,
              private dialogRefAsset: MatDialogRef<AssetChoosingFormComponent>,
              private dialog: MatDialog,
              private rosterService: RosterService, 
              private tradesService: TradesService) {
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
    // TODO: Napraviti ifove kako bi se odredilo koji je tip trgovine u pitanju
    // TODO: Bindovati idMenadzerPos i idMenazderPrim
    // TODO: Kreirati sve tradesubjecte neophodne
    let teamToSend;
    this.fullTeams.forEach(team => {
      if(team.nazTim == this.teamCtrl.value){
        teamToSend = team;
      }
    })

    const tradeProposal : TradeProposal = {
      datZahTrg: new Date(),
      tipZahTrg: TradeType.PICK_PICK,
      idMenadzerPos: 10, // NE ZAVORAVITI DA SE PROMENI
      //idMenadzerPrim: 1, // NE ZAVORAVITI DA SE PROMENI
      idMenadzerPrimTim: teamToSend!.idTim
    };

    this.tradesService.createTradeProposal(tradeProposal).subscribe({
      next: (result: TradeProposal) => {
        this.showNotification('Trade proposal successfully sent!');
        // OVDE MOZDA POMOCU RESULT MOZE DA SE DOBIJE ID MADA NE BIH BIO NAJSIGURNIJI IPAK - AKO NE TAKO ONDA DOBAVITI SA BEKA PA TAKO RADITI...
          // TODO: Kreirati sve tradesubjecte neophodne
          // if(this.chosenOwnPlayers){
          //   this.chosenOwnPlayers.forEach(player => {
          //     const tradeSubject: TradeSubject = {
          //       tipPredTrg: TradeSubjectType.IGRAC,
          //       idIgrac: player.id,
          //       idZahTrg: 0 // OVO MORAM SKONTATI KAKO DA NAMESTIM
          //     }
          //   })
          // }
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
    this.dialogRefAsset = this.dialog.open(AssetChoosingFormComponent, {
        data: {
          team: this.fullTeams[0],  // Za sad je nula da budem kao boston, treba skolniti da se namesti na loginovanog korisnika
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
