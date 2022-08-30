import { Component, OnInit, ViewContainerRef, ViewChild } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { AuthService } from '../authService/auth.service';
import { Router } from '@angular/router';
import { LoaderComponent } from '../loader/loader.component';
import { LoaderDirective } from '../loader/loader.directive';

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


  constructor(
    private router: Router,
    private formBuilder: FormBuilder,
    private authService: AuthService,
    private viewContainerRef: ViewContainerRef
  ) {}

  ngOnInit(): void {
    //this.title = 'Login';
  }

  onSubmit() {

    if (this.clicked) return;
    this.clicked = true;

    //console.log(this.loginForm.value);
    // this.authService.login("éajkshdg", "salékgdj").add(() => {
    //   this.router.navigate(["/home"]);
    // });

    this.createLoader();
    this.authService.login('éajkshdg', 'salékgdj').add(() => {
      if (typeof this.authService.error !== 'undefined') {
        console.warn('error');
        this.clicked = false;
      } else {
        this.router.navigate(['/home']);
      }
    });


    
    //alert("Connection successful");
    // send to api
    //form.username
    //form.password
  }

  createLoader(): void {
    this.dynamicChild.viewContainerRef.createComponent(LoaderComponent);
  }

}
