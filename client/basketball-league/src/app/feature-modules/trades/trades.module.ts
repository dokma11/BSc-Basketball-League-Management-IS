import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MaterialModule } from 'src/app/infrastructure/material/material-module';
import { RouterModule } from '@angular/router';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { ProposeTradeFormComponent } from './propose-trade-form/propose-trade-form.component';
import { TradeManagementComponent } from './trade-management/trade-management.component';
import { TradeRequestCardComponent } from './trade-request-card/trade-request-card.component';
import { AssetChoosingFormComponent } from './asset-choosing-form/asset-choosing-form.component';
import { PickCardComponent } from './pick-card/pick-card.component';
import { RosterModule } from "../roster-management/roster.module";
import { PlayerAssetCardComponent } from './player-asset-card/player-asset-card.component';

@NgModule({
  declarations: [
    ProposeTradeFormComponent,
    TradeManagementComponent,
    TradeRequestCardComponent,
    AssetChoosingFormComponent,
    PickCardComponent,
    PlayerAssetCardComponent,
  ],
  imports: [
    CommonModule,
    RouterModule,
    ReactiveFormsModule,
    MaterialModule,
    FontAwesomeModule,
    RosterModule,
    FormsModule
],
  exports: [
    ProposeTradeFormComponent,
    TradeManagementComponent,
    TradeRequestCardComponent
  ]
})
export class TradesModule { }
