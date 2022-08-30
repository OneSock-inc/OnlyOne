import { Component, OnInit } from '@angular/core';
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
  username!: string;
  password!: string;

  config?: Config;
  error: any;

  constructor(private configService: ConfigService) {
  }

  ngOnInit(): void {
    //this.title = 'Login';
  }

  onSubmit(form: any): void {
    console.log(this.configService.getConfig().subscribe({
      next: (data: Config) => this.config = {...data},
      error: error => this.error = error
    }));
    console.log(form);
    alert("Connection successful");
    // send to api
    //form.username
    //form.password
  }

}
