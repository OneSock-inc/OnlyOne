import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { BackendLinkService } from '../backendservice/backend-link.service';

import { User } from 'src/app/dataModel/user.model';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  constructor(private http: HttpClient, private backSrv: BackendLinkService) {
    this.user = UserService.userFromLocalStorage();
  }

  private user: User;

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
