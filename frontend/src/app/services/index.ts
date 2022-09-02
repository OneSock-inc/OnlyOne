import { HttpErrorService } from './http-interceptors/http-error.service';
import { AuthService } from './authService/auth.service';
import { BackendLinkService } from './backendservice/backend-link.service';
import { AuthInterceptor } from './http-interceptors/auth-interceptor.service';
import { TokenService } from './authService/token-service.service';
import { ConfigService } from './config/config.service';

export const services = [
  HttpErrorService,
  ConfigService,
  AuthService,
  BackendLinkService,
  AuthInterceptor,
  TokenService,
];
