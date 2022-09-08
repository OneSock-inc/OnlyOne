import { Component, OnInit, HostListener } from '@angular/core';

@Component({
  selector: 'app-add-sock-page',
  templateUrl: './add-sock-page.component.html',
  styleUrls: ['./add-sock-page.component.scss'],
  host: {'class': 'default-layout'}
})
export class AddSockPageComponent implements OnInit {

  displayArrow: boolean = true;

  constructor() { }

  ngOnInit(): void {
  }

}
