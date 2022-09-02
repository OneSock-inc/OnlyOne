import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BackendLinkService } from '../backendservice/backend-link.service';
import { JWToken } from 'src/app/dataModel/jwt.model';
import { TokenService } from './token-service.service';
import { Subscription } from 'rxjs';

// Inspired by : https://blog.angular-university.io/angular-jwt-authentication/

/**
 * This service provide some useful methods to login, logout and know if
 * the user is logged in.
 */
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
  login(username?: string, password?: string): Subscription {
    return this.http
      .post<JWToken>(this.backendLink.getLoginUrl(), { username, password })
      .subscribe({
        next: (data: JWToken) => {
          this.tokenService.setAuthorizationToken(data);
          // TODO: jwt token validator
        }, // success path
        error: (error) => { 
          this.error = error; // error path
        }
      });
  }

  loginV2(username: string, password: string, successCallback: Function, errorCallBack: Function) {
    this.http.post<JWToken>(this.backendLink.getLoginUrl(), {username, password})
    .subscribe({
      next: (response: any) => {
        successCallback(response);
      },
      error: (error: any) => {
        errorCallBack(error);
      }
    })
  }

  /**
   * Clear the local storage and the token saved in the TokenService 
   * instance.
   */
  logout() {
    localStorage.clear();
    this.tokenService.clearToken();
  }

  /**
   * Test if user is logged.
   * @returns true if the current user is logged in.
   */
  isLoggedIn(): boolean {
    return this.tokenService.getAuthorizationToken() !== '';
  }

  /**
   * Retrieve error occured on http request if any.
   * @returns error
   */
  getError() {
    return this.error;
  }

  /**
   * Clear the error saved by AuthService instance.
   * Must be called before made a new call to login method.
   */
  clearError() {
    this.error = undefined;
  }

}
