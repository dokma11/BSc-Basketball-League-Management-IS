import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from './infrastructure/material/material-module';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { LayoutModule } from "./feature-modules/layout/layout.module";
import { HttpClientModule } from '@angular/common/http';

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

],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
