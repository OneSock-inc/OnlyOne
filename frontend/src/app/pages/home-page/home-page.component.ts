import { Component, OnInit } from '@angular/core';
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

export class HomePageComponent implements OnInit {

  constructor(private authService: AuthService) {}

  isLoggedIn: boolean = this.authService.isLoggedIn()

  addSock: LinkElement = {text: "Add a lonely sock", href: '/add-sock'};
  sockList: LinkElement = {text: "My socks", href: '/sock-list'};
  myAccount: LinkElement = {text: "My account", href: '/my-account'};
  signup: LinkElement = {text: 'Signup', href: '/signup'};
  login: LinkElement = {text: 'Login', href: '/login'};
  logout: LinkElement = {text: 'Logout', href: ''};

  ngOnInit(): void {
  }

  onLogout() {
    this.authService.logout();
    location.assign('/');
  }
}
