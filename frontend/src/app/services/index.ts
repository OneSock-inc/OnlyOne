import { HttpErrorService } from './http-interceptors/http-error.service';
import { BackendLinkService } from './backendservice/backend-link.service';
import { AuthInterceptor } from './http-interceptors/auth-interceptor.service';
import { TokenService } from './authService/token-service.service';
import { ConfigService } from './config/config.service';
import { SocksManagerService } from './socksManager/socks-manager.service';
import { MatchService } from './match/match-service.service';
import { AccessControlMatchPage, AccessControlService } from './authService/access-control.service';
import { UserService } from './userService/user-service.service';

export const services = [
  AccessControlService,
  TokenService,
  BackendLinkService,
  ConfigService,
  HttpErrorService,
  AuthInterceptor,
  SocksManagerService,
  MatchService,
  UserService,
  AccessControlMatchPage,
];
