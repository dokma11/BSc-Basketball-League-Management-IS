import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, tap } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { TokenStorage } from './jwt/token.service';
import { environment } from 'src/env/environment';
import { AuthenticationResponse } from './model/authentication-response.model';
import { Login } from './model/login.model';
import { User } from './model/user.model';
import { EditEmployee } from './model/editEmployee.model';
import { Registration } from './model/regsitration.model';
import { JwtHelperService } from '@auth0/angular-jwt';

@Injectable({
  providedIn: 'root',
})
export class AuthService {

  user$ = new BehaviorSubject<User>({ email: '', id: -100, uloga: 0, teamId: -100 });
  basePath = environment.apiHost + 'auth/';

  constructor(
    private http: HttpClient,
    private tokenStorage: TokenStorage,
    private router: Router
  ) {}

  login(login: Login): Observable<AuthenticationResponse> {
    return this.http
      .post<AuthenticationResponse>(environment.host + 'login', login)
      .pipe(
        tap((authenticationResponse) => {
          this.tokenStorage.saveAccessToken(authenticationResponse.token);
          this.setUser();
        })
      );
  }

  register(registration: Registration): Observable<AuthenticationResponse> {
    return this.http
      .post<AuthenticationResponse>(environment.host + 'register', registration)
      .pipe(
        tap((authenticationResponse) => {
          this.tokenStorage.saveAccessToken(authenticationResponse.token);
          this.setUser();
        })
      );
  }

  registerEmployee(registration: Registration): Observable<AuthenticationResponse> {
    return this.http
      .post<AuthenticationResponse>(this.basePath + 'registerEmployee', registration);
  }

  getEmployeeById(id: number): Observable<EditEmployee> {
    console.log('get employee');
    console.log(id);
    return this.http.get<EditEmployee>(this.basePath + `${id}`)
  }

  updateEmployee(employee: EditEmployee, id: number): Observable<Registration> {
    return this.http.put<Registration>(this.basePath + `updateEmployee/${id}`, employee);
  }

  logout(): void {
    this.router.navigate(['/']).then((_) => {
      this.tokenStorage.clear();
      this.user$.next({ email: '', id: 0, uloga: 0 });
    });
  }

  checkIfUserExists(): void {
    const accessToken = this.tokenStorage.getAccessToken();
    if (accessToken == null) {
      return;
    }
    this.setUser();
  }

  private setUser(): void {
    const jwtHelperService = new JwtHelperService();
    const accessToken = this.tokenStorage.getAccessToken() || '';
    const user: User = {
      id: +jwtHelperService.decodeToken(accessToken).id,
      email: jwtHelperService.decodeToken(accessToken).username,
      uloga: jwtHelperService.decodeToken(accessToken).role,
      teamId: jwtHelperService.decodeToken(accessToken).teamId,
    };
    this.user$.next(user);
  }

  setToken(authenticationResponse: AuthenticationResponse) {
    this.tokenStorage.saveAccessToken(authenticationResponse.token);
    this.setUser();
  }

  getJwtToken(): string | null {
    return this.tokenStorage.getAccessToken(); 
  }

}