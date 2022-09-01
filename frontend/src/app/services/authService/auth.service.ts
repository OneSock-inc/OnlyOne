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
  ) { }

  private error: any;

  /**
   * Sends a login request to backend and expects a JWT token in return.
   * Set its jwt param value and save it in localStorage.
   * @param username provided by front end user
   * @param password provided by front end user
   * @returns Observable
   */
  login(username?: string, password?: string) {
    return this.http
      .post<JWToken>(this.backendLink.getLoginUrl(), { username, password })
      .subscribe({
        next: (data: JWToken) => {
          this.tokenService.setAuthorizationToken(data);
          // TODO: jwt token validator
          // this.jwt = { ...data };
        }, // success path
        error: (error) => { 
          this.error = error; // error path
        }
      });
  }

  logout(){
    localStorage.clear();
    this.tokenService.clearToken();
  }

  isLoggedIn(): boolean {
    return this.tokenService.getAuthorizationToken() !== '';
  }

  getError() {
    return this.error;
  }

  clearError() {
    this.error = undefined;
  }

}
