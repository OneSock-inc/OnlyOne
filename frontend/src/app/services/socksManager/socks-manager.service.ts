import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { concatAll, map, Observable } from 'rxjs';
import { Sock } from 'src/app/dataModel/sock.model';
import { User } from 'src/app/dataModel/user.model';
import { BackendLinkService } from '../backendservice/backend-link.service';
import { UserService } from '../userService/user-service.service';

export type UserSocks = Sock[];

export interface PostMatch {
  status: string;
  otherSockId: string;
}
export interface PostResponse {
  id: string;
}
export interface PatchResponse {
  message: string;
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
    this.retrieveSocks();
  }

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
  registerNewSock(newSock: Sock): Observable<PostResponse> {
    return this.http.post<PostResponse>(
      this.backendSrv.postSockUrl(),
      this.newSockToJson(newSock)
    );
  }

  retrieveSocks(): Observable<UserSocks> {
    return this.userService.getCurrentUser().pipe(
      map((data: User) => {
        return this.http.get<UserSocks>(this.userSocksUrl(data.username)).pipe(
          map((data: UserSocks) => {
            if (data) {
              return data;
            } else {
              return new Array();
            }
          })
        );
      }),
      concatAll()
    );
  }

  getPotencialMatches(sockid: string): Observable<UserSocks> {
    const url = `${this.backendSrv.getSockUrl()}/${sockid}/match`;
    return this.http.get<UserSocks>(url).pipe(
      map((data: UserSocks) => {
        if (data) {
          return data;
        } else {
          return new Array();
        }
      })
    );
  }

  setMatch(sockId: string, match: PostMatch): Observable<any> {
    const url: string = `${this.backendSrv.getSockUrl()}/${sockId}`;
    return this.http.patch<PatchResponse>(url, match);
  }

  private newSockToJson(newSock: Sock): string {
    return JSON.stringify(newSock, (key, value) => {
      if (value === '') {
        return undefined;
      } else {
        return value;
      }
    });
  }

  private userSocksUrl(username: string): string {
    const url: string = `${this.backendSrv.getUserUrl()}/${username}/sock`;
    return url;
  }
}
