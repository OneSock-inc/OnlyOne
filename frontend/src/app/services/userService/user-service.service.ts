import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { BackendLinkService } from '../backendservice/backend-link.service';

import { User } from 'src/app/dataModel/user.model';
import { catchError, map, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  constructor(private http: HttpClient, private backSrv: BackendLinkService) {
    this.user = UserService.userFromLocalStorage();
  }

  private user: User;
  updateUser(user : User, sucessCallb: Function, errorCallb : Function) : void {
    this.http.patch<any>(this.backSrv.getUpdateUrl(),user).subscribe({
      next: (response) => {
        sucessCallb(response)
      },
      error: (error) => errorCallb(error),
    });
  }

  registerNewUser(newUser: User, successCb: Function, errorCb: Function): void {
    UserService.registerUserInLocalStorage(newUser);
    this.http
      .post<any>(this.backSrv.getRegisterUrl(), newUser)
      .subscribe({
        next: (response) => {
          successCb(response)
        },
        error: (error) => errorCb(error),
      });
  }

  /**
   * @returns Retrieve the user if present in localStorage, return empty user otherwise.
   */
  getUser(): User {
    return this.user;
  }

  getUserV2(): Observable<User>{
    if (this.user.username === "") {
      const userName = localStorage.getItem('userName');
      const url: string = `${this.backSrv.getUserUrl()}/${userName}`;
      return this.http.get<User>(url).pipe(
        map((data: User) => {
          localStorage.setItem('fullUser', JSON.stringify(data));
          return data;
        }),
      );
    } else {
      return new Observable<User>((s) => s.next(this.user));
    }
  }

  saveUserName(username: string): void {
    localStorage.setItem('userName', username);
  }

  private static userFromLocalStorage(): User {
    const usrStr = localStorage.getItem('currentUser');
    if (typeof usrStr === 'string') {
      return JSON.parse(usrStr);
    } else {
      return new User();
    }
  }

  private static registerUserInLocalStorage(user: User): void {
    const usrStr = JSON.stringify(user);
    const usrClone = JSON.parse(usrStr);
    usrClone.password = '';
    localStorage.setItem('currentUser', JSON.stringify(usrClone));
  }
  
}
