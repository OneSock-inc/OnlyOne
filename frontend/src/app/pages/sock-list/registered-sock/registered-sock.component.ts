import { Component, OnInit, Input } from '@angular/core';
import { Sock, typeToString  } from '../../../dataModel/sock.model';

@Component({
  selector: 'app-registered-sock',
  templateUrl: './registered-sock.component.html',
  styleUrls: ['./registered-sock.component.scss']
})
export class RegisteredSockComponent implements OnInit {

  @Input() // to be accessed by the parent component
  sock = new Sock;

  typeToString: (sock: Sock) => string = typeToString;

  possibleMatches !: Sock[];

  constructor() { }

  ngOnInit(): void {
    this.possibleMatches = [new Sock, new Sock, new Sock]; // TODO: replace by API call
  }

  getBadge() : string {
    if (this.sock.match !== "") {
      return "\u{1F5A4}"; // black heart
    } else {
      return this.possibleMatches.length.toString();
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
