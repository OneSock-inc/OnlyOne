import { Component, ViewChild } from '@angular/core';
import {
  FormControl,
  FormGroup,
  Validators,
} from '@angular/forms';
import { AuthService } from '../services/authService/auth.service';
import { Router } from '@angular/router';
import { LoaderComponent } from '../loader/loader.component';
import { LoaderDirective } from '../loader/loader.directive';
import { MesageBannerDirective } from '../message-banner/mesage-banner.directive';
import { MessageBannerComponent } from '../message-banner/message-banner.component';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.scss'],
  host: { class: 'default-layout' },
})
export class LoginPageComponent {
  constructor(
    private router: Router,
    private authService: AuthService
  ) {
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

  hide = true;
  clicked = false;
  loginForm: FormGroup;
  loginFormInputs = { username: '', password: '' };

  @ViewChild(LoaderDirective, { static: true })
  dynamicChild!: LoaderDirective;

  @ViewChild(MesageBannerDirective, { static: true })
  dynamicBanner!: MesageBannerDirective;

  onSubmit() {
    if (this.loginForm.invalid) return;
    if (this.clicked) return;
    this.clicked = true;

    this.createLoader();
    this.authService.clearError();
    this.removeMessage();

    this.authService
      .login(this.loginForm.value.username, this.loginForm.value.password)
      .add(() => {
        if (typeof this.authService.getError() !== 'undefined') {
          this.clicked = false;
          this.removeLoader();          
          this.displayMessage(this.authService.getError());
        } else {
          this.router.navigate(['/home']);
        }
      });
  }

  createLoader(): void {
    this.dynamicChild.viewContainerRef.createComponent(LoaderComponent);
  }

  removeLoader(): void {
    this.dynamicChild.viewContainerRef.clear();
  }

  displayMessage(message: string) {
    const elem = this.dynamicBanner.vcref.createComponent(
      MessageBannerComponent
    );
    elem.instance.message = message;
  }
  removeMessage() {
    this.dynamicBanner.vcref.clear();
  }
}
