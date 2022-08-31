import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BackendLinkService } from '../backendservice/backend-link.service';
import { User } from '../../dataModel/user.model';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  constructor(private http: HttpClient, private backendLink: BackendLinkService) { }

  user?: User;

  private error:any;

  login(username?: string, password?: string) {
    return this.http.post<User>(this.backendLink.getLoginUrl(), {username, password}).subscribe({
      next: (data: User) => {
        this.user = {...data};
        console.log(this.user);
      }, // success path
      error: error => this.error = error, // error path
    });
  }

  getError() {
    return this.error;
  }

  clearError() {
    this.error = undefined;
  }



}

// https://blog.angular-university.io/angular-jwt-authentication/