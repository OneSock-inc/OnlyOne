import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { BackendLinkService } from '../backendservice/backend-link.service';

import { User } from 'src/app/dataModel/user.model';
import { Subscription } from 'rxjs';

interface Response {
  message: string;
}

@Injectable({
  providedIn: 'root',
})
export class UserService {
  constructor(private http: HttpClient, private backSrv: BackendLinkService) {
    this.user = UserService.userFromLocalStorage();
  }

  private user: User;
  private error: any;
  private successResponse: any;

  registerNewUser(newUser: User): Subscription {
    UserService.registerUserInLocalStorage(newUser);
    return this.http
      .post<any>(this.backSrv.getRegisterUrl(), newUser)
      .subscribe({
        next: (response) => {
          this.successResponse = response;
          console.log(response);
        },
        error: (error) => (this.error = error),
      });
  }

  getLastSuccessResponse() {
    return this.successResponse;
  }

  getLastErrorResponse() {
    return this.error;
  }

  clearMessages(){
    this.error = undefined;
    this.successResponse = undefined;
  }

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
    user.password = '';
    localStorage.setItem('currentUser', JSON.stringify(user));
  }
  
}
