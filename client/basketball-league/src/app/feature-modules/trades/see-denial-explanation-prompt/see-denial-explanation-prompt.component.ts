import { trigger, state, style, transition, animate } from '@angular/animations';
import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-see-denial-explanation-prompt',
  templateUrl: './see-denial-explanation-prompt.component.html',
  styleUrls: ['./see-denial-explanation-prompt.component.css'],
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
export class SeeDenialExplanationPromptComponent {
  closeButtonState: string = '';
  focused: string = '';
  // request: PersonalTourRequest | undefined;  OVDE IDE TRADE REQUEST KADA POVEZEM SA BEKOM
  // organizer: Organizer | undefined;
  denialReason: string = '';

  constructor(@Inject(MAT_DIALOG_DATA) public data: any,
              private dialogRef: MatDialogRef<SeeDenialExplanationPromptComponent>,) {
    //this.request = data;

    // this.toursService.getOrganizerById(this.request?.organizerId!).subscribe({
    //   next: (result: Organizer) => {
    //     this.organizer = result;
    //     this.denialReason = this.organizer?.firstName + ' ' + this.organizer?.lastName + ' wrote: ' + this.request?.denialReason;
    //   }
    // })
  }

  closeButtonClicked() {
    this.closeButtonState = 'clicked';
    setTimeout(() => { this.closeButtonState = 'idle'; }, 200);
    this.dialogRef.close();
  }

  overviewClicked(){
    this.dialogRef.close();
  }
}
