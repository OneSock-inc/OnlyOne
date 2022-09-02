import { Component, OnInit, HostListener } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import {DataUrl, NgxImageCompressService} from "ngx-image-compress";


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

  

  // display down arrow if the user has not scrolled to the bottom of the page
  @HostListener('window:scroll', ['$event'])
  onScroll(event: Event): void {
    if (window.pageYOffset >= (document.documentElement.scrollHeight - document.documentElement.clientHeight)) {
      this.displayArrow = false;
    }
    else {
      this.displayArrow = true;
    }
  }

}
