import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MaterialModule } from 'src/app/infrastructure/material/material-module';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { RosterComponent } from './roster/roster.component';
import { PlayerCardComponent } from './player-card/player-card.component';
import { AddPlayerToListPromptComponent } from './add-player-to-list-prompt/add-player-to-list-prompt.component';

@NgModule({
  declarations: [
    RosterComponent,
    PlayerCardComponent,
    AddPlayerToListPromptComponent
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
