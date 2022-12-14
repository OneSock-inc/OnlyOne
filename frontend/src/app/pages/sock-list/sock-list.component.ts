import { Component, OnInit } from '@angular/core';
import { SocksManagerService } from 'src/app/services/socksManager/socks-manager.service';
import { UserSocks } from 'src/app/services/socksManager/socks-manager.service';
import { catchError, map, Observable, throwError } from 'rxjs';
@Component({
  selector: 'app-sock-list',
  templateUrl: './sock-list.component.html',
  styleUrls: ['./sock-list.component.scss'],
})

export class SockListComponent implements OnInit {
  constructor(private socksManager: SocksManagerService) {}
  
  userSocks!: Observable<UserSocks>;

  ngOnInit(): void {
    this.userSocks = this.socksManager.retrieveSocks().pipe(
      catchError(this.errorHandling)
    );
  }

  private errorHandling(e: any) {
    alert(e.message);
    return throwError(() => new Error(e.message));
  }

}
