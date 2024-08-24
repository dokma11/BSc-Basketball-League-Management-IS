import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, OnInit } from '@angular/core';
import { RecruitsService } from '../recruits.service';
import { Recruit } from 'src/app/shared/model/recruit.model';
import { FormGroup, FormControl, Validators } from '@angular/forms';

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
  recruitsBeforeSort: Recruit[] = [];
  focused: string = '';
  buttonState: string = 'idle';
  resetButtonState: string = 'idle';

  constructor(private recruitsService: RecruitsService) { }
           
  ngOnInit(): void {
    this.getRecruits();
  }

  positionForm = new FormGroup({
    selectedPosition: new FormControl('None', [Validators.required]),
  });

  searchForm = new FormGroup({
    parameters: new FormControl('', [Validators.required]),
  });

  getRecruits() {
    this.recruitsService.getAllRecruits().subscribe({
      next: (result: Recruit[] | Recruit) => {
        if (Array.isArray(result)) {
          this.recruits = result;
        }
      }
    });
  }

  sortRecruits(position: any) {
    if (this.searchForm.value.parameters == '') {
      this.recruitsService.getAllRecruits().subscribe({
        next: (result: Recruit[] | Recruit) => {
          if (Array.isArray(result)) {
            this.recruitsBeforeSort = result;
            this.recruits = [];
            this.recruitsBeforeSort.forEach((recruit, index) => {
              if (recruit.pozReg.toString() == position){
                this.recruits.push(recruit);
              }
            });
          }
        }
      });
    } else {
      this.recruitsService.getAllRecruitsByName(this.searchForm.value.parameters!).subscribe({
        next: (result: Recruit[] | Recruit) => {
          if (Array.isArray(result)){
            this.recruitsBeforeSort = result;
            this.recruits = [];
            this.recruitsBeforeSort.forEach((recruit, index) => {
              if (recruit.pozReg.toString() == position){
                this.recruits.push(recruit);
              }
            });
          }
        }
      });
    }
  }

  searchButtonClicked() {
    this.buttonState = 'clicked';
    setTimeout(() => { this.buttonState = 'idle'; }, 200);

    if (this.positionForm.value.selectedPosition == 'None'){
      this.recruitsService.getAllRecruitsByName(this.searchForm.value.parameters!).subscribe({
        next: (result: Recruit[] | Recruit) => {
          if (Array.isArray(result)){
            this.recruits = [];
            this.recruits = result;
          }
        }
      });
    } else {
      this.recruitsService.getAllRecruitsByName(this.searchForm.value.parameters!).subscribe({
        next: (result: Recruit[] | Recruit) => {
          if (Array.isArray(result)){
            this.recruitsBeforeSort = result;
            this.recruits = [];
            this.recruitsBeforeSort.forEach((recruit, index) => {
              if (recruit.pozReg.toString() == this.positionForm.value.selectedPosition){
                this.recruits.push(recruit);
              }
            });
          }
        }
      });
    }
  }

  resetButtonClicked() {
    this.resetButtonState = 'clicked';
    setTimeout(() => { this.resetButtonState = 'idle'; }, 200);

    this.searchForm.get('parameters')?.setValue('');
    this.positionForm.get('selectedPosition')?.setValue('None');
    this.getRecruits();
  }

  onPositionChange(event: any) {
    if (this.positionForm.value.selectedPosition == 'None' && this.searchForm.value.parameters == ''){
      this.getRecruits();
    } else if (this.positionForm.value.selectedPosition == 'None' && this.searchForm.value.parameters != '') {
      this.searchButtonClicked();
    } else {
      this.sortRecruits(this.positionForm.value.selectedPosition);
    }
  }

  handleDialogClosed(result: any) {
    this.getRecruits();
  }
}
