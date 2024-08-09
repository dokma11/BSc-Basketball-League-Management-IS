import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Pick } from 'src/app/shared/model/pick.model';
import { WishlistAsset } from 'src/app/shared/model/wishlistAsset.model';
import { RosterService } from '../../roster-management/roster.service';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';

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
  wishlistItems: WishlistAsset[] = [];
  onWishlist: boolean = false;
  
  constructor(private snackBar: MatSnackBar,
              private rosterService: RosterService,
              private authService: AuthService) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    if(this.chosenPicks){
      this.chosenPicks.forEach(team => {
        if(team.idPik === this.pick.idPik){
          this.assetSelected = true;
        }
      })
    }

    if (this.pick.idTim == this.user?.teamId) {
      this.ownTeam = true;
    }

    this.rosterService.getWishlistByTeamID(this.user?.teamId!).subscribe({
      next: (result: WishlistAsset[] | WishlistAsset) => {
        if (Array.isArray(result)) {
          this.wishlistItems = result;

          this.wishlistItems.forEach(asset => {
            if (asset.idPik == this.pick.idPik){
              this.onWishlist = true;
            }
          });
        }
      }
    });
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
