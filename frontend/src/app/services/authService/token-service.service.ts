import { Injectable } from '@angular/core';
import { JWToken } from 'src/app/dataModel/jwt.model';
import {decodeJwt} from 'jose'


@Injectable({
  providedIn: 'root',
})
export class TokenService {
  constructor() {}

  private token?: JWToken;

  /**
   * Return the jwt in string format and non parsed.
   * @returns empty string if user there is no token in localStorage (i.e. user is not logged in).
   */
  getAuthorizationToken(): string {
    if (typeof this.token === 'undefined') {
      this.sessionFromLocalStorage();
    }
    if (this.token?.token) {
      return this.token.token;
    } else {
      return '';
    }
  }

  setAuthorizationToken(token: JWToken): void {
    this.token = token;
    this.setSession(token);
  }

  clearToken() {
    this.token = undefined;
  }

  private setSession(token: JWToken) {
    // const now = new Date('now');
    // TODO: check date validity
    localStorage.setItem('jwtoken', JSON.stringify(token));
  }

  private sessionFromLocalStorage(): void {
    const strObj: string | null = localStorage.getItem('jwtoken');
    if (strObj !== null) {
      const jsonObj: Object = JSON.parse(strObj);
      this.token = <JWToken>jsonObj;
    }
  }

  getUserIdFromJWT(): string {
    const tmp: any = decodeJwt(this.getAuthorizationToken());
    return tmp.identity;
  }

}
