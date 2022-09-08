import { Component, OnInit, ViewChild } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/authService/auth.service';
import { MessageBannerDirective } from 'src/app/message-banner/mesage-banner.directive';
import { UserService } from 'src/app/services/userService/user-service.service';
import { LoaderDirective } from 'src/app/loader/loader.directive';
import { LoaderComponent } from 'src/app/loader/loader.component';
@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.scss']
})
export class LoginFormComponent implements OnInit {

  constructor(private router: Router, private authService: AuthService, private userService: UserService) { }

  @ViewChild(MessageBannerDirective, {static: true})
  messageBanner!: MessageBannerDirective;


  hide = true;
  loginForm!: FormGroup;
  private clicked = false;
  private loginFormInputs = { username: '', password: '' };
  
  ngOnInit(): void {
    this.loginForm = new FormGroup({
      username: new FormControl(this.loginFormInputs.username, [
        Validators.required,
        Validators.minLength(3),
      ]),
      password: new FormControl(this.loginFormInputs.password, [
        Validators.required,
        Validators.minLength(10),
      ]),
    });
  }

  onSubmit(form: FormGroup): void {
    if (form.invalid) return;
    if (this.clicked) return;
    this.clicked = true;

    this.messageBanner.hideMessage();
    const userName = this.loginForm.value.username;
    const pwd = form.value.password;
    this.userService.login(userName, pwd,
      (response: any) => {
        this.router.navigate(['/home']);
      },
      (error: any) => {
        this.clicked = false;
        this.messageBanner.displayMessage(error)
      }
      )
  }

}


