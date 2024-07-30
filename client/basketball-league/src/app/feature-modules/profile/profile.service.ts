import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { environment } from 'src/env/environment';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor(private http: HttpClient) { }

  getUserByID(userId: number): Observable<User> {
    return this.http.get<User>(environment.host + 'user/' + userId);
  }

  updateUserProfile(user: User): Observable<User> {
    return this.http.put<User>(environment.host + 'user', user);
  }

}
