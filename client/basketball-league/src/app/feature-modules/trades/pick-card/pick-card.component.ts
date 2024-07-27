import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Pick } from 'src/app/shared/model/pick.model';

@Component({
  selector: 'app-pick-card',
  templateUrl: './pick-card.component.html',
  styleUrls: ['./pick-card.component.css'],
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
export class PickCardComponent implements OnInit{
  addAssetButtonState: string = 'idle';  
  removeAssetButtonState: string = 'idle';
  assetSelected: boolean = false;
  @Input() pick!: Pick;
  @Input() chosenPicks!: Pick[];
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  ownTeam: boolean = false;

  constructor(private snackBar: MatSnackBar) {
    
  }

  ngOnInit(): void {
    if(this.chosenPicks){
      this.chosenPicks.forEach(team => {
        if(team.idPik === this.pick.idPik){
          this.assetSelected = true;
        }
      })
    }
  }

  addAssetButtonClicked(asset: Pick): void {
    this.addAssetButtonState = 'clicked';
    setTimeout(() => { this.addAssetButtonState = 'idle'; }, 200);

    let alreadyThere = false;
    this.chosenPicks.forEach(team => {
        if(team.idPik === asset.idPik){
          alreadyThere = true
        }
      }
    )

    if(!alreadyThere){
      this.chosenPicks.push(asset);
      this.assetSelected = true;
      this.showNotification("Pick successfully added!");
    }
    else{
      this.showNotification("Pick already chosen!");
    }
  }

  removeAssetButtonClicked(asset: Pick): void {
    this.removeAssetButtonState = 'clicked';
    setTimeout(() => { this.removeAssetButtonState = 'idle'; }, 200);

    this.chosenPicks.forEach( (team, index) => {
        if(team.idPik === asset.idPik){
          this.chosenPicks.splice(index,1)
          this.assetSelected = false;
          this.showNotification("Pick successfully removed!");
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

  faPlus = faPlus;
  faMinus = faMinus;
}
