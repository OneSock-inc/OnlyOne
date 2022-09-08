import { Component, HostListener, OnInit } from '@angular/core';

@Component({
  selector: 'app-arrow-down',
  templateUrl: './arrow-down.component.html',
  styleUrls: ['./arrow-down.component.scss']
})
export class ArrowDownComponent {

  constructor() { }

  displayArrow: boolean = true;

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
