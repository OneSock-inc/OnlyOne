import { HttpErrorService } from './http-interceptors/http-error.service';
import { BackendLinkService } from './backendservice/backend-link.service';
import { AuthInterceptor } from './http-interceptors/auth-interceptor.service';
import { TokenService } from './authService/token-service.service';
import { ConfigService } from './config/config.service';
import { SocksManagerService } from './socksManager/socks-manager.service';

export const services = [
  HttpErrorService,
  ConfigService,
  BackendLinkService,
  AuthInterceptor,
  TokenService,
  SocksManagerService
];
