import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { faCheck, faTimes, faPen, faTrash } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';

@Component({
  selector: 'app-news-card',
  templateUrl: './news-card.component.html',
  styleUrls: ['./news-card.component.css']
})
export class NewsCardComponent implements OnInit{
  acceptButtonState: string = 'idle';
  declineButtonState: string = 'idle';
  detailsButtonState: string = 'idle';
  seeExplanationButtonState: string = 'idle';
  cancelRequestButtonState: string = 'idle';
  newsHeadline: string = "";  // DOK NE POVEZEM SA BEKOM
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
    this.newsHeadline = "Headline please";
  }

  faCheck = faCheck;
  faTimes = faTimes;
  faPen = faPen;
  faTrash = faTrash;
}
