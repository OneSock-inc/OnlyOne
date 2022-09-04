import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-my-account',
  templateUrl: './my-account.component.html',
  styleUrls: ['./my-account.component.scss'],
  host: {'class': 'default-layout'}
})
export class MyAccountComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
