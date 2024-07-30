import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { TradeProposal, TradeProposalStatus } from 'src/app/shared/model/tradeProposal.model';
import { DeclineRequestPromptComponent } from '../decline-request-prompt/decline-request-prompt.component';
import { TradesService } from '../trades.service';

@Component({
  selector: 'app-cancel-request-prompt',
  templateUrl: './cancel-request-prompt.component.html',
  styleUrls: ['./cancel-request-prompt.component.css']
})
export class CancelRequestPromptComponent {
  noButtonState: string = 'idle';
  yesButtonState: string = 'idle';
  focused: string = '';
  tradeProposal: TradeProposal | undefined;

  constructor(private snackBar: MatSnackBar,
              private dialogRef: MatDialogRef<DeclineRequestPromptComponent>,
              private tradesService: TradesService,
              @Inject(MAT_DIALOG_DATA) public data: any,) {
    this.tradeProposal = data;
  }

  yesButtonClicked(){
    this.yesButtonState = 'clicked';
    setTimeout(() => { this.yesButtonState = 'idle'; }, 200);

    this.tradeProposal!.statusZahTrg = TradeProposalStatus.CANCELLED;

    this.tradesService.updateTradeProposal(this.tradeProposal!).subscribe({
      next: () => {
        this.showNotification('Trade proposal successfully cancelled!');
        this.dialogRef.close();
      }
    });
  }

  noButtonClicked(){
    this.noButtonState = 'clicked';
    setTimeout(() => { this.noButtonState = 'idle'; }, 200);
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
