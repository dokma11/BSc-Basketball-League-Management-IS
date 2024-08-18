import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, OnInit } from '@angular/core';
import { RecruitsService } from '../recruits.service';
import { Recruit } from 'src/app/shared/model/recruit.model';
import { MatDialog } from '@angular/material/dialog';
import { InterviewInvitePromptComponent } from '../interview-invite-prompt/interview-invite-prompt.component';
import { TrainingInvitePromptComponent } from '../training-invite-prompt/training-invite-prompt.component';

@Component({
  selector: 'app-recruit-management',
  templateUrl: './recruit-management.component.html',
  styleUrls: ['./recruit-management.component.css'],
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
export class RecruitManagementComponent implements OnInit{
  recruits: Recruit[] = [];
  private dialogRef: any;

  constructor(private recruitsService: RecruitsService,
              private dialog: MatDialog,
  ) { }

  ngOnInit(): void {
    this.recruitsService.getAllRecruits().subscribe({
      next: (result: Recruit[] | Recruit) => {
        if (Array.isArray(result)) {
          this.recruits = result;
        }
      }
    });

    
  }

  probazaformuintervju() {
    this.dialogRef = this.dialog.open(InterviewInvitePromptComponent, {
      data: {
        recruitId: 0,
        coachId: 0
      }
    });
  }

  probazaformutrening() {
    this.dialogRef = this.dialog.open(TrainingInvitePromptComponent, {
      data: {
        recruitId: 0,
        coachId: 0
      }
    });
  }

}
