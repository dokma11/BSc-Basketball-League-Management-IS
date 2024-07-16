import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';

@Component({
  selector: 'app-pick-card',
  templateUrl: './pick-card.component.html',
  styleUrls: ['./pick-card.component.css'],
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
export class PickCardComponent implements OnInit{
  addAssetButtonState: string = 'idle';  
  removeAssetButtonState: string = 'idle';
  assetSelected: boolean = false;
  player: string = 'IGRAC';  // DOK NE POVEZEM SA BEKOM, PIK TREBA DA BUDE OVO STO SU PLAYER I REQUEST
  //@Input() request!: PersonalTourRequest; Ovde treba da budu pikovi kao lista
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  ownTeam: boolean = false;

  constructor(private dialog: MatDialog,
              private authService: AuthService,
              private snackBar: MatSnackBar) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  
    // Treba nekako videti za koji tim se traze pikovi i onda ih ucitati u this.picks recimo
  }

  ngOnInit(): void {
    // TODO: Dodati ovde sta treba pri inicijalizaciji komponenti
  }

  addAssetButtonClicked(assset: any): void {
    this.addAssetButtonState = 'clicked';
    setTimeout(() => { this.addAssetButtonState = 'idle'; }, 200);

    this.assetSelected = true;
    this.showNotification("Pick successfully added!");
    
    // TODO: Dodati logiku za dodavanje imovine na neku listu itd, trebalo bi da ima na isi
  }

  removeAssetButtonClicked(assset: any): void {
    this.removeAssetButtonState = 'clicked';
    setTimeout(() => { this.removeAssetButtonState = 'idle'; }, 200);

    this.assetSelected = false;
    this.showNotification("Pick successfully removed!");

    // TODO: Dodati logiku za uklanjanje imovine sa neke liste itd, trebalo bi da ima na isi
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }

  faPlus = faPlus;
  faMinus = faMinus;
}
