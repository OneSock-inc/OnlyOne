import { HttpClient } from '@angular/common/http';
import { Component, HostListener, ViewChild } from '@angular/core';
import { BackendLinkService } from '../../services/backendservice/backend-link.service';

@Component({
  selector: 'app-signup-page',
  templateUrl: './signup-page.component.html',
  styleUrls: ['./signup-page.component.scss'],
  host: {'class': 'default-layout'},
})
export class SignupPageComponent {
  constructor(private http: HttpClient, private backendLink: BackendLinkService) {
    
  }
  
  @ViewChild(SignupPageComponent)
  signupPageComponent!: SignupPageComponent;

  displayArrow: boolean = true;

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

