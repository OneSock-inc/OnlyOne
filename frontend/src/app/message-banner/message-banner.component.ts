import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-message-banner',
  templateUrl: './message-banner.component.html',
  styleUrls: ['./message-banner.component.scss']
})
export class MessageBannerComponent implements OnInit {

  message?: string;

  constructor() {}



  ngOnInit(): void {
  }

}
