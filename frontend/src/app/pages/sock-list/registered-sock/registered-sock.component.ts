import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { Observable } from 'rxjs';
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
    private router: Router
  ) {}

  @Input() // to be accessed by the parent component
  sock: Sock = new Sock();

  typeToString(sock: Sock): string {
    return tts(sock);
  }

  possibleMatches!: Observable<String>;

  private redirectUrl = new Array();

  ngOnInit(): void {
    if (this.sock.id !== '') {
      if (this.sock.match !== '') {
        this.possibleMatches = new Observable<String>((s) =>
          s.next('\u{2665}')
        );
        const url = this.sock.matchResult;
        this.redirectUrl = [
          `match-${url}`,
          { queryParams: { mySock: this.sock.id, otherSock: this.sock.match } },
        ];
      } else {
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

  onSelect() {
    if (this.sock.match !== '') {
      console.log('Go to win/lose page of sock : ' + this.sock.id); // TODO api call
    } else {
      console.log('Go to possible matches of sock: ' + this.sock.id); // TODO api call
    }
  }

  onClick() {
    this.router.navigate(this.redirectUrl);
  }
}
