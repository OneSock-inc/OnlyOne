import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.scss'],
  host: {'class': 'default-layout'}
})
export class LoginPageComponent implements OnInit {

  hide = true;
  username!: string;
  password!: string;



  ngOnInit(): void {
    //this.title = 'Login';
  }

  onSubmit(form: any): void {
    console.log(form);
    alert("Connection successful");
    // send to api
    //form.username
    //form.password
  }

}
