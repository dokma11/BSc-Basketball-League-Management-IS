import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { faHeart, faList, faBan, faWindowClose } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { AddPlayerToListPromptComponent } from '../add-player-to-list-prompt/add-player-to-list-prompt.component';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { Player } from 'src/app/shared/model/player.model';
import { MatSnackBar } from '@angular/material/snack-bar';
import { RosterService } from '../roster.service';
import { WishlistAsset } from 'src/app/shared/model/wishlistAsset.model';

@Component({
  selector: 'app-player-card',
  templateUrl: './player-card.component.html',
  styleUrls: ['./player-card.component.css'],
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
export class PlayerCardComponent implements OnInit{
  addToWishlistButtonState: string = 'idle';
  addToUntouchablesListButtonState: string = 'idle';
  addToTradeListButtonState: string = 'idle';
  @Input() player!: Player;
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  age: string = '';
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
    const today = new Date();
    const birthDate = new Date(this.player.datRodj!);
    this.age = (today.getFullYear() - birthDate.getFullYear()).toString();

    // Check if user is on the teams wishlist already
    if (!this.ownTeam){
      this.rosterService.getWishlistByTeamID(this.user?.teamId!).subscribe({
        next: (result: WishlistAsset) => {
          if (Array.isArray(result)){
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
  }

  addToWishlistButtonClicked(player: any){
    this.addToWishlistButtonState = 'clicked';
    setTimeout(() => { this.addToWishlistButtonState = 'idle'; }, 200);
    
    if (!this.player.nedodListIgr){
      this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
        data: {
          list: 'wishlist',
          player: this.player,
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
                if (asset.idIgrac == this.player.id){
                  this.onWishlist = true;
                }
              });
            }
          }
        });
      });
    } else {
      this.showNotification('Can not add an untouchable player to the wishlist!');
    }
  }

  removeFromWishlistButtonClicked(player: any){
    this.addToWishlistButtonState = 'clicked';
    setTimeout(() => { this.addToWishlistButtonState = 'idle'; }, 200);

    this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
      data: {
        list: 'wishlist',
        player: this.player,
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
              if (asset.idIgrac == this.player.id){
                this.onWishlist = true;
              }
            });
          }
        }
      });
    });
  }

  addToUntouchablesListButtonClicked(player: any){
    this.addToUntouchablesListButtonState = 'clicked';
    setTimeout(() => { this.addToUntouchablesListButtonState = 'idle'; }, 200);
    
    if(!this.player.trgListIgr){
      this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
        data: {
          list: 'untouchable list',
          player: this.player,
          action: 'add',
        }
      });
    } else {
      this.showNotification('Can not add a tradeable player to the untouchables list!')
    }
  }
  
  removeFromUntouchablesListButtonClicked(player: any){
    this.addToUntouchablesListButtonState = 'clicked';
    setTimeout(() => { this.addToUntouchablesListButtonState = 'idle'; }, 200);
    
    this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
      data: {
        list: 'untouchable list',
        player: this.player, 
        action: 'remove',
      }
    });
  }

  addToTradeListButtonClicked(player: any){
    this.addToTradeListButtonState = 'clicked';
    setTimeout(() => { this.addToTradeListButtonState = 'idle'; }, 200);
    
    if(!this.player.nedodListIgr){
      this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
        data: {
          list: 'trade list',
          player: this.player, 
          action: 'add',
        }
      });
    } else {
      this.showNotification('Can not add an untouchable player to the trade list!')
    }
  }

  removeFromTradeListButtonClicked(player: any){
    this.addToTradeListButtonState = 'clicked';
    setTimeout(() => { this.addToTradeListButtonState = 'idle'; }, 200);

    this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
      data: {
        list: 'trade list',
        player: this.player, 
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

  faHeart = faHeart;
  faBan = faBan;
  faList = faList;
  faWindowClose = faWindowClose;
}
