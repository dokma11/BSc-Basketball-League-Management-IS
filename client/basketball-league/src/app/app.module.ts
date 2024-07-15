import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from './infrastructure/material/material-module';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { LayoutModule } from "./feature-modules/layout/layout.module";
import { HttpClientModule } from '@angular/common/http';
import { AuthModule } from './infrastructure/auth/auth.module';
import { TradesModule } from './feature-modules/trades/trades.module';
import { NewsModule } from './feature-modules/news/news.module';
import { RosterModule } from './feature-modules/roster-management/roster.module';

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MaterialModule,
    FontAwesomeModule,
    LayoutModule,
    HttpClientModule,
    AuthModule,
    TradesModule,
    NewsModule,
    RosterModule
],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
