import { HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Sock } from 'src/app/dataModel/sock.model';
import { BackendLinkService } from '../backendservice/backend-link.service';
import { UserService } from '../userService/user-service.service';

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
    this.matchSocks = new Map();
  }

  private userSocks: Map<string, Sock>;
  private matchSocks: Map<string, Sock>;

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
    this.postData(
      this.backendSrv.postSockUrl(),
      newSock,
      successCallback,
      errorCallback
    );
  }

  private setSocks(): void {
    const url: string = this.backendSrv.getSockUrl() + this.userService.getUser().username;
    this.getData(url, (data: Map<string, Sock>) => {
      this.userSocks = data;
    }, (error: any) => {
      console.log("Cannot retrive socks from database");
    });
  }

  private setMatches(): void {
    const url: string = this.backendSrv.getSockUrl() + this.userService.getUser().username;
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

  private postData(
    url: string,
    body: any,
    successCallback: Function,
    errorCallback: Function
  ): void {
    this.http.post<any>(this.backendSrv.postSockUrl(), body).subscribe({
      next: (response) => {
        successCallback(response);
      },
      error: (error) => errorCallback(error),
    });
  }
}
