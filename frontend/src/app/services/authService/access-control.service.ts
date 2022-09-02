import { Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivate,
  Route,
  Router,
  RouterStateSnapshot,
  UrlSegment,
  UrlTree,
} from '@angular/router';
import { Observable } from 'rxjs';
import { AuthService } from './auth.service';
import { CanMatch } from '@angular/router';

// https://netbasal.com/introducing-the-canmatch-router-guard-in-angular-84e398046c9a
// https://levelup.gitconnected.com/route-guards-angular-1ea6e596ce65

@Injectable({
  providedIn: 'root',
})
export class AccessControlService implements CanMatch {
  constructor(private authService: AuthService, private router: Router) {}

  canMatch(
    route: Route,
    segments: UrlSegment[]
  ):
    | boolean
    | UrlTree
    | Observable<boolean | UrlTree>
    | Promise<boolean | UrlTree> {
    if (segments[0].path === 'login') {
      return !this.authService.isLoggedIn() || this.router.parseUrl('/');
    } else {
      return this.authService.isLoggedIn() || this.router.parseUrl('/');
    }
  }
}
