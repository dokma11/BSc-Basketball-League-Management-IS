import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { DraftRight } from 'src/app/shared/model/draftRight.model';
import { Pick } from 'src/app/shared/model/pick.model';
import { Player } from 'src/app/shared/model/player.model';

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
  @Input() chosenPlayers!: Player[];
  @Input() chosenPicks!: Pick[];
  @Input() chosenDraftRights!: DraftRight[];

  constructor(private snackBar: MatSnackBar) {
  
  }

  ngOnInit(): void {
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
