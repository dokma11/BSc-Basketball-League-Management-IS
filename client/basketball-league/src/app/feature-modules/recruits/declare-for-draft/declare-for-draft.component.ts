import { Component } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { faEye, faEyeSlash } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Pozicija } from 'src/app/shared/model/player.model';
import { Recruit } from 'src/app/shared/model/recruit.model';
import { RecruitsService } from '../recruits.service';
import { trigger, transition, style, animate, state } from '@angular/animations';

@Component({
  selector: 'app-declare-for-draft',
  templateUrl: './declare-for-draft.component.html',
  styleUrls: ['./declare-for-draft.component.css'],
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
export class DeclareForDraftComponent {
  buttonState: string = 'idle'; 
  focused: string = '';
  user: User | undefined;

  constructor(private authService: AuthService, 
              private recruitsService: RecruitsService,
              private router: Router,
              private snackBar: MatSnackBar,) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  }

  registrationForm = new FormGroup({
    placeOfBirth: new FormControl('', [Validators.required]),
    height: new FormControl('', [Validators.required]),
    weight: new FormControl('', [Validators.required]),
    position: new FormControl('', [Validators.required]),
    phoneNumber: new FormControl('', [Validators.required]),
    selectedPosition: new FormControl('', [Validators.required]),
  });

  register(): void {
    let pozicija : Pozicija;

    if (this.registrationForm.value.position === 'PG') {
      pozicija = Pozicija.PG;
    } else if (this.registrationForm.value.position === 'SG') {
      pozicija = Pozicija.SG;
    } else if (this.registrationForm.value.position === 'SF') {
      pozicija = Pozicija.SF;
    } else if (this.registrationForm.value.position === 'PF') {
      pozicija = Pozicija.PF;
    } else {
      pozicija = Pozicija.C;
    }
 
    const recruit: Recruit = {
      id: this.user!.id,
      email: this.user?.email!,
      konTelefonReg: this.registrationForm.value.phoneNumber || "",
      visReg: this.registrationForm.value.height?.toString() || "",
      tezReg: this.registrationForm.value.weight?.toString() || "",
      pozReg: pozicija,
      prosOcReg: '90',
      prosRankReg: '5'
    }

    console.log(recruit);

    this.buttonState = 'clicked'; 
    setTimeout(() => { this.buttonState = 'idle'; }, 200); 

    this.recruitsService.declareForDraft(recruit).subscribe({
      next: (result: Recruit) => {
        this.showNotification('Successfully declared for the upcoming draft');
        this.router.navigate(['']);
      }
    });
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
