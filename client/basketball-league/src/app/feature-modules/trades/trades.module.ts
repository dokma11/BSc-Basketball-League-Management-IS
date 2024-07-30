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
import { NgxMatSelectSearchModule } from 'ngx-mat-select-search';
import { ProposeTradeAssetCardComponent } from './propose-trade-asset-card/propose-trade-asset-card.component';
import { AcceptRequestPromptComponent } from './accept-request-prompt/accept-request-prompt.component';
import { DeclineRequestPromptComponent } from './decline-request-prompt/decline-request-prompt.component';
import { ShowRequestDetailsPromptComponent } from './show-request-details-prompt/show-request-details-prompt.component';
import { SeeDenialExplanationPromptComponent } from './see-denial-explanation-prompt/see-denial-explanation-prompt.component';
import { CancelRequestPromptComponent } from './cancel-request-prompt/cancel-request-prompt.component';

@NgModule({
  declarations: [
    ProposeTradeFormComponent,
    TradeManagementComponent,
    TradeRequestCardComponent,
    AssetChoosingFormComponent,
    PickCardComponent,
    PlayerAssetCardComponent,
    ProposeTradeAssetCardComponent,
    AcceptRequestPromptComponent,
    DeclineRequestPromptComponent,
    ShowRequestDetailsPromptComponent,
    SeeDenialExplanationPromptComponent,
    CancelRequestPromptComponent,
  ],
  imports: [
    CommonModule,
    RouterModule,
    ReactiveFormsModule,
    MaterialModule,
    FontAwesomeModule,
    RosterModule,
    FormsModule,
    NgxMatSelectSearchModule
],
  exports: [
    ProposeTradeFormComponent,
    TradeManagementComponent,
    TradeRequestCardComponent
  ]
})
export class TradesModule { }
