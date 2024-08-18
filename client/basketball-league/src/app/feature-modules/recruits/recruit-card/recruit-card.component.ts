import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { faClipboard, faDumbbell, faHeart, faWindowClose } from '@fortawesome/free-solid-svg-icons';
import { RecruitsService } from '../recruits.service';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Recruit } from 'src/app/shared/model/recruit.model';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { AddPlayerToListPromptComponent } from '../../roster-management/add-player-to-list-prompt/add-player-to-list-prompt.component';
import { InterviewInvitePromptComponent } from '../interview-invite-prompt/interview-invite-prompt.component';
import { TrainingInvitePromptComponent } from '../training-invite-prompt/training-invite-prompt.component';

@Component({
  selector: 'app-recruit-card',
  templateUrl: './recruit-card.component.html',
  styleUrls: ['./recruit-card.component.css'],
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
export class RecruitCardComponent implements OnInit{
  addToWishlistButtonState: string = '';
  inviteToTrainingButtonState: string = '';
  inviteToInterviewButtonState: string = '';
  user: User | undefined;
  @Input() player!: Recruit;
  onWishlist: boolean = false;
  age: string = '';
  private dialogRef: any;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();

  constructor(private authService: AuthService,
              private dialog: MatDialog,
              private snackBar: MatSnackBar,
              private recruitsService: RecruitsService) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    
  }

  addToWishlistButtonClicked(recruit: any) {
    this.addToWishlistButtonState = 'clicked';
    setTimeout(() => { this.addToWishlistButtonState = 'idle'; }, 200);
    
    this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
      data: {
        list: 'wishlist',
        player: this.player,  // ovde samo staviti regrut i to promeniti
        action: 'add',
        teamId: this.user?.teamId,
      }
    });
    
  }

  removeFromWishlistButtonClicked(recruit: any) {
    this.addToWishlistButtonState = 'clicked';
    setTimeout(() => { this.addToWishlistButtonState = 'idle'; }, 200);

    this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
      data: {
        list: 'wishlist',
        player: this.player, // ovde samo staviti regrut i to promeniti
        action: 'remove',
        teamId: this.user?.teamId,
      }
    });
  }

  inviteToInterviewButtonClicked(recruit: any) {
    this.inviteToInterviewButtonState = 'clicked';
    setTimeout(() => { this.inviteToInterviewButtonState = 'idle'; }, 200);

    this.dialogRef = this.dialog.open(InterviewInvitePromptComponent, {
      data: {
        player: this.player, // ovde samo staviti regrut i to promeniti
        teamId: this.user?.teamId, // mozda ovde da bude kao korisnik id da se zna koji trener je u pitanju
      }
    });
  }

  inviteToTrainingButtonClicked(recruit: any) {
    this.inviteToTrainingButtonState = 'clicked';
    setTimeout(() => { this.inviteToTrainingButtonState = 'idle'; }, 200);

    this.dialogRef = this.dialog.open(TrainingInvitePromptComponent, {
      data: {
        player: this.player, // ovde samo staviti regrut i to promeniti
        teamId: this.user?.teamId, // mozda ovde da bude kao korisnik id da se zna koji trener je u pitanju
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
  faClipboard = faClipboard;
  faDumbbell = faDumbbell;
  faWindowClose = faWindowClose;
}
