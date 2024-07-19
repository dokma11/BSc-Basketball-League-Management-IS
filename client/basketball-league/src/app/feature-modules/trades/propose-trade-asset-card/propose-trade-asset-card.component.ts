import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';

@Component({
  selector: 'app-propose-trade-asset-card',
  templateUrl: './propose-trade-asset-card.component.html',
  styleUrls: ['./propose-trade-asset-card.component.css'],
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
export class ProposeTradeAssetCardComponent implements OnInit{
  removeAssetButtonState: string = 'idle';
  player: string = 'IGRAC';  // DOK NE POVEZEM SA BEKOM
  description: string = ''; // Ovo treba da bdue ono sto ce se pokazati u samoj kartici
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  ownTeam: boolean = false;
  @Input() showButton: boolean = true;

  constructor(private dialog: MatDialog,
              private authService: AuthService,
              private snackBar: MatSnackBar) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  
    // TODO: Treba napraviti semu za prikaz svakog aseta posebno, sta pokazati itd.
  }

  ngOnInit(): void {
    // TODO: Dodati ovde sta treba pri inicijalizaciji komponenti
  }

  removeAssetButtonClicked(assset: any): void {
    this.removeAssetButtonState = 'clicked';
    setTimeout(() => { this.removeAssetButtonState = 'idle'; }, 200);

    this.showNotification("Asset successfully removed!");

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
