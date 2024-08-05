import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { Player } from 'src/app/shared/model/player.model';
import { WishlistAsset } from 'src/app/shared/model/wishlistAsset.model';
import { RosterService } from '../../roster-management/roster.service';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Team } from 'src/app/shared/model/team.model';

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
  wishlistItems: WishlistAsset[] = [];
  onWishlist: boolean = false;
  user: User | undefined;
  @Input() team!: Team;
  ownTeam: boolean = false;

  constructor(private snackBar: MatSnackBar,
              private rosterService: RosterService,
              private authService: AuthService) { 
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  }

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

    if (this.team.idTim == this.user?.teamId) {
      this.ownTeam = true;
    }

    this.rosterService.getWishlistByTeamID(this.user?.teamId!).subscribe({
      next: (result: WishlistAsset[] | WishlistAsset) => {
        if (Array.isArray(result)) {
          this.wishlistItems = result;

          this.wishlistItems.forEach(asset => {
            if (asset.idIgrac == this.player.id){
              this.onWishlist = true;
            }
          });
        }
      }
    });
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
