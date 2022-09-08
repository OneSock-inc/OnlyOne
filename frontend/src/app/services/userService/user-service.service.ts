import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BackendLinkService } from '../backendservice/backend-link.service';

import { User } from 'src/app/dataModel/user.model';
import { concatAll, map, Observable } from 'rxjs';
import { JWToken } from 'src/app/dataModel/index.model';
import { TokenService } from '../authService/token-service.service';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  constructor(
    private http: HttpClient,
    private backSrv: BackendLinkService,
    private tokenService: TokenService
  ) {
    if (!this.isLoggedIn()) {
      this.user$ = new Observable<User>((s) => s.next(new User()));
    } else {
      this.user$ = this.fetchUserById(this.getUserId());
    }
  }

  private user$: Observable<User> = new Observable<User>((s) =>
    s.next(new User())
  );

  /**
   * Sends a login request to backend.
   * Set its jwt param value and save it in localStorage.
   * @param username provided by front end user
   * @param password provided by front end user
   * @returns Observable
   */
  login(
    username: string,
    password: string,
    successCallback: Function,
    errorCallBack: Function
  ): void {
    this.http
      .post<JWToken>(this.backSrv.getLoginUrl(), { username, password })
      .subscribe({
        next: (response: any) => {
          this.tokenService.setAuthorizationToken(response);
          this.user$ = this.fetchUserById(this.getUserId());
          successCallback(response);
        },
        error: (error: any) => {
          errorCallBack(error);
        },
      });
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
   * Get the user id from JWT stored in localStorage
   * @returns userid string
   */
  getUserId(): string {
    return this.tokenService.getUserIdFromJWT();
  }

  registerNewUser(newUser: User, successCb: Function, errorCb: Function): void {
    this.http.post<any>(this.backSrv.getRegisterUrl(), newUser).subscribe({
      next: (response) => {
        successCb(response);
      },
      error: (error) => errorCb(error),
    });
  }

  /**
   * @returns Retrieve the current user
   */
  getCurrentUser(forceFetch: boolean = false): Observable<User> {
    if (forceFetch) {
      this.user$ = this.fetchUserById(this.getUserId());
    }
    return this.user$;
  }

  /**
   *
   * @param updatedData return updated user
   */
  updateUser(updatedData: User): Observable<User> {
    const url = `${this.backSrv.getUserUrl()}/update`;
    return this.http.patch<any>(url, updatedData).pipe(
      map((data: any) => {
        this.user$ = this.fetchUserById(this.getUserId());
        return this.user$;
      }),
      concatAll()
    );
  }

  private fetchUserById(userid: string): Observable<User> {
    const url: string = `${this.backSrv.getUserUrl_id()}/${userid}`;
    return this.http.get<User>(url);
  }
}