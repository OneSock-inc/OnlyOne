import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BackendLinkService } from '../backendservice/backend-link.service';
import { JWToken } from 'src/app/dataModel/jwt.model';
import { TokenService } from './token-service.service';
import { Subscription } from 'rxjs';

// Inspired by : https://blog.angular-university.io/angular-jwt-authentication/

/**
 * This service provide some useful methods to login, logout and know if
 * the user is logged in.
 */
@Injectable({
  providedIn: 'root',
})
export class AuthService {
  constructor(
    private http: HttpClient,
    private backendLink: BackendLinkService,
    private tokenService: TokenService
  ) {
  }



}
