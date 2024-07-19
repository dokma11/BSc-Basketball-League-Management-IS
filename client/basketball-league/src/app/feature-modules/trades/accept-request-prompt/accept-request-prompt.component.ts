import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { DeclineRequestPromptComponent } from '../decline-request-prompt/decline-request-prompt.component';
import { trigger, transition, style, animate, state } from '@angular/animations';

@Component({
  selector: 'app-accept-request-prompt',
  templateUrl: './accept-request-prompt.component.html',
  styleUrls: ['./accept-request-prompt.component.css'],
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
export class AcceptRequestPromptComponent {
  cancelButtonState: string = 'idle';
  acceptButtonState: string = 'idle';
  focused: string = '';
  private ownDialogRef: any;
  adultTicketPrice: string = "0";
  minorTicketPrice: string = "0";
  private request: string = ''; // OVO JE ZAMENA DOK NE DODJE POVEZIVANJE SA BEKOM

  constructor(private snackBar: MatSnackBar,
              private dialogRef: MatDialogRef<DeclineRequestPromptComponent>,
              @Inject(MAT_DIALOG_DATA) public data: any,
              private dialog: MatDialog,) {
    this.request = data;
  }

  acceptButtonClicked(){
      this.acceptButtonState = 'clicked';
      setTimeout(() => { this.acceptButtonState = 'idle'; }, 200);

      // TODO: Uraditi logiku za prihvatanje zahteva

      // this.request!.status = PersonalTourRequestStatus.ACCEPTED;

      // this.toursService.updateTourRequest(this.request!).subscribe({
      //   next: () => {

      //     const tour: PersonalTour = {
      //       occurrenceDateTime: this.request!.occurrenceDateTime || new Date(),
      //       adultTicketPrice: this.adultTicketPrice || "",
      //       minorTicketPrice: this.minorTicketPrice || "",
      //       guestNumber: this.request!.guestNumber || "",
      //       proposerId: this.request!.proposerId || 0,
      //       guideId: this.selectedCurator[0].id || 7,
      //       duration: "0",
      //       exhibitions: this.request?.exhibitions || []
      //     };

      //     this.toursService.addPersonalTour(tour).subscribe({
      //       next: () => {
      //         this.showNotification('Tour request successfully accepted!');
      //         this.dialogRef.close();
      //       }
      //     })
      //   }
      // });
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
