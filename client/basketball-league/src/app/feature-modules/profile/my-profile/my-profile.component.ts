import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { ProfileService } from '../profile.service';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { MatSnackBar } from '@angular/material/snack-bar';
import { FormGroup, FormControl, Validators } from '@angular/forms';

@Component({
  selector: 'app-my-profile',
  templateUrl: './my-profile.component.html',
  styleUrls: ['./my-profile.component.css'],
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
export class MyProfileComponent implements OnInit{
  user: User | undefined;
  editButtonState: string = 'idle';
  doneButtonState: string = 'idle';
  cancelButtonState: string = 'idle';
  canEdit: boolean = false;
  birthDayDate: string = '';
  birthDayTime: string = '';
  dateYear: string = '';
  dateMonth: string = '';
  dateDay: string = '';
  
  constructor(private authService: AuthService, 
              private profileService: ProfileService,
              private snackBar: MatSnackBar,) {
    this.authService.user$.subscribe((user) => {
      this.user = user;
    });
  }

  editProfileForm = new FormGroup({
    firstName: new FormControl('', [Validators.required]),
    lastName: new FormControl('', [Validators.required]),
    email: new FormControl('', [Validators.required]),
    dateOfBirth: new FormControl(null, [Validators.required]),
  });

  ngOnInit(): void {
    this.profileService.getUserByID(this.user?.id!).subscribe({
      next: (result: User) => {
        this.user = result;

        [this.birthDayDate, this.birthDayTime] = this.user?.datRodj!.toString().split('T');
        [this.dateYear, this.dateMonth, this.dateDay] = this.birthDayDate.split('-');
        this.birthDayDate = this.dateDay + '.' + this.dateMonth + '.' + this.dateYear + '.'
      }
    });
  }

  editButtonClicked() {
    this.editButtonState = 'clicked';
    setTimeout(() => { this.editButtonState = 'idle'; }, 200);

    this.canEdit = true;
    this.showNotification('You can now edit your profile information'); 
  }

  doneButtonClicked() {
    this.doneButtonState = 'clicked';
    setTimeout(() => { this.doneButtonState = 'idle'; }, 200);
    
    const editedUser: User = {
      id: this.user?.id || 0,
      ime: this.editProfileForm.value.firstName || this.user?.ime,
      prezime: this.editProfileForm.value.lastName || this.user?.prezime,
      email: this.editProfileForm.value.email || this.user?.email || "",
      datRodj: this.editProfileForm.value.dateOfBirth || this.user?.datRodj,
    };

    this.profileService.updateUserProfile(editedUser).subscribe({
      next: (result: User) => {
        this.canEdit = false;
        this.showNotification('Changes saved successfully!'); 
      }
    });
  }

  cancelButtonClicked() {
    this.cancelButtonState = 'clicked';
    setTimeout(() => { this.cancelButtonState = 'idle'; }, 200);

    this.canEdit = false;

    this.showNotification('No changes were made'); 
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }
}
