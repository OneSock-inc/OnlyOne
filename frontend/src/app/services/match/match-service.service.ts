import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, mergeMap, Observable } from 'rxjs';
import { Sock } from 'src/app/dataModel/sock.model';
import { User } from 'src/app/dataModel/user.model';
import { BackendLinkService } from '../backendservice/backend-link.service';
import { SocksManagerService } from '../socksManager/socks-manager.service';


export interface OtherInfo {
  sock: Sock;
  user: User;
}

@Injectable({
  providedIn: 'root',
})
export class MatchService {
  constructor(
    private sockManager: SocksManagerService,
    private httpClient: HttpClient,
    private bl: BackendLinkService
  ) {}

  otherSockId: string = '';
  selfSockId: string = '';

  otherSock?: Sock;
  selfSock?: Sock;

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

}
