import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MaterialModule } from 'src/app/infrastructure/material/material-module';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { MyProfileComponent } from './my-profile/my-profile.component';

@NgModule({
  declarations: [
    MyProfileComponent
  ],
  imports: [
    CommonModule,
    RouterModule,
    ReactiveFormsModule,
    MaterialModule,
    FontAwesomeModule,
  ],
  exports: [
    MyProfileComponent
  ]
})
export class ProfileModule { }
