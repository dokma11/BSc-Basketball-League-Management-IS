import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MaterialModule } from 'src/app/infrastructure/material/material-module';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { RecruitManagementComponent } from './recruit-management/recruit-management.component';
import { RecruitCardComponent } from './recruit-card/recruit-card.component';
import { InterviewInvitePromptComponent } from './interview-invite-prompt/interview-invite-prompt.component';
import { TrainingInvitePromptComponent } from './training-invite-prompt/training-invite-prompt.component';
import { ProposalManagementComponent } from './proposal-management/proposal-management.component';
import { TrainingProposalCardComponent } from './training-proposal-card/training-proposal-card.component';
import { InterviewProposalCardComponent } from './interview-proposal-card/interview-proposal-card.component';
import { AcceptProposalPromptComponent } from './accept-proposal-prompt/accept-proposal-prompt.component';
import { DeclineProposalPromptComponent } from './decline-proposal-prompt/decline-proposal-prompt.component';

@NgModule({
  declarations: [
    RecruitManagementComponent,
    RecruitCardComponent,
    InterviewInvitePromptComponent,
    TrainingInvitePromptComponent,
    ProposalManagementComponent,
    TrainingProposalCardComponent,
    InterviewProposalCardComponent,
    AcceptProposalPromptComponent,
    DeclineProposalPromptComponent,
  ],
  imports: [
    CommonModule,
    RouterModule,
    ReactiveFormsModule,
    MaterialModule,
    FontAwesomeModule,
  ],
  exports: [
    RecruitManagementComponent,
  ]
})
export class RecruitsModule { }
