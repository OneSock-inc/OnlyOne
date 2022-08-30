import { Component, OnInit, ViewContainerRef, ViewChild } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { AuthService } from '../authService/auth.service';
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
  providers: [AuthService],
})
export class LoginPageComponent implements OnInit {

  hide = true;
  clicked = false;

  loginForm = this.formBuilder.group({
    username: '',
    password: '',
  });

  @ViewChild(LoaderDirective, { static: true })
  dynamicChild!: LoaderDirective;

  @ViewChild(MesageBannerDirective, { static: true })
  dynamicBanner!: MesageBannerDirective;

  constructor(
    private router: Router,
    private formBuilder: FormBuilder,
    private authService: AuthService,
  ) {}

  ngOnInit(): void {
  }

  onSubmit() {

    if (this.clicked) return;
    this.clicked = true;

    this.createLoader();
    this.authService.clearError();
    this.removeMessage();

    this.authService.login('éajkshdg', 'salékgdj').add(() => {
      if (typeof this.authService.getError() !== 'undefined') {
        console.warn('error');
        this.clicked = false;
        this.removeLoader();
        this.displayMessage("error");
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
    const elem = this.dynamicBanner.vcref.createComponent(MessageBannerComponent)
    elem.instance.message = message;
  }
  removeMessage() {
    this.dynamicBanner.vcref.clear();
  }

}
