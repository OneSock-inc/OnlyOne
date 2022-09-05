import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-match-lose',
  templateUrl: './match-lose.component.html',
  styleUrls: ['./match-lose.component.scss']
})
export class MatchLoseComponent implements OnInit {

  constructor() { }

  address!: string;
  sock!: string;
  otherSock!: string;

  ngOnInit(): void {
    this.address = "Big Road 38"
    // TODO: remplace base 64 by an image of height=500px
    this.sock = "iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="
    this.otherSock = "iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="
  }

}
