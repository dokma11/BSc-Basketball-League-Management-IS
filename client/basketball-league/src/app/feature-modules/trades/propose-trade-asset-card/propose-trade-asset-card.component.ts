import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { DraftRight } from 'src/app/shared/model/draftRight.model';
import { Pick } from 'src/app/shared/model/pick.model';
import { Player } from 'src/app/shared/model/player.model';
import { TradeSubject } from 'src/app/shared/model/tradeSubject.model';
import { TradesService } from '../trades.service';

@Component({
  selector: 'app-propose-trade-asset-card',
  templateUrl: './propose-trade-asset-card.component.html',
  styleUrls: ['./propose-trade-asset-card.component.css'],
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
export class ProposeTradeAssetCardComponent implements OnInit{
  removeAssetButtonState: string = 'idle';
  description: string | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  ownTeam: boolean = false;
  @Input() showButton: boolean = true;
  @Input() chosenPlayer!: Player;
  @Input() chosenPick!: Pick;
  @Input() chosenDraftRight!: DraftRight;
  @Input() ownTradeSubject!: TradeSubject;
  @Input() partnerTradeSubject!: TradeSubject;
  @Input() chosenPlayers!: Player[];
  @Input() chosenPicks!: Pick[];
  @Input() chosenDraftRights!: DraftRight[];
  @Input() ownTradeSubjects!: TradeSubject[];
  @Input() partnerTradeSubjects!: TradeSubject[];
  @Input() detailedView: boolean = false;
  player: Player | undefined;
  pick: Pick | undefined;
  draftRights: DraftRight | undefined;

  constructor(private snackBar: MatSnackBar, 
              private tradesService: TradesService) {}

  ngOnInit(): void {
    if(!this.detailedView){
      if(this.chosenPlayer) {
        this.description = 'Player: ' + this.chosenPlayer.ime + ' ' + this.chosenPlayer.prezime;
      } else if(this.chosenPick) {
        if(this.chosenPick.redBrPik == '0'){
          if(this.chosenPick.brRunPik == '1') {
            this.description = 'Pick: 1st round in year ' + this.chosenPick.godPik;
          } else{
            this.description = 'Pick: 2nd round in year ' + this.chosenPick.godPik;
          }
        }else{
          this.description = 'Pick order: ' + this.chosenPick.redBrPik + ' Round: ' + this.chosenPick.brRunPik + ' Year:' + this.chosenPick.godPik;
        }
      } else if(this.chosenDraftRight) {
        console.log('ma n p');
      }
    } else {
      if(this.ownTradeSubject){
        if(this.ownTradeSubject.tipPredTrg == 0) { // Player type
          this.tradesService.getPlayerByID(this.ownTradeSubject.idIgrac!).subscribe({
            next: (result: Player) => {
              this.player = result;
              this.description = 'Player: ' + this.player.ime + ' ' + this.player.prezime;
            }
          });
        } else if(this.ownTradeSubject.tipPredTrg == 1) { // Pick type
          this.tradesService.getPickByID(this.ownTradeSubject.idIgrac!).subscribe({
            next: (result: Pick) => {
              this.pick = result;
              if(this.pick.redBrPik == '0'){
                if(this.pick.brRunPik == '1') {
                  this.description = 'Pick: 1st round in year ' + this.pick.godPik;
                } else{
                  this.description = 'Pick: 2nd round in year ' + this.pick.godPik;
                }
              }else{
                this.description = 'Pick order: ' + this.pick.redBrPik + ' Round: ' + this.pick.brRunPik + ' Year:' + this.pick.godPik;
              }
            }
          });
        } else if(this.ownTradeSubject.tipPredTrg == 2) { // Draft Rights
          this.tradesService.getDraftRightsByID(this.ownTradeSubject.idIgrac!).subscribe({
            next: (result: DraftRight) => {
              this.draftRights = result;
              console.log('ma nek se npk')
            }
          });
        }  
      } else {
        if(this.partnerTradeSubject.tipPredTrg == 0) { // Player type
          this.tradesService.getPlayerByID(this.partnerTradeSubject.idIgrac!).subscribe({
            next: (result: Player) => {
              this.player = result;
              this.description = 'Player: ' + this.player.ime + ' ' + this.player.prezime;
            }
          });
        } else if(this.partnerTradeSubject.tipPredTrg == 1) { // Pick type
          this.tradesService.getPickByID(this.partnerTradeSubject.idIgrac!).subscribe({
            next: (result: Pick) => {
              this.pick = result;
              if(this.pick.redBrPik == '0'){
                if(this.pick.brRunPik == '1') {
                  this.description = 'Pick: 1st round in year ' + this.pick.godPik;
                } else{
                  this.description = 'Pick: 2nd round in year ' + this.pick.godPik;
                }
              }else{
                this.description = 'Pick order: ' + this.pick.redBrPik + ' Round: ' + this.pick.brRunPik + ' Year:' + this.pick.godPik;
              }
            }
          });
        } else if(this.partnerTradeSubject.tipPredTrg == 2) { // Draft Rights
          this.tradesService.getDraftRightsByID(this.partnerTradeSubject.idIgrac!).subscribe({
            next: (result: DraftRight) => {
              this.draftRights = result;
              console.log('ma nek se npk')
            }
          });
        }
      }
    }
  }

  removeAssetButtonClicked(asset: any): void {
    this.removeAssetButtonState = 'clicked';
    setTimeout(() => { this.removeAssetButtonState = 'idle'; }, 200);

    console.log('prvo')

    if(this.chosenPlayer) {
      this.chosenPlayers.forEach((player, index) => {
        if(player.id === asset.id){
          this.chosenPlayers.splice(index,1);
          this.showNotification("Player successfully removed!");
        }
      });
    } else if(this.chosenPick) {
      this.chosenPicks.forEach((pick, index) => {
        if(pick.idPik === asset.idPik){
          this.chosenPicks.splice(index,1);
          this.showNotification("Pick successfully removed!");
        }
      });
    } else if(this.chosenDraftRight) {
      this.chosenDraftRights.forEach((draftRight, index) => {
        if(draftRight.idPrava === asset.idPrava){
          this.chosenDraftRights.splice(index,1);
          this.showNotification("Draft Right successfully removed!");
        }
      });
    }

    this.showNotification("Asset successfully removed!");
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }

  faPlus = faPlus;
  faMinus = faMinus;
}
