import { Component, OnInit, ViewChild } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/authService/auth.service';
import { MessageBannerDirective } from 'src/app/message-banner/mesage-banner.directive';
import { LoaderDirective } from 'src/app/loader/loader.directive';
import { LoaderComponent } from 'src/app/loader/loader.component';
@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.scss']
})
export class LoginFormComponent implements OnInit {

  constructor(private router: Router, private authService: AuthService) { }

  @ViewChild(MessageBannerDirective, {static: true})
  messageBanner!: MessageBannerDirective;

  // @ViewChild(LoaderDirective, {static: true})
  // loader!: LoaderDirective;


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
        Validators.minLength(10),
      ]),
    });
  }

  onSubmit(form: FormGroup): void {
    if (this.loginForm.invalid) return;
    if (this.clicked) return;
    this.clicked = true;

    //this.createLoader();
    this.authService.clearError();
    //this.removeMessage();
    this.messageBanner.hideMessage();

    // Call the method that send login request to the server
    this.authService
      .login(this.loginForm.value.username, this.loginForm.value.password)
      .add(() => {
        if (typeof this.authService.getError() !== 'undefined') {
          this.clicked = false;
          //this.removeLoader();
          this.messageBanner.displayMessage(this.authService.getError());
        } else {
          this.router.navigate(['/home']);
        }
      });
  }

  // createLoader(): void {
  //   this.loader.viewContainerRef.createComponent(LoaderComponent);
  // }

  // removeLoader(): void {
  //   this.loader.viewContainerRef.clear();
  // }

}
