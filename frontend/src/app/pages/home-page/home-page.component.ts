import { Component, OnInit } from '@angular/core';
import { UserService } from 'src/app/services/userService/user-service.service';
import { AuthService } from '../../services/authService/auth.service';

type LinkElement = {
  text: string;
  href: string;
};

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.scss'],
  host: {'class': 'default-layout'}
})

export class HomePageComponent {

  constructor(private userSrv: UserService) {}

  isLoggedIn: boolean = this.userSrv.isLoggedIn()

  addSock: LinkElement = {text: "Add a lonely sock", href: '/add-sock'};
  sockList: LinkElement = {text: "My socks", href: '/sock-list'};
  myAccount: LinkElement = {text: "My account", href: '/my-account'};
  signup: LinkElement = {text: 'Signup', href: '/signup'};
  login: LinkElement = {text: 'Login', href: '/login'};
  logout: LinkElement = {text: 'Logout', href: ''};

  onLogout() {
    this.userSrv.logout();
    location.assign('/');
  }
}
