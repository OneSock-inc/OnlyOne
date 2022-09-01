import { Component, OnInit } from '@angular/core';
import { AuthService } from '../services/authService/auth.service';
import { LinkElement } from '../dataModel/interface/link.model';

@Component({
  selector: 'app-authentication-button',
  templateUrl: './authentication-button.component.html',
  styleUrls: ['./authentication-button.component.scss']
})
export class AuthenticationButtonComponent implements OnInit {

  constructor(private authService: AuthService) { }

  signup: LinkElement = {text: 'Sigup', path: '/signup'};
  login: LinkElement = {text: 'Login', path: '/login'};
  logout: LinkElement = {text: 'Logout', path: '/logout'};

  isLoggedIn: boolean = this.authService.isLoggedIn()

  ngOnInit(): void {

  }

  onLogout() {
    this.authService.logout();
    location.reload();
  }


}
