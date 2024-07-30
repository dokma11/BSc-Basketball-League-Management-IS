import { trigger, state, style, transition, animate } from '@angular/animations';
import { Component, Inject } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { TradeProposal, TradeProposalStatus } from 'src/app/shared/model/tradeProposal.model';
import { TradesService } from '../trades.service';

@Component({
  selector: 'app-decline-request-prompt',
  templateUrl: './decline-request-prompt.component.html',
  styleUrls: ['./decline-request-prompt.component.css'],
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
export class DeclineRequestPromptComponent {
  cancelButtonState: string = 'idle';
  declineButtonState: string = 'idle';
  tradeProposal: TradeProposal | undefined;
  focused: string = '';

  constructor(private snackBar: MatSnackBar,
              private dialogRef: MatDialogRef<DeclineRequestPromptComponent>,
              private tradesService: TradesService,
              @Inject(MAT_DIALOG_DATA) public data: any) {
    this.tradeProposal = data;
  }

  declineRequestForm = new FormGroup({
    explanation: new FormControl('', [Validators.required]),
  });

  declineButtonClicked(){
    this.declineButtonState = 'clicked';
    setTimeout(() => { this.declineButtonState = 'idle'; }, 200);

    this.tradeProposal!.statusZahTrg = TradeProposalStatus.DECLINED;
    this.tradeProposal!.razlogOdbij = this.declineRequestForm.value.explanation || undefined;

    this.tradesService.updateTradeProposal(this.tradeProposal!).subscribe({
      next: () => {
        this.showNotification('Trade proposal successfully declined!');
        this.dialogRef.close();
      }
    });
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
