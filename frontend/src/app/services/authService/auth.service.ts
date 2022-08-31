import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BackendLinkService } from '../backendservice/backend-link.service';
import { User } from 'src/app/dataModel/user.model';
import { JWToken } from 'src/app/dataModel/jwt.model';
import { TokenService } from './token-service.service';

// Inspired by : https://blog.angular-university.io/angular-jwt-authentication/

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  constructor(
    private http: HttpClient,
    private backendLink: BackendLinkService,
    private tokenService: TokenService
  ) {}

  user?: User;

  private jwt?: JWToken;
  private error: any;

  login(username?: string, password?: string) {
    return this.http
      .post<JWToken>(this.backendLink.getLoginUrl(), { username, password })
      .subscribe({
        next: (data: JWToken) => {
          this.tokenService.setAutoriuationToken(data);
          // this.jwt = { ...data };
          console.log(this.jwt);
        }, // success path
        error: (error) => { 
          console.error(error);
          this.error = error; // error path
        }
      });
  }

  getAuthorizationToken(): string {
    if (typeof this.jwt !== 'undefined') return this.jwt.token;
    else return '';
  }

  getError() {
    return this.error;
  }

  clearError() {
    this.error = undefined;
  }

}
