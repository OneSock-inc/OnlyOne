import { Component, OnInit } from '@angular/core';
import { map, Observable } from 'rxjs';
import { User } from 'src/app/dataModel/user.model';
import { UserService } from 'src/app/services/userService/user-service.service';
@Component({
  selector: 'app-my-account',
  templateUrl: './my-account.component.html',
  styleUrls: ['./my-account.component.scss'],
  host: {'class': 'default-layout'}
})
export class MyAccountComponent implements OnInit {

  constructor(private userService: UserService) { }

  user$!: Observable<User>;

  ngOnInit(): void {
    this.user$ = this.userService.getCurrentUser();
  }

}
