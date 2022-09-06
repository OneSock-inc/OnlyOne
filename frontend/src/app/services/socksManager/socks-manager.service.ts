import { HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { Sock } from 'src/app/dataModel/sock.model';
import { BackendLinkService } from '../backendservice/backend-link.service';
import { UserService } from '../userService/user-service.service';

export type UserSocks = Sock[];
export interface PostResponse {
  id: string;
}
@Injectable({
  providedIn: 'root',
})
export class SocksManagerService {
  constructor(
    private http: HttpClient,
    private userService: UserService,
    private backendSrv: BackendLinkService
  ) {
    this.userSocks = new Array();
    this.retrieveSocks();

    this.potencialMatches = new Map();
  }

  private userSocks: UserSocks;
  private potencialMatches: Map<string, UserSocks>;


  /**
   * Make an http request to retrieve sock
   * @param sockId
   */
  getSockById(sockId: string): Observable<Sock> {
    const url = this.backendSrv.getSockUrl() + '/' + sockId;
    return this.http.get<Sock>(url);
  }

  /**
   * Make a request to backend to register a new sock.
   *
   * @param newSock
   */
  registerNewSock(
    newSock: Sock
  ): Observable<PostResponse> {
    return this.http
      .post<PostResponse>(this.backendSrv.postSockUrl(), this.newSockToJson(newSock))
      .pipe(
        map(
          (data: PostResponse) =>
          {
            this.getSockById(data.id).subscribe(
              (newSock: Sock) => {
                this.userSocks.push(newSock);
              }
            )
            return data;
          }
        )
      );
  }


  retrieveSocks(): Observable<UserSocks> {
    if (this.userSocks.length) {
      return new Observable<UserSocks>(
        (subscriber) => {
          subscriber.next(this.userSocks);
          subscriber.complete();
        }
      )
    } else {
      const url = this.userSocksUrl();
      return this.http.get<UserSocks>(url).pipe(
        map((data: UserSocks) => {
          if (data) {
            this.userSocks = data;
            // data.forEach((sock: Sock) => {
            //   this.getPotencialMatches(sock.id);
            // });
            return data;
          } else {
            return new Array();
          }
        })
      )
    }
  }

  getPotencialMatches(sockid: string): Observable<UserSocks> {
    if (this.potencialMatches.has(sockid)) {
      return new Observable((subscriber) => subscriber.next(this.potencialMatches.get(sockid)));
    } else {
      const url = `${this.backendSrv.getSockUrl()}/${sockid}/match`;
      return this.http.get<UserSocks>(url).pipe(
        map((data: UserSocks) => {
          this.potencialMatches.set(sockid, data);
          return data;
        })
      );
    }
  }
 
  private setMatches(): void {
    const url: string =
      this.backendSrv.getSockUrl() + this.userService.getUser().username;
  }

  private getData(
    url: string,
    successCallback: Function,
    errorCallback: Function
  ): void {
    this.http.get<any>(url).subscribe({
      next: (response) => {
        successCallback(response);
      },
      error: (error) => {
        errorCallback(error);
      },
    });
  }

  private newSockToJson(newSock: Sock): string {
    return JSON.stringify(newSock, (key, value) => {
      if (value === ''){
        return undefined;
      } else {
        return value;
      }
    });
  }

  private userSocksUrl(): string {
    return this.backendSrv.getUserUrl() + '/' +
    + this.userService.getUser().username + '/sock';
  }

}
