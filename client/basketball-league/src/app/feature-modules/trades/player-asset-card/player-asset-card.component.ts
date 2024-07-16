import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faHeart, faHandPaper, faList, faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';

@Component({
  selector: 'app-player-asset-card',
  templateUrl: './player-asset-card.component.html',
  styleUrls: ['./player-asset-card.component.css']
})
export class PlayerAssetCardComponent implements OnInit{
  addAssetButtonState: string = 'idle';
  removeAssetButtonState: string = 'idle';
  assetSelected: boolean = false;
  player: string = 'IGRAC';  // DOK NE POVEZEM SA BEKOM
  //@Input() request!: PersonalTourRequest;
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();

  constructor(private dialog: MatDialog,
              private authService: AuthService,
              private snackBar: MatSnackBar) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    
  }

  addAssetButtonClicked(assset: any): void {
    this.addAssetButtonState = 'clicked';
    setTimeout(() => { this.addAssetButtonState = 'idle'; }, 200);

    this.assetSelected = true;
    this.showNotification("Player successfully added!");
    
    // TODO: Dodati logiku za dodavanje imovine na neku listu itd, trebalo bi da ima na isi
  }

  removeAssetButtonClicked(assset: any): void {
    this.removeAssetButtonState = 'clicked';
    setTimeout(() => { this.removeAssetButtonState = 'idle'; }, 200);

    this.assetSelected = false;
    this.showNotification("Player successfully removed!");

    // TODO: Dodati logiku za uklanjanje imovine sa neke liste itd, trebalo bi da ima na isi
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }
  faMinus = faMinus;
  faPlus = faPlus;
}
