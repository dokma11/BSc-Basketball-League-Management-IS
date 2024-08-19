import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, HostListener } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { faEye, faEyeSlash } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from '../auth.service';
import { Registration } from '../model/regsitration.model';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css'],
  animations: [
    trigger('fadeInOut', [
      transition(':enter', [
        style({ opacity: 0 }),
        animate('500ms ease-out', style({ opacity: 1 })),
      ]),
      transition(':leave', [
        animate('500ms ease-in', style({ opacity: 0 })),
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
export class RegistrationComponent {
  isPasswordVisible: boolean;
  isRepeatPasswordVisible: boolean;

  constructor(private authService: AuthService, 
              private router: Router,
              private snackBar: MatSnackBar,) {
    this.isPasswordVisible = false;
    this.isRepeatPasswordVisible = false;
  }

  registrationForm = new FormGroup({
    firstName: new FormControl('', [Validators.required]),
    lastName: new FormControl('', [Validators.required]),
    email: new FormControl('', [Validators.required]),
    dateOfBirth: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required]),
    repeatpassword: new FormControl('', [Validators.required]),
  });

  buttonState: string = 'idle'; 
  focused: string = '';
  backgroundSize: string = '100% 110%';

  register(): void {
    const registration: Registration = {
      ime: this.registrationForm.value.firstName || "",
      prezime: this.registrationForm.value.lastName || "",
      email: this.registrationForm.value.email || "",
      lozinka: this.registrationForm.value.password || "",
      datRodj: new Date(this.registrationForm.value.dateOfBirth!),
      role: 0
    };

    console.log(registration);

    if (this.registrationForm.valid) {
      if(this.registrationForm.value.password === this.registrationForm.value.repeatpassword){
        this.buttonState = 'clicked'; 
        setTimeout(() => { this.buttonState = 'idle'; }, 200); 
        this.authService.register(registration).subscribe({
          next: () => {
            this.router.navigate(['']);
          },
          error: (error) => {
            if (error.error instanceof ErrorEvent) {
                alert('An error occurred: ' + error.error.message);
            } else {
              if (error.status === 400) { 
                const errorMessage: string = error.error.message || "";

                if (errorMessage.includes('email')) {
                    this.showNotification('Email already exists. Please use a different one.');
                }
            } else {
              this.showNotification('Server error occurred. Please try again later.');
            }
            }
        }
        });
      } 
      else{
        this.showNotification('Passwords do not match!'); 
      }
    }
    else{
      this.showNotification('Sign up form not valid!');
    }
  }

  togglePasswordVisibility() {
    this.isPasswordVisible = !this.isPasswordVisible;
  }

  toggleRepeatPasswordVisibility() {
    this.isRepeatPasswordVisible = !this.isRepeatPasswordVisible;
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
      panelClass: ['custom-snackbar'] 
    });
  }

  faEye = faEye;
  faEyeSlash = faEyeSlash;
}
