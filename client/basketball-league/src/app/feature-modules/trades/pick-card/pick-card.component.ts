import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { faHeart, faHandPaper, faList, faPlus } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';

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
  addToWishlistButtonState: string = 'idle';
  addToUntouchablesListButtonState: string = 'idle';
  addToTradeListButtonState: string = 'idle';
  player: string = 'IGRAC';  // DOK NE POVEZEM SA BEKOM
  //@Input() request!: PersonalTourRequest; Ovde treba da budu pikovi kao lista
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  ownTeam: boolean = false;

  constructor(private dialog: MatDialog,
              private authService: AuthService) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  
    // Treba nekako videti za koji tim se traze pikovi i onda ih ucitati u this.picks recimo
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

  addAssetButtonClicked(assset: any): void {
    this.addAssetButtonState = 'clicked';
    setTimeout(() => { this.addAssetButtonState = 'idle'; }, 200);

  }

  addToWishlistButtonClicked(player: any){

  }

  addToUntouchablesListButtonClicked(player: any){

  }
  
  addToTradeListButtonClicked(player: any){

  }

  faHeart = faHeart;
  faHandPaper = faHandPaper;
  faList = faList;
  faPlus = faPlus
}
