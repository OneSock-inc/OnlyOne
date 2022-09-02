import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/authService/auth.service';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.scss']
})
export class LoginFormComponent implements OnInit {

  constructor(private router: Router, private authService: AuthService) { }

  hide = true;
  loginForm!: FormGroup;
  private clicked = false;
  private loginFormInputs = { username: '', password: '' };
  
  ngOnInit(): void {
    this.loginForm = new FormGroup({
      username: new FormControl(this.loginFormInputs.username, [
        Validators.required,
        Validators.minLength(4),
      ]),
      password: new FormControl(this.loginFormInputs.password, [
        Validators.required,
        Validators.minLength(8),
      ]),
    });
  }

}
