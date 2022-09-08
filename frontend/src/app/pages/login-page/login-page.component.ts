import { Component } from '@angular/core';
import { UserService } from 'src/app/services/userService/user-service.service';
import { AuthService } from '../../services/authService/auth.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.scss'],
  host: { class: 'default-layout' },
})
export class LoginPageComponent {

  constructor(private userSrv: UserService) { }

  notLogged(): boolean {
    return !this.userSrv.isLoggedIn();
  }
  
}
