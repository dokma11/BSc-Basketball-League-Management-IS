import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { InterviewProposal, InterviewProposalStatus } from 'src/app/shared/model/interviewProposal.model';
import { TrainingProposal, TrainingProposalStatus } from 'src/app/shared/model/trainingProposal.model';
import { RecruitsService } from '../recruits.service';

@Component({
  selector: 'app-accept-proposal-prompt',
  templateUrl: './accept-proposal-prompt.component.html',
  styleUrls: ['./accept-proposal-prompt.component.css']
})
export class AcceptProposalPromptComponent {
  cancelButtonState: string = 'idle';
  acceptButtonState: string = 'idle';
  focused: string = '';
  interviewProposal: InterviewProposal | undefined;
  trainingProposal: TrainingProposal | undefined;
  interviewChosen:  boolean = true;

  constructor(private snackBar: MatSnackBar,
              private dialogRef: MatDialogRef<AcceptProposalPromptComponent>,
              private recruitsService: RecruitsService,
              @Inject(MAT_DIALOG_DATA) public data: any,) {
    if (data.interview) {
      this.interviewProposal = data.interviewProposal;
    } else if (data.training) {
      this.trainingProposal = data.trainingProposal;
      this.interviewChosen = false;
    }
  }

  acceptButtonClicked(){
    this.acceptButtonState = 'clicked';
    setTimeout(() => { this.acceptButtonState = 'idle'; }, 200);

    if(this.interviewChosen) {
      this.interviewProposal!.statusPozInt = InterviewProposalStatus.AFFIRMED;
      this.recruitsService.updateInterviewProposal(this.interviewProposal!).subscribe({
        next: () => {
          this.showNotification('Interview proposal successfully accepted!');
          this.dialogRef.close();
        }
      });
    } else {
      this.trainingProposal!.statusPozTrng = TrainingProposalStatus.APPROVED;
      this.recruitsService.updateTrainingProposal(this.trainingProposal!).subscribe({
        next: () => {
          this.showNotification('Training proposal successfully accepted!');
          this.dialogRef.close();
        }
      });
    }
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
