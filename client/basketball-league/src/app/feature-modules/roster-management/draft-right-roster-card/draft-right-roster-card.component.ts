import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { faHeart, faBan, faList } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Team } from 'src/app/shared/model/team.model';
import { AddPlayerToListPromptComponent } from '../add-player-to-list-prompt/add-player-to-list-prompt.component';
import { RosterService } from '../roster.service';
import { DraftRight } from 'src/app/shared/model/draftRight.model';

@Component({
  selector: 'app-draft-right-roster-card',
  templateUrl: './draft-right-roster-card.component.html',
  styleUrls: ['./draft-right-roster-card.component.css']
})
export class DraftRightRosterCardComponent implements OnInit{
  addToWishlistButtonState: string = 'idle';
  addToUntouchablesListButtonState: string = 'idle';
  addToTradeListButtonState: string = 'idle';
  @Input() draftRight!: DraftRight;
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  playerBirthDayDate: string = '';
  playerBirthDayTime: string = '';
  timovi: Team[] = [];

  constructor(private dialog: MatDialog,
              private authService: AuthService,
              private rosterService: RosterService) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  }

  ngOnInit(): void {
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
