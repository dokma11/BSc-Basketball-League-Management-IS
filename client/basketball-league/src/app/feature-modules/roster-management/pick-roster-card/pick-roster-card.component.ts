import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faList, faBan, faHeart, faWindowClose } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Pick } from 'src/app/shared/model/pick.model';
import { AddPlayerToListPromptComponent } from '../add-player-to-list-prompt/add-player-to-list-prompt.component';
import { RosterService } from '../roster.service';
import { WishlistAsset } from 'src/app/shared/model/wishlistAsset.model';

@Component({
  selector: 'app-pick-roster-card',
  templateUrl: './pick-roster-card.component.html',
  styleUrls: ['./pick-roster-card.component.css'],
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
export class PickRosterCardComponent implements OnInit{
  addToTradeListButtonState: string = 'idle';  
  addToUntouchablesListButtonState: string = 'idle';
  addToWishlistButtonState: string = 'idle';
  assetSelected: boolean = false;
  @Input() pick!: Pick;
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  @Input() ownTeam!: boolean;
  wishlistItems: WishlistAsset[] = [];
  onWishlist: boolean = false;

  constructor(private dialog: MatDialog,
              private authService: AuthService,
              private snackBar: MatSnackBar,
              private rosterService: RosterService) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    // Check if user is on the teams wishlist already
    if (!this.ownTeam){
      this.rosterService.getWishlistByTeamID(this.user?.teamId!).subscribe({
        next: (result: WishlistAsset) => {
          if (Array.isArray(result)){
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
  }

  addToWishlistButtonClicked(pick: any){
    this.addToWishlistButtonState = 'clicked';
    setTimeout(() => { this.addToWishlistButtonState = 'idle'; }, 200);
    
    if (!this.pick.nedodListPik) {
      this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
        data: {
          list: 'wishlist',
          pick: this.pick,
          action: 'add',
          teamId: this.user?.teamId,
        }
      });

      this.dialogRef.afterClosed().subscribe((result: any) => {
        this.rosterService.getWishlistByTeamID(this.user?.teamId!).subscribe({
          next: (result: WishlistAsset) => {
            if (Array.isArray(result)){
              this.wishlistItems = result;
    
              this.wishlistItems.forEach(asset => {
                if (asset.idPik == this.pick.idPik){
                  this.onWishlist = true;
                }
              });
            }
          }
        });
      });
    }
  }

  removeFromWishlistButtonClicked(pick: Pick) {
    this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
      data: {
        list: 'wishlist',
        pick: this.pick,
        action: 'remove',
        teamId: this.user?.teamId,
      }
    });

    this.dialogRef.afterClosed().subscribe((result: any) => {
      this.onWishlist = false;
      this.rosterService.getWishlistByTeamID(this.user?.teamId!).subscribe({
        next: (result: WishlistAsset) => {
          if (Array.isArray(result)){
            this.wishlistItems = result;
  
            this.wishlistItems.forEach(asset => {
              if (asset.idPik == this.pick.idPik){
                this.onWishlist = true;
              }
            });
          }
        }
      });
    });
  }

  addToTradeListButtonClicked(pick: Pick) {
    this.addToUntouchablesListButtonState = 'clicked';
    setTimeout(() => { this.addToUntouchablesListButtonState = 'idle'; }, 200);
    
    if(!this.pick.nedodListPik){
      this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
        data: {
          list: 'trade list',
          pick: this.pick, 
          action: 'add',
        }
      });
    } else {
      this.showNotification('Can not add an untouchable pick to the trade list!')
    }
  }

  removeFromTradeListButtonClicked(pick: any){
    this.addToTradeListButtonState = 'clicked';
    setTimeout(() => { this.addToTradeListButtonState = 'idle'; }, 200);

    this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
      data: {
        list: 'trade list',
        pick: this.pick, 
        action: 'remove',
      }
    });
  }

  addToUntouchablesListButtonClicked(pick: Pick) {
    this.addToTradeListButtonState = 'clicked';
    setTimeout(() => { this.addToTradeListButtonState = 'idle'; }, 200);
    
    if(!this.pick.trgListPik){
      this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
        data: {
          list: 'untouchable list',
          pick: this.pick,
          action: 'add',
        }
      });
    } else {
      this.showNotification('Can not add a tradeable pick to the untouchables list!')
    }
  }

  removeFromUntouchablesListButtonClicked(pick: any){
    this.addToUntouchablesListButtonState = 'clicked';
    setTimeout(() => { this.addToUntouchablesListButtonState = 'idle'; }, 200);
    
    this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
      data: {
        list: 'untouchable list',
        pick: this.pick, 
        action: 'remove',
      }
    });
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }

  faList = faList;
  faBan = faBan;
  faHeart = faHeart;
  faWindowClose = faWindowClose;
}
