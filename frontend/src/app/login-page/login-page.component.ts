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

  constructor(private router: Router, private authService: AuthService) {
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

  @ViewChild(LoaderDirective, { static: true })
  dynamicChild!: LoaderDirective;

  hide = true;
  loginForm: FormGroup;

  private clicked = false;
  private loginFormInputs = { username: '', password: '' };


  onSubmit() {
    if (this.loginForm.invalid) return;
    if (this.clicked) return;
    this.clicked = true;

    this.createLoader();
    this.authService.clearError();
    this.removeMessage();

    // Call the method that send login request to the server
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
    const elem = document.querySelector("#errorMsg");
    if (elem) {
      elem.innerHTML = message;
      elem.classList.add('visible');
      elem.classList.remove('hidden');
    }
  }

  removeMessage() {
    const elem = document.querySelector("#errorMsg");
    if (elem) {
      elem.innerHTML = '';
      elem.classList.remove('visible');
      elem.classList.add('hidden');
    }
  }

  notLogged(): boolean {
    return !this.authService.isLoggedIn();
  }
}
