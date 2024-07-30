import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { TradeProposal } from 'src/app/shared/model/tradeProposal.model';
import { TradeSubject } from 'src/app/shared/model/tradeSubject.model';
import { TradesService } from '../trades.service';

@Component({
  selector: 'app-show-request-details-prompt',
  templateUrl: './show-request-details-prompt.component.html',
  styleUrls: ['./show-request-details-prompt.component.css'],
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
export class ShowRequestDetailsPromptComponent implements OnInit{
  buttonState: string = 'idle';
  focused: string = '';
  user: User | undefined;
  allTradeSubjects: TradeSubject[] = [];
  ownTradeSubjects: TradeSubject[] = [];
  partnerTradeSubjects: TradeSubject[] = [];
  tradeProposal: TradeProposal | undefined;

  constructor(private snackBar: MatSnackBar,
              private dialogRef: MatDialogRef<ShowRequestDetailsPromptComponent>,
              private authService: AuthService,
              private tradesService: TradesService,
              @Inject(MAT_DIALOG_DATA) public data: any,) {
    this.tradeProposal = data;
    this.authService.user$.subscribe((user) => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    this.tradesService.getTradeProposalDetailsByID(this.tradeProposal?.idZahTrg!).subscribe({
      next: (result: TradeSubject[] | TradeSubject) => {
        if(Array.isArray(result)){
          this.allTradeSubjects = result;

          // Filter the trade subjects into correct arrays
          if(this.tradeProposal?.statusZahTrg == 1){  // ACCEPTED
            this.allTradeSubjects.forEach(tradeSubject => {
              if(tradeSubject.tipPredTrg == 0 && tradeSubject.idTim != this.user?.teamId){ // Player type 
                this.ownTradeSubjects.push(tradeSubject);
              } else if(tradeSubject.tipPredTrg == 0 && tradeSubject.idTim == this.user?.teamId){ // Player type 
                this.partnerTradeSubjects.push(tradeSubject);
              } else if(tradeSubject.tipPredTrg == 1 && tradeSubject.idTim != this.user?.teamId){ // Pick type 
                this.ownTradeSubjects.push(tradeSubject);
              }else if(tradeSubject.tipPredTrg == 1 && tradeSubject.idTim == this.user?.teamId){ // Pick type 
                this.partnerTradeSubjects.push(tradeSubject);
              }else if(tradeSubject.tipPredTrg == 2 && tradeSubject.idTim != this.user?.teamId){ // Draft Rights type 
                this.ownTradeSubjects.push(tradeSubject);
              }else if(tradeSubject.tipPredTrg == 2 && tradeSubject.idTim == this.user?.teamId){ // Draft Rights type 
                this.partnerTradeSubjects.push(tradeSubject);
              }
            });
          } else {
            this.allTradeSubjects.forEach(tradeSubject => {
              if(tradeSubject.tipPredTrg == 0 && tradeSubject.idTim == this.user?.teamId){ // Player type 
                this.ownTradeSubjects.push(tradeSubject);
              } else if(tradeSubject.tipPredTrg == 0 && tradeSubject.idTim != this.user?.teamId){ // Player type 
                this.partnerTradeSubjects.push(tradeSubject);
              } else if(tradeSubject.tipPredTrg == 1 && tradeSubject.idTim == this.user?.teamId){ // Pick type 
                this.ownTradeSubjects.push(tradeSubject);
              }else if(tradeSubject.tipPredTrg == 1 && tradeSubject.idTim != this.user?.teamId){ // Pick type 
                this.partnerTradeSubjects.push(tradeSubject);
              }else if(tradeSubject.tipPredTrg == 2 && tradeSubject.idTim == this.user?.teamId){ // Draft Rights type 
                this.ownTradeSubjects.push(tradeSubject);
              }else if(tradeSubject.tipPredTrg == 2 && tradeSubject.idTim != this.user?.teamId){ // Draft Rights type 
                this.partnerTradeSubjects.push(tradeSubject);
              }
            });
          }
        }
      }
    })
  }

  closeButtonClicked() {
    this.buttonState = 'clicked';
    setTimeout(() => { this.buttonState = 'idle'; }, 200);
    this.dialogRef.close();
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }

  overviewClicked(){
    this.dialogRef.close();
  }

  faPlus = faPlus;
  faMinus = faMinus;
}
