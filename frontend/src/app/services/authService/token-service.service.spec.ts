import { HttpClient, HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';
import { JWToken } from 'src/app/dataModel/jwt.model';
import { services } from '..';

import { TokenService } from './token-service.service';

const token: JWToken = {
  token: '123456789',
  expire: new Date('now'),
};

describe('TokenServiceService', () => {
  let tokenService: TokenService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, HttpClientModule],
      providers: [services, HttpClient],
    });
    tokenService = TestBed.inject(TokenService);
  });

  it('should be created', () => {
    expect(tokenService).toBeTruthy();
  });

  it('should retrieve token from localStorage', () => {
    localStorage.setItem('jwtoken', JSON.stringify(token));
    expect(tokenService.getAuthorizationToken()).toEqual('123456789');
    localStorage.clear();
  });

});
