import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
  animations: [
    trigger('fadeInOut', [
      transition(':enter', [
        style({ opacity: 0 }),
        animate('500ms ease-out', style({ opacity: 1 })),
      ]),
      transition(':leave', [
        animate('500ms ease-in', style({ opacity: 0 })),
      ]),
    ]), 
    trigger('buttonState', [
      state('clicked', style({
        transform: 'scale(0.9)',
        opacity: 0.5
      })),
      transition('* => clicked', [
        animate('200ms')
      ]),
      transition('clicked => idle', [
        animate('200ms')
      ])
    ]),
  ]
})
export class HomeComponent {
  rosterButtonState: string = 'idle'; 
  tradesButtonState: string = 'idle'; 
  recruitsButtonState: string = 'idle'; 

  constructor(private router: Router) {
  }
  
  cardGroupClicked(){
    console.log('click');
  }

  rosterButtonClicked(){
    this.rosterButtonState = 'clicked'; 
    setTimeout(() => { this.rosterButtonState = 'idle'; }, 200);
    this.router.navigate(['/roster']); 
  }

  tradesButtonClicked(){
    this.tradesButtonState = 'clicked'; 
    setTimeout(() => { this.tradesButtonState = 'idle'; }, 200);
    this.router.navigate(['/trade-management']); 
  }

  recruitsButtonClicked(){
    this.recruitsButtonState = 'clicked'; 
    setTimeout(() => { this.recruitsButtonState = 'idle'; }, 200);
    this.router.navigate(['/recruit-management']); 
  }
}
