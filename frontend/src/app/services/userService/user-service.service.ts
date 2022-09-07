import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BackendLinkService } from '../backendservice/backend-link.service';

import { User } from 'src/app/dataModel/user.model';
import { concatAll, map, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  constructor(private http: HttpClient, private backSrv: BackendLinkService) {
    this.userFromLocalStorage();
  }

  private user!: User;
  

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
  getUser(): User | undefined{
    return this.user;
  }

  getUserV2(force: boolean = false): Observable<User>{
    if (this.user.username === "" || force) {
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

  /**
   * 
   * @param updatedData return updated user
   */
  saveUser(updatedData: User): Observable<User> {
    const url = `${this.backSrv.getUserUrl()}/update`;
    updatedData.username = this.user.username;
    return this.http.patch<any>(url, updatedData)
    .pipe(
      map((data: any) => {
          return this.getUserV2(true)
            .pipe(
              map((data:User) => {
                UserService.registerUserInLocalStorage(data);
                return data;
              })
            );
      }),
      concatAll()
    )
  }

  private userFromLocalStorage(): void {
    const usrStr = localStorage.getItem('fullUser');
    const userN = localStorage.getItem('userName');
    if (typeof usrStr === 'string') {
      this.user = JSON.parse(usrStr);
    } else if(typeof userN === 'string') {
      this.getUserV2().subscribe(
        {
          next: (u: User) => this.user = u,
        }
      )
    } else {
      this.user = new User();
    }
  }

  private static registerUserInLocalStorage(user: User): void {
    const usrStr = JSON.stringify(user);
    const usrClone = JSON.parse(usrStr);
    usrClone.password = '';
    localStorage.setItem('fullUser', JSON.stringify(usrClone));
  }
  
}
