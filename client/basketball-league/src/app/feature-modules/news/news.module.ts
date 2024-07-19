import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MaterialModule } from 'src/app/infrastructure/material/material-module';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { LeagueNewsComponent } from './league-news/league-news.component';
import { NewsCardComponent } from './news-card/news-card.component';

@NgModule({
  declarations: [
    LeagueNewsComponent,
    NewsCardComponent
  ],
  imports: [
    CommonModule,
    RouterModule,
    ReactiveFormsModule,
    MaterialModule,
    FontAwesomeModule,
  ],
  exports: [
    LeagueNewsComponent,
    NewsCardComponent
  ]
})
export class NewsModule { }
