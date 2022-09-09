import { HttpClient } from '@angular/common/http';
import { Injectable, OnInit } from '@angular/core';
import { map, mergeMap, Observable } from 'rxjs';
import { Sock } from 'src/app/dataModel/sock.model';
import { User } from 'src/app/dataModel/user.model';
import { BackendLinkService } from '../backendservice/backend-link.service';
import { SocksManagerService } from '../socksManager/socks-manager.service';

export interface OtherInfo {
  sock: Sock;
  user: User;
}

export interface MatchesInfo {
  otherSockId: string;
  selfSockId: string;
  otherSock: Sock;
  selfSock: Sock;
}

@Injectable({
  providedIn: 'root',
})
export class MatchService implements OnInit {
  constructor(
    private sockManager: SocksManagerService,
    private httpClient: HttpClient,
    private bl: BackendLinkService
  ) {}

  matchesInfos: Map<string, MatchesInfo> = new Map();

  otherSockId: string = '';
  selfSockId: string = '';

  otherSock?: Sock;
  selfSock?: Sock;

  ngOnInit(){
    this.init();
  }

  getSelfSock(): Observable<Sock> {
    return new Observable<Sock>((s) => s.next(this.selfSock));
  }

  getOtherInfos(): Observable<OtherInfo> {
    return this.sockManager.getSockById(this.otherSockId).pipe(
      mergeMap((otherSock: Sock) => {
        const url: string = `${this.bl.getUserUrl_id()}/${otherSock.owner}`;
        return this.httpClient.get<User>(url).pipe(
          map((user: User) => {
            return { user: user, sock: otherSock };
          })
        );
      })
    );
  }

  getSelfSock_id(selfSockId: string): Observable<Sock> {
    return new Observable<Sock>((s) => s.next(this.matchesInfos.get(selfSockId)?.selfSock))
  }

  getOtherInfos_id(selfSockId: string): Observable<OtherInfo> {
    const matchInfo: MatchesInfo | undefined = this.matchesInfos.get(selfSockId);
    return this.sockManager.getSockById(matchInfo?.otherSockId).pipe(
      mergeMap((otherSock: Sock) => {
        const url: string = `${this.bl.getUserUrl_id()}/${otherSock.owner}`;
        return this.httpClient.get<User>(url).pipe(
          map((user: User) => {
            return { user: user, sock: otherSock };
          })
        );
      })
    );
  }

  init() {
    this.otherSock = undefined;
    this.selfSock = undefined;
    this.otherSockId = '';
    this.selfSockId = '';
  }

}
