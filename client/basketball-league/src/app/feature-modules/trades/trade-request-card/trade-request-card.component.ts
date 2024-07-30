import { trigger, state, style, transition, animate } from '@angular/animations';
import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { faCheck, faTimes, faPen, faTrash, faCircleInfo, faBan } from '@fortawesome/free-solid-svg-icons';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { DeclineRequestPromptComponent } from '../decline-request-prompt/decline-request-prompt.component';
import { AcceptRequestPromptComponent } from '../accept-request-prompt/accept-request-prompt.component';
import { SeeDenialExplanationPromptComponent } from '../see-denial-explanation-prompt/see-denial-explanation-prompt.component';
import { ShowRequestDetailsPromptComponent } from '../show-request-details-prompt/show-request-details-prompt.component';
import { TradeProposal } from 'src/app/shared/model/tradeProposal.model';
import { Employee } from 'src/app/shared/model/employee.model';
import { Team } from 'src/app/shared/model/team.model';
import { TradesService } from '../trades.service';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { CancelRequestPromptComponent } from '../cancel-request-prompt/cancel-request-prompt.component';

@Component({
  selector: 'app-trade-request-card',
  templateUrl: './trade-request-card.component.html',
  styleUrls: ['./trade-request-card.component.css'],
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
export class TradeRequestCardComponent implements OnInit{
  acceptButtonState: string = 'idle';
  declineButtonState: string = 'idle';
  detailsButtonState: string = 'idle';
  seeExplanationButtonState: string = 'idle';
  cancelButtonState: string = 'idle';
  @Input() tradeProposal!: TradeProposal;
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  proposalOccurrenceTime: string = "";
  proposalOccurrenceDate: string = "";
  manager: Employee | undefined;
  managersTeam: Team | undefined;
  dateDay: string = "";
  dateMonth: string = "";
  dateYear: string = "";

  constructor(private dialog: MatDialog, 
              private tradesService: TradesService, 
              private authService: AuthService) {
    this.authService.user$.subscribe((user) => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    const proposalDateTimeString = this.tradeProposal.datZahTrg.toString();
    [this.proposalOccurrenceDate, this.proposalOccurrenceTime] = proposalDateTimeString.split('T');
    [this.dateYear, this.dateMonth, this.dateDay] = this.proposalOccurrenceDate.split('-');
    this.proposalOccurrenceDate = this.dateDay + '.' + this.dateMonth + '.' + this.dateYear + '.'

    this.tradesService.getTeamByManagerID(this.tradeProposal.idMenadzerPos).subscribe({
      next: (result: Team) => {
        this.managersTeam = result;
      }
    });
    
    this.tradesService.getManagerByID(this.tradeProposal.idMenadzerPos).subscribe({
      next: (result: Employee) => {
        this.manager = result;
      }
    });
  }

  acceptButtonClicked() {
    this.acceptButtonState = 'clicked';
    setTimeout(() => { this.acceptButtonState = 'idle'; }, 200);
    
    this.dialogRef = this.dialog.open(AcceptRequestPromptComponent, {
      data: this.tradeProposal
    });
    
    this.dialogRef.afterClosed().subscribe((result: any) => {
      this.dialogRefClosed.emit(result);
    });
  }

  declineButtonClicked() {
    this.declineButtonState = 'clicked';
    setTimeout(() => { this.declineButtonState = 'idle'; }, 200);
    
    this.dialogRef = this.dialog.open(DeclineRequestPromptComponent, {
      data: this.tradeProposal
    });
    
    this.dialogRef.afterClosed().subscribe((result: any) => {
      this.dialogRefClosed.emit(result);
    });
  }

  seeDetailsButtonClicked(){
    this.detailsButtonState = 'clicked';
    setTimeout(() => { this.detailsButtonState = 'idle'; }, 200);
   
    this.dialogRef = this.dialog.open(ShowRequestDetailsPromptComponent, {
      data: this.tradeProposal
    });
  }

  seeExplanationButtonClicked() {
    this.seeExplanationButtonState = 'clicked';
    setTimeout(() => { this.seeExplanationButtonState = 'idle'; }, 200);
    
    this.dialogRef = this.dialog.open(SeeDenialExplanationPromptComponent, {
      data: this.tradeProposal
    });
  }

  cancelButtonClicked(){
    this.cancelButtonState = 'clicked';
    setTimeout(() => { this.cancelButtonState = 'idle'; }, 200);
    
    this.dialogRef = this.dialog.open(CancelRequestPromptComponent, {
      data: this.tradeProposal
    });
    
    this.dialogRef.afterClosed().subscribe((result: any) => {
      this.dialogRefClosed.emit(result);
    });
  }

  faCheck = faCheck;
  faTimes = faTimes;
  faPen = faPen;
  faTrash = faTrash;
  faCircleInfo = faCircleInfo;
  faBan = faBan;
}
  