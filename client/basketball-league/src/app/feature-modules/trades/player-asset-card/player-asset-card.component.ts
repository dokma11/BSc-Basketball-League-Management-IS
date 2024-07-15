import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { faHeart, faHandPaper, faList, faPlus } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';

@Component({
  selector: 'app-player-asset-card',
  templateUrl: './player-asset-card.component.html',
  styleUrls: ['./player-asset-card.component.css']
})
export class PlayerAssetCardComponent implements OnInit{
  addAssetButtonState: string = 'idle';
  player: string = 'IGRAC';  // DOK NE POVEZEM SA BEKOM
  //@Input() request!: PersonalTourRequest;
  private dialogRef: any;
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();

  constructor(private dialog: MatDialog,
              private authService: AuthService) {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  }

  ngOnInit(): void {
    
  }

  addAssetButtonClicked(player: any){

  }

  faHeart = faHeart;
  faHandPaper = faHandPaper;
  faList = faList;
  faPlus = faPlus;
}
