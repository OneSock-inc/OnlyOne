import { Injectable } from '@angular/core';
import {
  Route,
  Router,
  UrlSegment,
  UrlTree,
} from '@angular/router';
import { Observable } from 'rxjs';
import { CanMatch } from '@angular/router';
import { UserService } from '../userService/user-service.service';

// Documentation
// https://netbasal.com/introducing-the-canmatch-router-guard-in-angular-84e398046c9a
// https://levelup.gitconnected.com/route-guards-angular-1ea6e596ce65

@Injectable({
  providedIn: 'root',
})
export class AccessControlService implements CanMatch {
  constructor(private userSrv: UserService, private router: Router) {}

  canMatch(
    route: Route,
    segments: UrlSegment[]
  ):
    | boolean
    | UrlTree
    | Observable<boolean | UrlTree>
    | Promise<boolean | UrlTree> {
    if (segments[0].path === 'login') {
      return !this.userSrv.isLoggedIn() || this.router.parseUrl('/');
    } else {
      return this.userSrv.isLoggedIn() || this.router.parseUrl('/login');
    }
  }

}

@Injectable({
  providedIn: 'root',
})
export class AccessControlMatchPage implements CanMatch {
  constructor(private router: Router) { }
  
  previousUrl!: string;
  currentUrl!: string;

  canMatch(
    route: Route,
    segments: UrlSegment[]
  ):
    | boolean
    | UrlTree
    | Observable<boolean | UrlTree>
    | Promise<boolean | UrlTree> {
      if ((segments[0].path === 'match-lose' || segments[0].path === 'match-win') && this.router.url === '/sock-list') {
        return true;
      } else {
        return this.router.parseUrl('sock-list');
      }
  }

}