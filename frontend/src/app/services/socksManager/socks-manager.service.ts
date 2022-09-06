import { HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, map, Observable, Subscription } from 'rxjs';
import { Sock } from 'src/app/dataModel/sock.model';
import { BackendLinkService } from '../backendservice/backend-link.service';
import { UserService } from '../userService/user-service.service';

interface PostResponse {
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

    this.userSocks = new Map();
    this.dataSource = new BehaviorSubject<ReadonlyMap<string, Sock>>(this.userSocks);
    this.retrieveSocks();

    this.matchSocks = new Map();
  }

  private dataSource: BehaviorSubject<ReadonlyMap<string, Sock>>;
  private userSocks: Map<string, Sock>;
  private matchSocks: Map<string, Sock>;

  getCurrentUserSocks(): Observable<ReadonlyMap<string, Sock>> {
    return this.dataSource.asObservable();
  }

  /**
   * Try to retrieve sock locally, make a request otherwise
   * @param sockId
   * @param successCallback
   * @param errorCallback
   */
  getSockById(
    sockId: string,
    successCallback: Function,
    errorCallback: Function
  ) {
    if (this.userSocks.has(sockId)) {
      successCallback(this.userSocks.get(sockId));
    } else {
      const url = this.backendSrv.getSockUrl() + '/' + sockId;
      this.getData(url, successCallback, errorCallback);
    }
  }

  /**
   * Make a request to backend to register a new sock.
   *
   * @param newSock
   * @param successCallback
   * @param errorCallback
   */
  registerNewSock(
    newSock: Sock,
    successCallback: Function,
    errorCallback: Function
  ): void {
    this.http
      .post<PostResponse>(this.backendSrv.postSockUrl(), this.newSockToJson(newSock))
      .subscribe({
        next: (response: PostResponse) => {
          this.getSockById(
            response.id,
            (data: Sock) => {
              successCallback(response);
              this.userSocks.set(response.id, data);
              this.dataSource.next(this.userSocks);
            },
            (e: any) => {
              console.error(
                'Seams like new sock was successfully registerd, however something went wrong.'
              );
              errorCallback(e);
            }
          );
        },
        error: (e) => errorCallback(e),
      });
  }

  private retrieveSocks(): void {
    const url: string =
      this.backendSrv.getUserUrl() + '/' +
      + this.userService.getUser().username + '/sock';
    this.getData(url, this.retrCbSuccess, (e: any) => {
      console.error(
        'Seams like new sock was successfully registered, however something went wrong.'
      );
    });
  }

  private retrCbSuccess = (data: Sock[]) => {
    const dataMap: Map<string, Sock> = new Map(
      data.map((sock: Sock) => {
        return [sock.id, sock];
      })
    );
    this.userSocks = dataMap;
    this.dataSource.next(this.userSocks);
  };

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

}
