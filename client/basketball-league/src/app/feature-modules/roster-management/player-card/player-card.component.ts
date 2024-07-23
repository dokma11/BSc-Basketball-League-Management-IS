import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { faHeart, faHandPaper, faList, faBan } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { AddPlayerToListPromptComponent } from '../add-player-to-list-prompt/add-player-to-list-prompt.component';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { RosterService } from '../roster.service';
import { Igrac } from '../model/igrac.model';
import { Team } from 'src/app/shared/model/team.model';

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
  player: string = 'IGRAC';  // DOK NE POVEZEM SA BEKOM
  //@Input() request!: PersonalTourRequest;
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();

  timovi: Team[] = [];

  constructor(private dialog: MatDialog,
              private authService: AuthService,
              private rosterService: RosterService) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  }

  ngOnInit(): void {
  //   const tourOccurrenceDateTimeString = this.request.occurrenceDateTime.toString();
  //   [this.tourOccurrenceDate, this.tourOccurrenceTime] = tourOccurrenceDateTimeString.split('T');

  //   if(this.request.proposerId){
  //     this.toursService.getGuestById(this.request.proposerId).subscribe({
  //       next : (result: Guest) => {
  //         this.request.proposer = result;
  //       }
  //     });
  //   }

  //   this.request.exhibitions!.forEach((exhibition: Exhibition) => {
  //     this.exhibitionsString += exhibition.name + ", ";
  //   });

  //   this.exhibitionsString = this.exhibitionsString.slice(0, -2);

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

    // Cisto proba da vidim da li je dobro povezano
    this.rosterService.getAllTeams().subscribe({
      next: (result: Team[] | Team) => {
        console.log('Usao u next')
        if(Array.isArray(result)){
          console.log('Usao u if')
          this.timovi = result;
          console.log(this.timovi);

          for (let i = 0; i < this.timovi.length; i++) {
            console.log('id: ' + this.timovi[i].idTim);
            console.log('lokacija: ' + this.timovi[i].lokTim);
            console.log('naziv: ' + this.timovi[i].nazTim);
            console.log('godosn: ' + this.timovi[i].godOsnTim);
          }
        }
        console.log('nesto')
      }
    })
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
