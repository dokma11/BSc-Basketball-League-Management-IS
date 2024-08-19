import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { InterviewProposal } from 'src/app/shared/model/interviewProposal.model';
import { TrainingProposal } from 'src/app/shared/model/trainingProposal.model';
import { RecruitsService } from '../recruits.service';

@Component({
  selector: 'app-proposal-management',
  templateUrl: './proposal-management.component.html',
  styleUrls: ['./proposal-management.component.css'],
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
  ]
})
export class ProposalManagementComponent implements OnInit{
  user: User | undefined;
  backgroundSize: string = '100% 100%';
  trainingProposals: TrainingProposal[] = [];
  interviewProposals: InterviewProposal[] = [];
  proposeTradeButtonState: string = "";

  proposalForm = new FormGroup({
    selectedProposalType: new FormControl('Interview', [Validators.required]),
  });
  
  proposalStatusForm = new FormGroup({
    selectedStatus: new FormControl('None', [Validators.required]),
  });

  constructor(private authService: AuthService,
              private snackBar: MatSnackBar,
              private recruitsService: RecruitsService) {
    this.authService.user$.subscribe((user) => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    this.getProposals();
  }

  getProposals() {
    if(this.proposalForm.value.selectedProposalType == 'Interview'){
      this.recruitsService.getInterviewProposalsByRecruitId(this.user?.id!).subscribe({
        next: (result: InterviewProposal[] | InterviewProposal) => {
          if(Array.isArray(result)){
            this.trainingProposals = [];
            this.interviewProposals = [];
            this.interviewProposals = result;
            
            if(this.proposalStatusForm.value.selectedStatus == 'In progress'){
              this.interviewProposals = this.interviewProposals.filter(interviewProposal => interviewProposal.statusPozInt === 0);
            } else if(this.proposalStatusForm.value.selectedStatus == 'Accepted'){
              this.interviewProposals = this.interviewProposals.filter(interviewProposal => interviewProposal.statusPozInt === 1);
            } else if(this.proposalStatusForm.value.selectedStatus == 'Declined'){
              this.interviewProposals = this.interviewProposals.filter(interviewProposal => interviewProposal.statusPozInt === 2);
            } 
          }
        }
      });
    } else if (this.proposalForm.value.selectedProposalType == 'Training'){
      this.recruitsService.getTrainingProposalsByRecruitId(this.user?.id!).subscribe({
        next: (result: TrainingProposal[] | TrainingProposal) => {
          if(Array.isArray(result)){
            this.interviewProposals = [];
            this.trainingProposals = [];
            this.trainingProposals = result;
            
            if(this.proposalStatusForm.value.selectedStatus == 'In progress'){
              this.trainingProposals = this.trainingProposals.filter(trainingProposal => trainingProposal.statusPozTrng === 0);
            } else if(this.proposalStatusForm.value.selectedStatus == 'Accepted'){
              this.trainingProposals = this.trainingProposals.filter(trainingProposal => trainingProposal.statusPozTrng === 1);
            } else if(this.proposalStatusForm.value.selectedStatus == 'Declined'){
              this.trainingProposals = this.trainingProposals.filter(trainingProposal => trainingProposal.statusPozTrng === 2);
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
    this.showNotification('Selected proposal type: ' + this.proposalForm.value.selectedProposalType);
    this.getProposals();
  }

  onTradeStatusChange(event: any) {
    this.showNotification('Selected proposal status: ' + this.proposalStatusForm.value.selectedStatus);
    this.getProposals();
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }
}
