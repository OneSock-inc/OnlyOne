import { Injectable } from '@angular/core';

import { ConfigService } from '../config/config.service';

@Injectable({
  providedIn: 'root'
})
export class BackendLinkService {

  constructor(private configService: ConfigService) { }

  getLoginUrl(): string {
    return "";

  }

  getLogoutUrl(): string {
    return "";
  }

}
