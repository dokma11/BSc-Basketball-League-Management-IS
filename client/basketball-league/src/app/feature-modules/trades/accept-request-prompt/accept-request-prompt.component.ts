import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { DeclineRequestPromptComponent } from '../decline-request-prompt/decline-request-prompt.component';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { TradeProposal, TradeProposalStatus } from 'src/app/shared/model/tradeProposal.model';
import { TradesService } from '../trades.service';

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
  tradeProposal: TradeProposal | undefined;

  constructor(private snackBar: MatSnackBar,
              private dialogRef: MatDialogRef<DeclineRequestPromptComponent>,
              private tradesService: TradesService,
              @Inject(MAT_DIALOG_DATA) public data: any,) {
    this.tradeProposal = data;
  }

  acceptButtonClicked(){
    this.acceptButtonState = 'clicked';
    setTimeout(() => { this.acceptButtonState = 'idle'; }, 200);

    this.tradeProposal!.statusZahTrg = TradeProposalStatus.ACCEPTED;

    this.tradesService.updateTradeProposal(this.tradeProposal!).subscribe({
      next: () => {
        this.showNotification('Trade proposal successfully accepted!');
        this.dialogRef.close();

        this.tradesService.commitTrade(this.tradeProposal!).subscribe({
          next: (result: TradeProposal) => {}
        })
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
