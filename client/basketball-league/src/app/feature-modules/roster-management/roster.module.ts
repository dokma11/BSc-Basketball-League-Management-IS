import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MaterialModule } from 'src/app/infrastructure/material/material-module';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { RosterComponent } from './roster/roster.component';
import { PlayerCardComponent } from './player-card/player-card.component';

@NgModule({
  declarations: [
    RosterComponent,
    PlayerCardComponent
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
