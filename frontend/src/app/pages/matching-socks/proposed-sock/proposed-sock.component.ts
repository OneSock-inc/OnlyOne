import { ThisReceiver } from '@angular/compiler';
import { Component, OnInit, Input } from '@angular/core';
import { Observable } from 'rxjs';
import { SocksManagerService, PostMatch } from 'src/app/services/socksManager/socks-manager.service';
import { Sock, typeToString } from '../../../dataModel/sock.model';

@Component({
  selector: 'app-proposed-sock',
  templateUrl: './proposed-sock.component.html',
  styleUrls: ['./proposed-sock.component.scss']
})
export class ProposedSockComponent implements OnInit {

  constructor(private sockManager: SocksManagerService) { }

  @Input() // to be accessed by the parent component
  proposedSock = new Sock;

  @Input()
  parentSock = new Sock;

  message: string = "";
  
  typeToString: (sock: Sock) => string = typeToString;

  notRefusedOrAccepted$!: Observable<boolean>; 

  ngOnInit(): void {
    this.notRefusedOrAccepted$ = new Observable<boolean>((subscriber) => subscriber.next(this.checkMatch()));  
  }

  checkMatch(): boolean {
    let inAcceptedList = false;
    let inRefusedList = false;
    if(this.parentSock.acceptedList) {
      inAcceptedList = this.parentSock.acceptedList.includes(this.proposedSock.id);
    }
    if(this.parentSock.refusedList) {
      inRefusedList = this.parentSock.refusedList.includes(this.proposedSock.id);
    }
    return !inAcceptedList && !inRefusedList;
  }
  
  accept(sockId : string) {
    this.message = 'Accepted !'
    this.matchRequest(sockId, "accept");
  }

  refuse(sockId : string) {
    this.message = 'Refused !'
    this.matchRequest(sockId, "refuse");
  }

  private matchRequest(sockId: string, status: string) {
    this.sockManager.setMatch(this.parentSock.id, {otherSockId: sockId, status: status}).subscribe(
      {
        next: (data) => {
          console.log(data);
          this.notRefusedOrAccepted$ = new Observable((s) => s.next(false));
        },
        error: (e:any) => alert(e.message)
      }
    )
  }

}
