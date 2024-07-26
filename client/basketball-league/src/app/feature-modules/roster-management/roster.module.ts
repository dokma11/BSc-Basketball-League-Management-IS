import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MaterialModule } from 'src/app/infrastructure/material/material-module';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { RosterComponent } from './roster/roster.component';
import { PlayerCardComponent } from './player-card/player-card.component';
import { AddPlayerToListPromptComponent } from './add-player-to-list-prompt/add-player-to-list-prompt.component';
import { PickRosterCardComponent } from './pick-roster-card/pick-roster-card.component';
import { DraftRightRosterCardComponent } from './draft-right-roster-card/draft-right-roster-card.component';

@NgModule({
  declarations: [
    RosterComponent,
    PlayerCardComponent,
    AddPlayerToListPromptComponent,
    PickRosterCardComponent,
    DraftRightRosterCardComponent
  ],
  imports: [
    CommonModule,
    RouterModule,
    ReactiveFormsModule,
    MaterialModule,
    FontAwesomeModule,
  ],
  exports: [
    RosterComponent,
    PlayerCardComponent
  ]
})
export class RosterModule { }
