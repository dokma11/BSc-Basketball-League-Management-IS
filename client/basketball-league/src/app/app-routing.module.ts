import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './feature-modules/layout/home/home.component';
import { LoginComponent } from './infrastructure/auth/login/login.component';
import { RegistrationComponent } from './infrastructure/auth/registration/registration.component';
import { RosterComponent } from './feature-modules/roster-management/roster/roster.component';
import { LeagueNewsComponent } from './feature-modules/news/league-news/league-news.component';
import { TradeManagementComponent } from './feature-modules/trades/trade-management/trade-management.component';
import { MyProfileComponent } from './feature-modules/profile/my-profile/my-profile.component';

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
  },
  {
    path: 'my-profile',
    component: MyProfileComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {
    scrollPositionRestoration: 'enabled', 
    anchorScrolling: 'enabled', 
    scrollOffset: [0, 0] 
  })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
