import { trigger, state, style, transition, animate } from '@angular/animations';
import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { TradeProposal } from 'src/app/shared/model/tradeProposal.model';
import { TradesService } from '../trades.service';
import { Employee } from 'src/app/shared/model/employee.model';

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
  tradeProposal: TradeProposal | undefined;
  manager: Employee | undefined;
  denialReason: string = '';

  constructor(@Inject(MAT_DIALOG_DATA) public data: any,
              private dialogRef: MatDialogRef<SeeDenialExplanationPromptComponent>,
              private tradesService: TradesService) {
    this.tradeProposal = data;

    this.tradesService.getManagerByID(this.tradeProposal?.idMenadzerPrim!).subscribe({
      next: (result: Employee) => {
        this.manager = result;
        this.denialReason = this.manager?.ime + ' ' + this.manager?.prezime + ' wrote: ' + this.tradeProposal?.razlogOdbij;
      }
    })
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
