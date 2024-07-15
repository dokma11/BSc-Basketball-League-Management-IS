import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { ProposeTradeFormComponent } from '../propose-trade-form/propose-trade-form.component';

@Component({
  selector: 'app-trade-management',
  templateUrl: './trade-management.component.html',
  styleUrls: ['./trade-management.component.css'],
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
export class TradeManagementComponent implements OnInit{
  user: User | undefined;
  backgroundSize: string = '100% 100%';
  //requests: PersonalTourRequest[] = [];   ovde treba da budu zahtevi za trejdove bas
  proposeTradeButtonState: string = "";

  private dialogRef: any;

  constructor(private authService: AuthService,
              private dialog: MatDialog,
              private snackBar: MatSnackBar,) {

  }

  ngOnInit(): void {
    this.getRequests();
  }

  getRequests() {
    // this.authService.user$.subscribe(user => {
    //   this.user = user;
    //   if(this.user.role === 'GUEST'){
    //     this.toursService.getGuestsTourRequests(this.user.id).subscribe({
    //       next: (result: PersonalTourRequest[] | PersonalTourRequest) => {
    //         if(Array.isArray(result)){
    //           this.requests = result;
    //         }
    //       }
    //     });
    //   }
    //   else{
    //     this.toursService.getTourRequestsOnHold().subscribe({
    //       next: (result: PersonalTourRequest[] | PersonalTourRequest) => {
    //         if(Array.isArray(result)){
    //           this.requests = result;
    //         }
    //       }
    //     });
    //   }
    // });
  }

  handleDialogClosed(result: any) {
    this.getRequests();
  }

  proposeTradeButtonClicked() {
    this.proposeTradeButtonState = 'clicked';
    setTimeout(() => { this.proposeTradeButtonState = 'idle'; }, 200);
    this.dialogRef = this.dialog.open(ProposeTradeFormComponent, {
    });

    if (this.dialogRef) {
      this.dialogRef.afterClosed().subscribe((result: any) => {
        this.getRequests();
      });
    }
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }

  faPlus = faPlus;
}
