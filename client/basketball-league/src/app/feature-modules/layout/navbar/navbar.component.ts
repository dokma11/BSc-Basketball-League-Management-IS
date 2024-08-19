import { Component } from '@angular/core';
import { faUser, faSignOut, faSignIn, faPencilSquare, faHome, faBinoculars, faAddressBook, faEnvelopeOpen, faAddressCard, faNewspaper, faBasketball, faPenClip } from '@fortawesome/free-solid-svg-icons';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Uloga } from 'src/app/shared/model/player.model';
import { RecruitsService } from '../../recruits/recruits.service';
import { Recruit } from 'src/app/shared/model/recruit.model';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {
  user: User | undefined;
  regrutUloga: Uloga = Uloga.UloRegrut;
  declaredRecruits: Recruit[] = [];
  showDeclaration: boolean = true;

  constructor(private authService: AuthService,
              private recruitsServie: RecruitsService) {}

  ngOnInit(): void {
    this.authService.user$.subscribe(user => {
      this.user = user;

      if(this.user.uloga == this.regrutUloga) {
        this.recruitsServie.getAllRecruits().subscribe({
          next: (result: Recruit[] | Recruit) => {
            if(Array.isArray(result)) {
              this.declaredRecruits = result;

              this.declaredRecruits.forEach(recruit => {
                if (recruit.id == this.user?.id) {
                  this.showDeclaration = false;
                }
              })
            }
          }
        });
      }
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
  faPenClip = faPenClip;
}
