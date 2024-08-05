import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { faHeart, faBan, faList } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Team } from 'src/app/shared/model/team.model';
import { AddPlayerToListPromptComponent } from '../add-player-to-list-prompt/add-player-to-list-prompt.component';
import { RosterService } from '../roster.service';
import { DraftRight } from 'src/app/shared/model/draftRight.model';
import { Recruit } from 'src/app/shared/model/recruit.model';
import { MatSnackBar } from '@angular/material/snack-bar';
import { trigger, transition, style, animate, state } from '@angular/animations';

@Component({
  selector: 'app-draft-right-roster-card',
  templateUrl: './draft-right-roster-card.component.html',
  styleUrls: ['./draft-right-roster-card.component.css'],
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
export class DraftRightRosterCardComponent implements OnInit{
  addToWishlistButtonState: string = 'idle';
  addToUntouchablesListButtonState: string = 'idle';
  addToTradeListButtonState: string = 'idle';
  @Input() draftRight!: DraftRight;
  draftRightPlayer: Recruit | undefined;
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  age: string = '';
  @Input() ownTeam!: boolean;

  constructor(private dialog: MatDialog,
              private authService: AuthService,
              private rosterService: RosterService,
              private snackBar: MatSnackBar) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    this.rosterService.getRecruitById(this.draftRight.idRegrut).subscribe({
      next: (result: Recruit) => {
        this.draftRightPlayer = result;

        const today = new Date();
        const birthDate = new Date(this.draftRightPlayer.datRodj!);
        this.age = (today.getFullYear() - birthDate.getFullYear()).toString();
      }
    });
  }

  addToWishlistButtonClicked(player: any){
    // TODO: Implementirati lpogiku za dodavanje odredjenog igraca na listu zelja, vrv treba u samom modalnom da se to uradi

    this.addToWishlistButtonState = 'clicked';
    setTimeout(() => { this.addToWishlistButtonState = 'idle'; }, 200);
    this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
      data: 'wishlist'
    });
  }

  addToUntouchablesListButtonClicked(player: any){
    this.addToUntouchablesListButtonState = 'clicked';
    setTimeout(() => { this.addToUntouchablesListButtonState = 'idle'; }, 200);
    
    if(!this.draftRight.nedodListPrava){
      this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
        data: {
          list: 'untouchable list',
          draftRights: this.draftRight,
          action: 'add',
        }
      });
    } else {
      this.showNotification('Can not add a tradeable draft right to the trade list!')
    }
  }
  
  removeFromUntouchablesListButtonClicked(player: any){
    this.addToUntouchablesListButtonState = 'clicked';
    setTimeout(() => { this.addToUntouchablesListButtonState = 'idle'; }, 200);
    
    this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
      data: {
        list: 'untouchable list',
        draftRights: this.draftRight,
        action: 'remove',
      }
    });
  }

  addToTradeListButtonClicked(player: any){
    this.addToTradeListButtonState = 'clicked';
    setTimeout(() => { this.addToTradeListButtonState = 'idle'; }, 200);
    
    if(!this.draftRight.nedodListPrava){
      this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
        data: {
          list: 'trade list',
          draftRights: this.draftRight,
          action: 'add',
        }
      });
    } else {
      this.showNotification('Can not add an untouchable draft right to the trade list!')
    }
  }

  removeFromTradeListButtonClicked(player: any){
    this.addToTradeListButtonState = 'clicked';
    setTimeout(() => { this.addToTradeListButtonState = 'idle'; }, 200);
    
    this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
      data: {
        list: 'trade list',
        draftRights: this.draftRight,
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
}
