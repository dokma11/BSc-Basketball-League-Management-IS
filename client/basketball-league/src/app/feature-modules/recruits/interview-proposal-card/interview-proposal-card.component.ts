import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { faCheck, faTimes } from '@fortawesome/free-solid-svg-icons';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Employee } from 'src/app/shared/model/employee.model';
import { Team } from 'src/app/shared/model/team.model';
import { AcceptRequestPromptComponent } from '../../trades/accept-request-prompt/accept-request-prompt.component';
import { DeclineRequestPromptComponent } from '../../trades/decline-request-prompt/decline-request-prompt.component';
import { TradesService } from '../../trades/trades.service';
import { InterviewProposal } from 'src/app/shared/model/interviewProposal.model';
import { RecruitsService } from '../recruits.service';

@Component({
  selector: 'app-interview-proposal-card',
  templateUrl: './interview-proposal-card.component.html',
  styleUrls: ['./interview-proposal-card.component.css']
})
export class InterviewProposalCardComponent implements OnInit{
  acceptButtonState: string = 'idle';
  declineButtonState: string = 'idle';
  @Input() interviewProposal!: InterviewProposal;
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  proposalOccurrenceTime: string = "";
  proposalOccurrenceDate: string = "";
  coach: Employee | undefined;
  coachsTeam: Team | undefined;
  dateDay: string = "";
  dateMonth: string = "";
  dateYear: string = "";

  constructor(private dialog: MatDialog, 
              private tradesService: TradesService, 
              private recruitsService: RecruitsService) {
    
  }

  ngOnInit(): void {
    const proposalDateTimeString = this.interviewProposal.datVrePozInt.toString();
    [this.proposalOccurrenceDate, this.proposalOccurrenceTime] = proposalDateTimeString.split('T');
    [this.dateYear, this.dateMonth, this.dateDay] = this.proposalOccurrenceDate.split('-');
    this.proposalOccurrenceDate = this.dateDay + '.' + this.dateMonth + '.' + this.dateYear + '.'

    this.recruitsService.getTeamByCoachID(this.interviewProposal.idTrener).subscribe({
      next: (result: Team) => {
        this.coachsTeam = result;
      }
    });
    
    this.tradesService.getManagerByID(this.interviewProposal.idTrener).subscribe({
      next: (result: Employee) => {
        this.coach = result;
      }
    });
  }

  acceptButtonClicked() {
    this.acceptButtonState = 'clicked';
    setTimeout(() => { this.acceptButtonState = 'idle'; }, 200);
    
    this.dialogRef = this.dialog.open(AcceptRequestPromptComponent, {
      data: this.interviewProposal
    });
    
    this.dialogRef.afterClosed().subscribe((result: any) => {
      this.dialogRefClosed.emit(result);
    });
  }

  declineButtonClicked() {
    this.declineButtonState = 'clicked';
    setTimeout(() => { this.declineButtonState = 'idle'; }, 200);
    
    this.dialogRef = this.dialog.open(DeclineRequestPromptComponent, {
      data: this.interviewProposal
    });
    
    this.dialogRef.afterClosed().subscribe((result: any) => {
      this.dialogRefClosed.emit(result);
    });
  }

  faCheck = faCheck;
  faTimes = faTimes;
}
