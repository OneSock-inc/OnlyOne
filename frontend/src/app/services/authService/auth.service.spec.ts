import { HttpClient, HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';
import { JWToken } from 'src/app/dataModel/jwt.model';
import { services } from '..';

import { AuthService } from './auth.service';
import { TokenService } from './token-service.service';

describe('AuthService', () => {
  let service: AuthService;
  let tokenService: TokenService;

  const token: JWToken = {
    token: '123456789',
    expire: new Date('now'),
  };

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, HttpClientModule],
      providers: [services, HttpClient],
    });
    service = TestBed.inject(AuthService);
    tokenService = TestBed.inject(TokenService);

  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('User must not be logged', () => {
    expect(service.isLoggedIn()).toBeFalsy();
  })

  it('User must be logged', () => {
    localStorage.setItem('jwtoken', JSON.stringify(token));
    tokenService.getAuthorizationToken();
    expect(service.isLoggedIn()).toBeTruthy();
    localStorage.clear();
  })
  
});
