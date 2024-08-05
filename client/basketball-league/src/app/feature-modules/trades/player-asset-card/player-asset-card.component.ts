import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { Player } from 'src/app/shared/model/player.model';

@Component({
  selector: 'app-player-asset-card',
  templateUrl: './player-asset-card.component.html',
  styleUrls: ['./player-asset-card.component.css'],
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
export class PlayerAssetCardComponent implements OnInit{
  addAssetButtonState: string = 'idle';
  removeAssetButtonState: string = 'idle';
  assetSelected: boolean = false;
  @Input() player!: Player;
  @Input() chosenPlayers!: Player[];
  age: string = '';
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();

  constructor(private snackBar: MatSnackBar) { }

  ngOnInit(): void {
    const today = new Date();
    const birthDate = new Date(this.player.datRodj!);
    this.age = (today.getFullYear() - birthDate.getFullYear()).toString();

    if(this.chosenPlayers){
      this.chosenPlayers.forEach(player => {
        if(player.id === this.player.id){
          this.assetSelected = true;
        }
      })
    }
  }

  addAssetButtonClicked(asset: Player): void {
    this.addAssetButtonState = 'clicked';
    setTimeout(() => { this.addAssetButtonState = 'idle'; }, 200);
    
    let alreadyThere = false;
    this.chosenPlayers.forEach(player => {
        if(player.id === asset.id){
          alreadyThere = true
        }
      }
    )

    if(!alreadyThere){
      this.chosenPlayers.push(asset);
      this.assetSelected = true;
      this.showNotification("Player successfully added!");
    }
    else{
      this.showNotification("Player already chosen!");
    }
  }

  removeAssetButtonClicked(asset: Player): void {
    this.removeAssetButtonState = 'clicked';
    setTimeout(() => { this.removeAssetButtonState = 'idle'; }, 200);

    this.chosenPlayers.forEach( (player, index) => {
        if(player.id === asset.id){
          this.chosenPlayers.splice(index,1)
          this.assetSelected = false;
          this.showNotification("Player successfully removed!");      
        }
      }
    )  
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }

  faMinus = faMinus;
  faPlus = faPlus;
}
