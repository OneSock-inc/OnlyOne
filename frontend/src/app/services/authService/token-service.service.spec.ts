import { HttpClient, HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';
import { JWToken } from 'src/app/dataModel/jwt.model';
import { services } from '..';
import { AuthService } from './auth.service';

import { TokenService } from './token-service.service';

const token: JWToken = {
  token: '123456789',
  expire: new Date('now'),
};

describe('TokenServiceService', () => {
  let tokenService: TokenService;
  let authService: AuthService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, HttpClientModule],
      providers: [services, HttpClient],
    });
    tokenService = TestBed.inject(TokenService);
    authService = TestBed.inject(AuthService);
  });

  it('should be created', () => {
    expect(tokenService).toBeTruthy();
  });

  it('should retrive token from localStorage', () => {
    localStorage.setItem('jwtoken', JSON.stringify(token));
    expect(tokenService.getAuthorizationToken()).toEqual('123456789');
    localStorage.clear();
  });

});
