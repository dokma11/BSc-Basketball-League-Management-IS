import { trigger, state, style, transition, animate } from '@angular/animations';
import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { faCheck, faTimes, faPen, faTrash, faCircleInfo, faBan } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { DeclineRequestPromptComponent } from '../decline-request-prompt/decline-request-prompt.component';
import { AcceptRequestPromptComponent } from '../accept-request-prompt/accept-request-prompt.component';
import { SeeDenialExplanationPromptComponent } from '../see-denial-explanation-prompt/see-denial-explanation-prompt.component';
import { ShowRequestDetailsPromptComponent } from '../show-request-details-prompt/show-request-details-prompt.component';

@Component({
  selector: 'app-trade-request-card',
  templateUrl: './trade-request-card.component.html',
  styleUrls: ['./trade-request-card.component.css'],
  animations: [
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
export class TradeRequestCardComponent implements OnInit{
  acceptButtonState: string = 'idle';
  declineButtonState: string = 'idle';
  detailsButtonState: string = 'idle';
  seeExplanationButtonState: string = 'idle';
  cancelRequestButtonState: string = 'idle';
  request: string = 'das';  // DOK NE POVEZEM SA BEKOM
  //@Input() request!: PersonalTourRequest;
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  tourOccurrenceTime: string = "";
  tourOccurrenceDate: string = "";
  exhibitionsString: string = "";

  constructor(private dialog: MatDialog,
              private authService: AuthService) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    // TODO: Uraaditi sta treba vec za ngOnInit 

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

  acceptButtonClicked(request: any) {
    // TODO: Dodati mali prozor za prihvatanje zahteva za trgovinu 

    this.acceptButtonState = 'clicked';
    setTimeout(() => { this.acceptButtonState = 'idle'; }, 200);
    this.dialogRef = this.dialog.open(AcceptRequestPromptComponent, {
      data: request
    });
    this.dialogRef.afterClosed().subscribe((result: any) => {
      this.dialogRefClosed.emit(result);
    });
  }

  declineButtonClicked(request: any) {
    // TODO: Dodati mali prozor za potvrdu odbijanja 

    this.declineButtonState = 'clicked';
    setTimeout(() => { this.declineButtonState = 'idle'; }, 200);
    this.dialogRef = this.dialog.open(DeclineRequestPromptComponent, {
      data: request
    });
    this.dialogRef.afterClosed().subscribe((result: any) => {
      this.dialogRefClosed.emit(result);
    });
  }

  seeDetailsButtonClicked(request: any){
    // TODO: Dodati dijalog u kojem ce lepo biti ispisano sve sto treba o trgovini

    this.detailsButtonState = 'clicked';
    setTimeout(() => { this.detailsButtonState = 'idle'; }, 200);
    this.dialogRef = this.dialog.open(ShowRequestDetailsPromptComponent, {
      data: request
    });
  }

  seeExplanationButtonClicked(request: any) {
    // TODO: Dodati manji dijalog koji ce ispisati razlog odbijanja trgovine
    
    this.seeExplanationButtonState = 'clicked';
    setTimeout(() => { this.seeExplanationButtonState = 'idle'; }, 200);
    this.dialogRef = this.dialog.open(SeeDenialExplanationPromptComponent, {
      data: request
    });
  }

  cancelRequestButtonClicked(request: any){
    // TODO: Proveriti da li ovo treba da postoji, kontam da nije problem dodati
    // Ako postoji onda samo napraviti mali prozor slicno kao na IIS za potvrdu odbijanja

    // this.cancelRequestButtonState = 'clicked';
    // setTimeout(() => { this.cancelRequestButtonState = 'idle'; }, 200);

    // this.toursService.cancelTourRequest(request.id!).subscribe({
    //   next: (response: any) => {
    //     this.dialogRefClosed.emit(response);
    //   }
    // })
  }

  faCheck = faCheck;
  faTimes = faTimes;
  faPen = faPen;
  faTrash = faTrash;
  faCircleInfo = faCircleInfo;
  faBan = faBan;
}
