import { Component } from '@angular/core';
import { AuthService } from '../../services/authService/auth.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.scss'],
  host: { class: 'default-layout' },
})
export class LoginPageComponent {

  constructor(private authService: AuthService) { }

  notLogged(): boolean {
    return !this.authService.isLoggedIn();
  }
}
