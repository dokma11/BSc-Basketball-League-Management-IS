import { Component } from '@angular/core';
import { faUser, faSignOut, faSignIn, faPencilSquare, faHome, faBinoculars, faAddressBook, faEnvelopeOpen, faAddressCard, faNewspaper, faBasketball } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Uloga } from 'src/app/shared/model/player.model';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {
  user: User | undefined;
  regrutUloga: Uloga = Uloga.UloRegrut;

  constructor(private authService: AuthService) {}

  ngOnInit(): void {
    this.authService.user$.subscribe(user => {
      this.user = user;
    });
  }

  onLogout(): void {
    this.authService.logout();
  }

  faUser = faUser;
  faSignOut = faSignOut;
  faSignIn = faSignIn;
  faPencilSquare = faPencilSquare;
  faHome = faHome;
  faAddressBook = faAddressBook;
  faEnvelopeOpen = faEnvelopeOpen;
  faAddressCard = faAddressCard;
  faNewsPaper = faNewspaper;
  faBasketball = faBasketball;
  faBinoculars = faBinoculars;
}
