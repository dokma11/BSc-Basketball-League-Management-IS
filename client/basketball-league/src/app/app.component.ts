import { Component } from '@angular/core';
import { AuthService } from './infrastructure/auth/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'basketall-league';

  constructor(
    private authService: AuthService,
  ) {}


  ngOnInit(): void {
    this.checkIfUserExists();
  }
  
  private checkIfUserExists(): void {
    this.authService.checkIfUserExists();
  }
  
}
