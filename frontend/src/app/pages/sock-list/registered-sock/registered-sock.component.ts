
import { Component, Input, OnInit } from '@angular/core';

import { Observable } from 'rxjs';
import { SocksManagerService, UserSocks } from 'src/app/services/socksManager/socks-manager.service';
import { Sock, typeToString as tts } from '../../../dataModel/sock.model';

@Component({
  selector: 'app-registered-sock',
  templateUrl: './registered-sock.component.html',
  styleUrls: ['./registered-sock.component.scss']
})
export class RegisteredSockComponent implements OnInit{
  
  constructor(private socksManager: SocksManagerService) { }

  @Input() // to be accessed by the parent component
  sock: Sock = new Sock();

  typeToString(sock: Sock): string {
    return tts(sock);
  }

  possibleMatches!: Observable<String>;
  
  ngOnInit(): void {
    if (this.sock.id !== "") {
      this.socksManager.getPotencialMatches(this.sock.id).subscribe(
        (data: UserSocks) => {
          this.possibleMatches = new Observable<String>((subscriber) => {
            if (data.length > 0) {
              subscriber.next(data.length.toString());
            } else {
              subscriber.next("\u{1F5A4}");
            }
            subscriber.complete();

          })
        }
      );
    }
  }

  onSelect() {
    if (this.sock.match !== "") {
      console.log("Go to win/lose page of sock : " + this.sock.id); // TODO api call
    } else {
      console.log("Go to possible matches of sock: " + this.sock.id); // TODO api call
    }
  }
}
