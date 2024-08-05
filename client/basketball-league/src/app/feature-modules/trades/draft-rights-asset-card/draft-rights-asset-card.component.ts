import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { DraftRight } from 'src/app/shared/model/draftRight.model';
import { Recruit } from 'src/app/shared/model/recruit.model';
import { RosterService } from '../../roster-management/roster.service';

@Component({
  selector: 'app-draft-rights-asset-card',
  templateUrl: './draft-rights-asset-card.component.html',
  styleUrls: ['./draft-rights-asset-card.component.css']
})
export class DraftRightsAssetCardComponent implements OnInit{
  addAssetButtonState: string = 'idle';  
  removeAssetButtonState: string = 'idle';
  assetSelected: boolean = false;
  @Input() draftRight!: DraftRight;
  draftRightPlayer: Recruit | undefined;
  @Input() chosenDraftRights!: DraftRight[];
  user: User | undefined;
  @Output() dialogRefClosed: EventEmitter<any> = new EventEmitter<any>();
  ownTeam: boolean = false;
  age: string = '';

  constructor(private snackBar: MatSnackBar,
              private rosterService: RosterService) { }

  ngOnInit(): void {
    if(this.chosenDraftRights){
      this.chosenDraftRights.forEach(dr => {
        if(dr.idPrava === this.draftRight.idPrava){
          this.assetSelected = true;
        }
      })
    }

    this.rosterService.getRecruitById(this.draftRight.idRegrut).subscribe({
      next: (result: Recruit) => {
        this.draftRightPlayer = result;

        const today = new Date();
        const birthDate = new Date(this.draftRightPlayer.datRodj!);
        this.age = (today.getFullYear() - birthDate.getFullYear()).toString();
      }
    });
  }

  addAssetButtonClicked(asset: DraftRight): void {
    this.addAssetButtonState = 'clicked';
    setTimeout(() => { this.addAssetButtonState = 'idle'; }, 200);

    let alreadyThere = false;
    this.chosenDraftRights.forEach(dr => {
        if(dr.idPrava === asset.idPrava){
          alreadyThere = true
        }
      }
    )

    if(!alreadyThere){
      this.chosenDraftRights.push(asset);
      this.assetSelected = true;
      this.showNotification("Draft Rights successfully added!");
    }
    else{
      this.showNotification("Draft Rights already chosen!");
    }
  }

  removeAssetButtonClicked(asset: DraftRight): void {
    this.removeAssetButtonState = 'clicked';
    setTimeout(() => { this.removeAssetButtonState = 'idle'; }, 200);

    this.chosenDraftRights.forEach( (dr, index) => {
        if(dr.idPrava === asset.idPrava){
          this.chosenDraftRights.splice(index,1)
          this.assetSelected = false;
          this.showNotification("Draft Rights successfully removed!");
        }
      }
    )
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
