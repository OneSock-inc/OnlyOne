import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap, Router, TitleStrategy } from '@angular/router';
import { Sock, typeToString } from '../../dataModel/sock.model';
import { catchError, switchMap } from 'rxjs/operators';
import { map, Observable, Subscriber, throwError } from 'rxjs';
import { SocksManagerService, UserSocks } from 'src/app/services/socksManager/socks-manager.service';

@Component({
  selector: 'app-matching-socks',
  templateUrl: './matching-socks.component.html',
  styleUrls: ['./matching-socks.component.scss'],
  host: { class: 'default-layout' },
})
export class MatchingSocksComponent implements OnInit {
  typeToString: (sock: Sock) => string = typeToString;

  sock!: Observable<Sock>;

  propositionMatches!: Observable<UserSocks>;

  constructor(
    private route: ActivatedRoute,
    private sockManager: SocksManagerService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.sock = this.route.paramMap.pipe(
      switchMap((params: ParamMap) => {
        return this.sockManager.getSockById(params.get('id')!).pipe(
          map((sock: Sock) => {
            this.propositionMatches = this.sockManager.getPotencialMatches(sock.id);
            return sock;
          }),
          catchError(this.errorHandling)
        );
      })
    );
  }

  private errorHandling(e: any) {
    alert(e.message);
    return throwError(() => {
      new Error(e.message);
    });
  }

}