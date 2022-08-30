import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';

import { ConfigService } from '../config/config.service';
import { Config } from '../dataModel/config.model';
@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.scss'],
  host: {'class': 'default-layout'},
  providers: [ConfigService]
})
export class LoginPageComponent implements OnInit {

  hide = true;
 
  loginForm = this.formBuilder.group({
    username: '',
    password: ''
  });

  constructor(private formBuilder: FormBuilder){

  }

  ngOnInit(): void {
    //this.title = 'Login';
  }

  onSubmit(): void {
    console.log(this.loginForm.value);
    //alert("Connection successful");
    // send to api
    //form.username
    //form.password
  }

}
