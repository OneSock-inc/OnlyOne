import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { Observable } from 'rxjs';
import { MatchService, MatchesInfo } from 'src/app/services/match/match-service.service';
import {
  SocksManagerService,
  UserSocks,
} from 'src/app/services/socksManager/socks-manager.service';
import { Sock, typeToString as tts } from '../../../dataModel/sock.model';

@Component({
  selector: 'app-registered-sock',
  templateUrl: './registered-sock.component.html',
  styleUrls: ['./registered-sock.component.scss'],
})
export class RegisteredSockComponent implements OnInit {
  constructor(
    private socksManager: SocksManagerService,
    private router: Router,
    private matchSrv: MatchService
  ) {}

  @Input() // to be accessed by the parent component
  sock: Sock = new Sock();

  typeToString(sock: Sock): string {
    return tts(sock);
  }

  possibleMatches!: Observable<String>;
  badgeColor !: string;

  private redirectUrl = new Array();

  ngOnInit(): void {
    if (this.sock.id !== '') {
      if (this.sock.match !== '') {
        this.matchSrv.init();
        this.badgeColor = 'transparent';
        this.possibleMatches = new Observable<String>((s) =>
          s.next('\u{2764}')
        );
        const url = this.sock.matchResult;
        this.redirectUrl = [`match-${url}`, {id: this.sock.id}];
        const matchInfo: MatchesInfo = {
          otherSockId: this.sock.match,
          selfSockId: this.sock.id,
          otherSock: new Sock(),
          selfSock: this.sock
        }
        this.matchSrv.matchesInfos.set(this.sock.id, matchInfo)
        this.matchSrv.selfSock = this.sock;
        this.matchSrv.otherSockId = this.sock.match;
      } else {
        this.badgeColor = 'red';
        this.socksManager
          .getPotencialMatches(this.sock.id)
          .subscribe((data: UserSocks) => {
            this.possibleMatches = new Observable<String>((subscriber) => {
              subscriber.next(data.length.toString());
            });
          });
        this.redirectUrl = ['/sock', this.sock.id];
      }
    }
  }

  onClick() {
    this.router.navigate(this.redirectUrl);
  }
}
