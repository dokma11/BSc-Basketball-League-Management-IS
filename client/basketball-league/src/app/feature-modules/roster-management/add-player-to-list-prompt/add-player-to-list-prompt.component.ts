import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { RosterService } from '../roster.service';
import { Player } from 'src/app/shared/model/player.model';
import { Pick } from 'src/app/shared/model/pick.model';
import { DraftRight } from 'src/app/shared/model/draftRight.model';
import { Recruit } from 'src/app/shared/model/recruit.model';

@Component({
  selector: 'app-add-player-to-list-prompt',
  templateUrl: './add-player-to-list-prompt.component.html',
  styleUrls: ['./add-player-to-list-prompt.component.css'],
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
export class AddPlayerToListPromptComponent {
  cancelButtonState: string = 'idle';
  addButtonState: string = 'idle';
  focused: string = '';
  list: string = ''; 
  action: string = '';
  player: Player | undefined;
  pick: Pick | undefined;
  draftRights: DraftRight | undefined;
  recruit: Recruit | undefined;
  loggedInUserTeamId: number | undefined;

  constructor(private rosterService: RosterService,
              private snackBar: MatSnackBar,
              private dialogRef: MatDialogRef<AddPlayerToListPromptComponent>,
              @Inject(MAT_DIALOG_DATA) public data: any) {
    this.list = data.list;
    this.action = data.action;
    if (data.player) {
      this.player = data.player;
    }
    if (data.pick) {
      this.pick = data.pick;
    }
    if (data.draftRights) {
      this.draftRights = data.draftRights;
    }
    if (data.recruit) {
      this.recruit = data.recruit;
    }
    if (data.teamId != undefined) {
      this.loggedInUserTeamId = data.teamId;
    }
  }

  addButtonClicked(){
      this.addButtonState = 'clicked';
      setTimeout(() => { this.addButtonState = 'idle'; }, 200);

      if(this.list == 'untouchable list' && this.action == 'add' && this.player){
        this.player!.nedodListIgr = true;
        this.rosterService.updatePlayer(this.player!).subscribe({
          next: (result: Player) => {
            this.dialogRef.close();
            this.showNotification('Player successfully added to the untouchables list!');
          }
        })
      } else if(this.list == 'untouchable list' && this.action == 'remove' && this.player){
        this.player!.nedodListIgr = false;
        this.rosterService.updatePlayer(this.player!).subscribe({
          next: (result: Player) => {
            this.dialogRef.close();
            this.showNotification('Player successfully removed from the untouchables list!');
          }
        })
      } else if(this.list == 'trade list' && this.action == 'add' && this.player){
        this.player!.trgListIgr = true;
        this.rosterService.updatePlayer(this.player!).subscribe({
          next: (result: Player) => {
            this.dialogRef.close();
            this.showNotification('Player successfully added to the trade list!');
          }
        })
      } else if(this.list == 'trade list' && this.action == 'remove' && this.player){
        this.player!.trgListIgr = false;
        this.rosterService.updatePlayer(this.player!).subscribe({
          next: (result: Player) => {
            this.dialogRef.close();
            this.showNotification('Player successfully removed from the trade list!');
          }
        })
      }  else if(this.list == 'wishlist' && this.action == 'add' && this.player){
        console.log(this.loggedInUserTeamId);
        this.rosterService.addPlayerToWishlist(this.player!, this.loggedInUserTeamId!).subscribe({
          next: (result: Player) => {
            this.dialogRef.close();
            this.showNotification('Player successfully added to the wishlist!');
          }
        })
      } else if(this.list == 'wishlist' && this.action == 'remove' && this.player){
        this.rosterService.removePlayerFromWishlist(this.player!, this.loggedInUserTeamId!).subscribe({
          next: (result: Player) => {
            this.dialogRef.close();
            this.showNotification('Player successfully removed from the wishlist!');
          }
        })
      } else if(this.list == 'untouchable list' && this.action == 'add' && this.pick){
        this.pick!.nedodListPik = true;
        this.rosterService.updatePick(this.pick!).subscribe({
          next: (result: Pick) => {
            this.dialogRef.close();
            this.showNotification('Pick successfully added to the untouchables list!');
          }
        })
      } else if(this.list == 'untouchable list' && this.action == 'remove' && this.pick){
        this.pick!.nedodListPik = false;
        this.rosterService.updatePick(this.pick!).subscribe({
          next: (result: Pick) => {
            this.dialogRef.close();
            this.showNotification('Pick successfully removed from the untouchables list!');
          }
        })
      } else if(this.list == 'trade list' && this.action == 'add' && this.pick){
        this.pick!.trgListPik = true;
        this.rosterService.updatePick(this.pick!).subscribe({
          next: (result: Pick) => {
            this.dialogRef.close();
            this.showNotification('Pick successfully added to the trade list!');
          }
        })
      } else if(this.list == 'trade list' && this.action == 'remove' && this.pick){
        this.pick!.trgListPik = false;
        this.rosterService.updatePick(this.pick!).subscribe({
          next: (result: Pick) => {
            this.dialogRef.close();
            this.showNotification('Pick successfully removed form the trade list!');
          }
        })
      } else if(this.list == 'wishlist' && this.action == 'add' && this.pick){
        this.rosterService.addPickToWishlist(this.pick!, this.loggedInUserTeamId!).subscribe({
          next: (result: Pick) => {
            this.dialogRef.close();
            this.showNotification('Pick successfully added to the wishlist!');
          }
        })
      } else if(this.list == 'wishlist' && this.action == 'remove' && this.pick){
        this.rosterService.removePickFromWishlist(this.pick!, this.loggedInUserTeamId!).subscribe({
          next: (result: Pick) => {
            this.dialogRef.close();
            this.showNotification('Pick successfully removed from the wishlist!');
          }
        })
      } else if(this.list == 'untouchable list' && this.action == 'add' && this.draftRights){
        this.draftRights!.nedodListPrava = true;
        this.rosterService.updateDraftRights(this.draftRights!).subscribe({
          next: (result: DraftRight) => {
            this.dialogRef.close();
            this.showNotification('Draft Rights successfully added to the untouchables list!');
          }
        })
      } else if(this.list == 'untouchable list' && this.action == 'remove' && this.draftRights){
        this.draftRights!.nedodListPrava = false;
        this.rosterService.updateDraftRights(this.draftRights!).subscribe({
          next: (result: DraftRight) => {
            this.dialogRef.close();
            this.showNotification('Draft Rights successfully removed from the untouchables list!');
          }
        })
      } else if(this.list == 'wishlist' && this.action == 'add' && this.draftRights){
        this.rosterService.addDraftRightsToWishlist(this.draftRights!, this.loggedInUserTeamId!).subscribe({
          next: (result: DraftRight) => {
            this.dialogRef.close();
            this.showNotification('Draft Rights successfully added to the wishlist!');
          }
        })
      } else if(this.list == 'wishlist' && this.action == 'remove' && this.draftRights){
        this.rosterService.removeDraftRightsFromWishlist(this.draftRights!, this.loggedInUserTeamId!).subscribe({
          next: (result: DraftRight) => {
            this.dialogRef.close();
            this.showNotification('Draft Rights successfully removed form the wishlist!');
          }
        })
      } else if(this.list == 'trade list' && this.action == 'add' && this.draftRights){
        this.draftRights!.trgListPrava = true;
        this.rosterService.updateDraftRights(this.draftRights!).subscribe({
          next: (result: DraftRight) => {
            this.dialogRef.close();
            this.showNotification('Draft Rights successfully added to the trade list!');
          }
        })
      } else if(this.list == 'trade list' && this.action == 'remove' && this.draftRights){
        this.draftRights!.trgListPrava = false;
        this.rosterService.updateDraftRights(this.draftRights!).subscribe({
          next: (result: DraftRight) => {
            this.dialogRef.close();
            this.showNotification('Draft Rights successfully removed form the trade list!');
          }
        })
      } else if(this.list == 'wishlist' && this.action == 'add' && this.recruit){
        console.log(this.loggedInUserTeamId);
        this.rosterService.addRecruitToWishlist(this.recruit!, this.loggedInUserTeamId!).subscribe({
          next: (result: Recruit) => {
            this.dialogRef.close();
            this.showNotification('Recruit successfully added to the wishlist!');
          }
        })
      } else if(this.list == 'wishlist' && this.action == 'remove' && this.recruit){
        this.rosterService.removeRecruitFromWishlist(this.recruit!, this.loggedInUserTeamId!).subscribe({
          next: (result: Recruit) => {
            this.dialogRef.close();
            this.showNotification('Recruit successfully removed from the wishlist!');
          }
        })
      } 
  }

  cancelButtonClicked(){
    this.cancelButtonState = 'clicked';
    setTimeout(() => { this.cancelButtonState = 'idle'; }, 200);
    this.dialogRef.close();
  }

  overviewClicked(){
    this.dialogRef.close();
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }
}
