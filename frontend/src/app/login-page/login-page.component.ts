import { Component, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AuthService } from '../services/authService/auth.service';
import { Router } from '@angular/router';
import { LoaderComponent } from '../loader/loader.component';
import { LoaderDirective } from '../loader/loader.directive';

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
