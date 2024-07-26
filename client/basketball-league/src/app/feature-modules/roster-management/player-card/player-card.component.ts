import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { faHeart, faList, faBan } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { AddPlayerToListPromptComponent } from '../add-player-to-list-prompt/add-player-to-list-prompt.component';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { Player } from 'src/app/shared/model/player.model';

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

  constructor(private dialog: MatDialog,
              private authService: AuthService) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    const today = new Date();
    const birthDate = new Date(this.player.datRodj!);
    this.age = (today.getFullYear() - birthDate.getFullYear()).toString();
    // TODO: Implementirati logiku za proveru koji tim je u pitanju i samim time koje dugmice prikazati
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
    // TODO: Implementirati lpogiku za dodavanje odredjenog igraca na listu zelja, vrv treba u samom modalnom da se to uradi

    this.addToUntouchablesListButtonState = 'clicked';
    setTimeout(() => { this.addToUntouchablesListButtonState = 'idle'; }, 200);
    this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
      data: 'untouchable list'
    });
  }
  
  addToTradeListButtonClicked(player: any){
    // TODO: Implementirati lpogiku za dodavanje odredjenog igraca na listu zelja, vrv treba u samom modalnom da se to uradi

    this.addToTradeListButtonState = 'clicked';
    setTimeout(() => { this.addToTradeListButtonState = 'idle'; }, 200);
    this.dialogRef = this.dialog.open(AddPlayerToListPromptComponent, {
      data: 'trade list'
    });
  }

  faHeart = faHeart;
  faBan = faBan;
  faList = faList;
}
