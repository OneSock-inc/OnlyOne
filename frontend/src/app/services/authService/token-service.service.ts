import { Injectable } from '@angular/core';
import { JWToken } from 'src/app/dataModel/jwt.model';

@Injectable({
  providedIn: 'root',
})
export class TokenService {
  constructor() {}

  private token?: JWToken;

  getAuthorizationToken(): string {
    if (typeof this.token === 'undefined') {
      this.setSessionFromLocalStorage();
    }
    if (this.token?.token) {
      return this.token.token;
    } else {
      return '';
    }
  }

  setAutoriuationToken(token: JWToken): void {
    this.token = token;
    this.setSession(token);
  }

  private setSession(token: JWToken) {
    // const now = new Date('now');
    // TODO: check date validity
    localStorage.setItem('jwtoken', JSON.stringify(token));
  }

  private setSessionFromLocalStorage(): void {
    const strObj: string | null = localStorage.getItem('jwtoken');
    if (strObj !== null) {
      const jsonObj: Object = JSON.parse(strObj);
      this.token = <JWToken>jsonObj;
    }
  }
}
