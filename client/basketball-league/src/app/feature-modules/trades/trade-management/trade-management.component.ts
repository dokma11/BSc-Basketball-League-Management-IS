import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { ProposeTradeFormComponent } from '../propose-trade-form/propose-trade-form.component';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { TradeProposal } from 'src/app/shared/model/tradeProposal.model';
import { TradesService } from '../trades.service';

@Component({
  selector: 'app-trade-management',
  templateUrl: './trade-management.component.html',
  styleUrls: ['./trade-management.component.css'],
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
export class TradeManagementComponent implements OnInit{
  user: User | undefined;
  backgroundSize: string = '100% 100%';
  tradeProposals: TradeProposal[] = [];
  proposeTradeButtonState: string = "";

  private dialogRef: any;

  tradeForm = new FormGroup({
    selectedTradeType: new FormControl('Received', [Validators.required]),
  });
  
  tradeStatusForm = new FormGroup({
    selectedStatus: new FormControl('None', [Validators.required]),
  });

  constructor(private authService: AuthService,
              private dialog: MatDialog,
              private snackBar: MatSnackBar,
              private tradesService: TradesService) {
    this.authService.user$.subscribe((user) => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    this.getProposals();
  }

  getProposals() {
    if(this.tradeForm.value.selectedTradeType == 'Received'){
      this.tradesService.getAllReceivedTradeProposalsByManagerID(this.user?.id!).subscribe({
        next: (result: TradeProposal[] | TradeProposal) => {
          if(Array.isArray(result)){
            this.tradeProposals = [];
            this.tradeProposals = result;
            
            if(this.tradeStatusForm.value.selectedStatus == 'In progress'){
              this.tradeProposals = this.tradeProposals.filter(tradeProposal => tradeProposal.statusZahTrg === 0);
            } else if(this.tradeStatusForm.value.selectedStatus == 'Accepted'){
              this.tradeProposals = this.tradeProposals.filter(tradeProposal => tradeProposal.statusZahTrg === 1);
            } else if(this.tradeStatusForm.value.selectedStatus == 'Declined'){
              this.tradeProposals = this.tradeProposals.filter(tradeProposal => tradeProposal.statusZahTrg === 2);
            } else if(this.tradeStatusForm.value.selectedStatus == 'Cancelled'){
              this.tradeProposals = this.tradeProposals.filter(tradeProposal => tradeProposal.statusZahTrg === 3);
            } 
          }
        }
      });
    } else if (this.tradeForm.value.selectedTradeType == 'Sent'){
      this.tradesService.getAllSentTradeProposalsByManagerID(this.user?.id!).subscribe({
        next: (result: TradeProposal[] | TradeProposal) => {
          if(Array.isArray(result)){
            this.tradeProposals = [];
            this.tradeProposals = result;
            
            if(this.tradeStatusForm.value.selectedStatus == 'In progress'){
              this.tradeProposals = this.tradeProposals.filter(tradeProposal => tradeProposal.statusZahTrg === 0);
            } else if(this.tradeStatusForm.value.selectedStatus == 'Accepted'){
              this.tradeProposals = this.tradeProposals.filter(tradeProposal => tradeProposal.statusZahTrg === 1);
            } else if(this.tradeStatusForm.value.selectedStatus == 'Declined'){
              this.tradeProposals = this.tradeProposals.filter(tradeProposal => tradeProposal.statusZahTrg === 2);
            } else if(this.tradeStatusForm.value.selectedStatus == 'Cancelled'){
              this.tradeProposals = this.tradeProposals.filter(tradeProposal => tradeProposal.statusZahTrg === 3);
            } 
          }
        }
      });
    }
  }

  handleDialogClosed(result: any) {
    this.getProposals();
  }

  onTradeTypeChange(event: any) {
    this.showNotification('Selected trades type: ' + this.tradeForm.value.selectedTradeType);
    this.getProposals();
  }

  onTradeStatusChange(event: any) {
    this.showNotification('Selected trade status: ' + this.tradeStatusForm.value.selectedStatus);
    this.getProposals();
  }

  proposeTradeButtonClicked() {
    this.proposeTradeButtonState = 'clicked';
    setTimeout(() => { this.proposeTradeButtonState = 'idle'; }, 200);
    this.dialogRef = this.dialog.open(ProposeTradeFormComponent, {
    });

    if (this.dialogRef) {
      this.dialogRef.afterClosed().subscribe((result: any) => {
        this.getProposals();
      });
    }
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }

  faPlus = faPlus;
}
