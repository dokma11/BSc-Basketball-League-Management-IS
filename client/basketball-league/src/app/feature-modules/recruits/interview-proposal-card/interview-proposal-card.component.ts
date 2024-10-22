import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { faCheck, faTimes } from '@fortawesome/free-solid-svg-icons';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Employee } from 'src/app/shared/model/employee.model';
import { Team } from 'src/app/shared/model/team.model';
import { TradesService } from '../../trades/trades.service';
import { InterviewProposal } from 'src/app/shared/model/interviewProposal.model';
import { RecruitsService } from '../recruits.service';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { AcceptProposalPromptComponent } from '../accept-proposal-prompt/accept-proposal-prompt.component';
import { DeclineProposalPromptComponent } from '../decline-proposal-prompt/decline-proposal-prompt.component';

@Component({
  selector: 'app-interview-proposal-card',
  templateUrl: './interview-proposal-card.component.html',
  styleUrls: ['./interview-proposal-card.component.css'],
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
export class InterviewProposalCardComponent implements OnInit{
  acceptButtonState: string = 'idle';
  declineButtonState: string = 'idle';
  @Input() interviewProposal!: InterviewProposal;
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  proposalOccurrenceTime: string = '';
  proposalOccurrenceDate: string = '';
  coach: Employee | undefined;
  coachsTeam: Team | undefined;
  dateDay: string = '';
  dateMonth: string = '';
  dateYear: string = '';
  time: string = '';
  timeAddition: string = '';

  constructor(private dialog: MatDialog, 
              private tradesService: TradesService, 
              private recruitsService: RecruitsService) {
    
  }

  ngOnInit(): void {
    const proposalDateTimeString = this.interviewProposal.datVrePozInt.toString();
    [this.proposalOccurrenceDate, this.proposalOccurrenceTime] = proposalDateTimeString.split('T');
    [this.time, this.timeAddition] = this.proposalOccurrenceTime.split('+');
    [this.dateYear, this.dateMonth, this.dateDay] = this.proposalOccurrenceDate.split('-');
    
    const [hours, minutes] = this.time.split(':'); 
    this.time = `${parseInt(hours, 10) - 1}:${minutes}`; 
    
    this.proposalOccurrenceDate = this.dateDay + '.' + this.dateMonth + '.' + this.dateYear + '.'
    this.proposalOccurrenceTime = this.time + 'h';

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
    
    this.dialogRef = this.dialog.open(AcceptProposalPromptComponent, {
      data: {
        interviewProposal: this.interviewProposal
      } 
    });
    
    this.dialogRef.afterClosed().subscribe((result: any) => {
      this.dialogRefClosed.emit(result);
    });
  }

  declineButtonClicked() {
    this.declineButtonState = 'clicked';
    setTimeout(() => { this.declineButtonState = 'idle'; }, 200);
    
    this.dialogRef = this.dialog.open(DeclineProposalPromptComponent, {
      data: {
        interviewProposal: this.interviewProposal
      } 
    });
    
    this.dialogRef.afterClosed().subscribe((result: any) => {
      this.dialogRefClosed.emit(result);
    });
  }

  faCheck = faCheck;
  faTimes = faTimes;
}
