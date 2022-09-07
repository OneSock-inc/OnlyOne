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
  
  typeToString: (sock: Sock) => string = typeToString;

  notMatchedObs: Observable<boolean> = new Observable<boolean>((subscriber) => subscriber.next(this.notMatched));
  private notMatched = true;

  ngOnInit(): void {

  }

  checkMatch(): boolean {
    const notAccepted = this.parentSock.acceptedList.every((value: string, index:number) => {
      value !== this.parentSock.id;
    })
    const notRefused = this.parentSock.refusedList.every((value: string, index:number) => {
      value !== this.parentSock.id;
    })
    return notAccepted && notRefused;
  }

  accept(sockId : string) {
    console.log("accepting sock " + sockId);
    this.matchRequest(sockId, "accept");
  }

  refuse(sockId : string){
    console.log("refusing sock " + sockId);
    this.matchRequest(sockId, "refuse");
  }

  private matchRequest(sockId: string, status: string) {
    this.sockManager.setMatch(this.parentSock.id, {otherSockId: sockId, status: status}).subscribe(
      {
        next: (data) => {
          console.log(data);
          this.notMatchedObs = new Observable((s) => s.next(false));
        },
        error: (e:any) => alert(e.message)
      }
    )
  }

}
