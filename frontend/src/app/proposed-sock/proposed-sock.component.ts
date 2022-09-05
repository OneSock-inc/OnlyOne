import { Component, OnInit, Input } from '@angular/core';
import { Sock } from '../dataModel/sock.model';

@Component({
  selector: 'app-proposed-sock',
  templateUrl: './proposed-sock.component.html',
  styleUrls: ['./proposed-sock.component.scss']
})
export class ProposedSockComponent implements OnInit {

  @Input() // to be accessed by the parent component
  sock : Sock = {
    id: "",
    description: "",
    shoeSize: 0,
    color: "",
    picture:""
  };


  constructor() { }

  ngOnInit(): void {

  }

  accept(sockId : string){
    console.log("accepting sock " + sockId);
  }

  refuse(sockId : string){
    console.log("refusing sock " + sockId);
  }
}
