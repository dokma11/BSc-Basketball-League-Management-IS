import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './feature-modules/layout/home/home.component';
import { LoginComponent } from './infrastructure/auth/login/login.component';
import { RegistrationComponent } from './infrastructure/auth/registration/registration.component';
import { LeagueNewsComponent } from './feature-modules/league-news/league-news.component';
import { TradeManagementComponent } from './feature-modules/trade-management/trade-management.component';
import { RosterComponent } from './feature-modules/roster/roster.component';

const routes: Routes = [
  {
    path: '',
    component: HomeComponent
  },
  {
    path: 'login',
    component: LoginComponent
  },
  {
    path: 'register',
    component: RegistrationComponent
  },
  {
    path: 'roster',
    component: RosterComponent
  },
  {
    path: 'trade-management',
    component: TradeManagementComponent
  },
  {
    path: 'league-news',
    component: LeagueNewsComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
