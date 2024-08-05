import { Component, Inject, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { ProposeTradeFormComponent } from '../propose-trade-form/propose-trade-form.component';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { RosterService } from '../../roster-management/roster.service';
import { Player } from 'src/app/shared/model/player.model';
import { DraftRight } from 'src/app/shared/model/draftRight.model';
import { Pick } from 'src/app/shared/model/pick.model';
import { Team } from 'src/app/shared/model/team.model';

@Component({
  selector: 'app-asset-choosing-form',
  templateUrl: './asset-choosing-form.component.html',
  styleUrls: ['./asset-choosing-form.component.css'],
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
export class AssetChoosingFormComponent implements OnInit{
  finishButtonState: string = 'idle';
  focused: string = '';
  team: Team | undefined;
  picks: Pick[] = [];
  players: Player[] = [];
  draftRights: DraftRight[] = [];
  chosenPicks: Pick[] = [];
  chosenPlayers: Player[] = [];
  chosenDraftRights: DraftRight[] = [];
  assetForm = new FormGroup({
    selectedAssetType: new FormControl('Players', [Validators.required]),
  });

  constructor(private snackBar: MatSnackBar,
              private dialogRef: MatDialogRef<ProposeTradeFormComponent>,
              private rosterService: RosterService,
              @Inject(MAT_DIALOG_DATA) public data: any,) {
    this.team = data.team;
    this.chosenPicks = data.chosenPicks;
    this.chosenPlayers = data.chosenPlayers;
    this.chosenDraftRights = data.chosenDraftRights;
  }

  ngOnInit(): void {
    this.getAssets();
  }

  addTourButtonClicked() {
    
  }

  finishButtonClicked(): void {
    this.finishButtonState = 'clicked';
    setTimeout(() => { this.finishButtonState = 'idle'; }, 200);
    this.showNotification('Assets successfully selected!');

    // TODO: Dodati neko prosledjivanje liste odabranih aseta i sl, nekako moram povezati sa samim karticama

    this.dialogRef.close();
  }

  onAssetTypeChange(event: any) {
    this.showNotification('Selected asset type: ' + this.assetForm.value.selectedAssetType);
    this.getAssets();
  }

  getAssets() {
      if (this.assetForm.value.selectedAssetType === 'Players') {
        this.rosterService.getAllAvailablePlayersByTeamId(this.team?.idTim!).subscribe({
          next: (result: Player[] | Player) => {
            if(Array.isArray(result)){
              this.players = result;
              // Reset other unncessary lists
              this.picks = [];
              this.draftRights = [];
            }
          }
        })
      } 
      else if (this.assetForm.value.selectedAssetType === 'Picks') {
        this.rosterService.getAllAvailablePicksByTeamId(this.team?.idTim!).subscribe({
          next: (result: Pick[] | Pick) => {
            if(Array.isArray(result)){
              this.picks = result;
              // Reset other unncessary lists
              this.players = [];
              this.draftRights = [];
            }
          }
        })
      }  
      else if (this.assetForm.value.selectedAssetType === 'Draft rights') {
        this.rosterService.getAllAvailableDraftRightsByTeamId(this.team?.idTim!).subscribe({
          next: (result: DraftRight[] | DraftRight) => {
            if(Array.isArray(result)){
              this.draftRights = result;
              // Reset other unncessary lists
              this.picks = [];
              this.players = [];
            }
          }
        })
      }  
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }

  overviewClicked(){
    this.dialogRef.close();
  }

  faPlus = faPlus;
  faMinus = faMinus;
}
