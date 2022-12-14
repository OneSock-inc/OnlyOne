import { Injectable } from '@angular/core';

import { ConfigService } from '../config/config.service';
import { Config } from '../../dataModel/config.model';

@Injectable({
  providedIn: 'root'
})
export class BackendLinkService {
  
  
  constructor(private configService: ConfigService) { 
    this.config = this.configService.getConfig();
  }
  private config: Config;
  
  getLoginUrl(): string {
    return this.config.backendUrl + "/user/login"
  }
  
  getRegisterUrl(): string {
    return this.config.backendUrl + "/user/register";
  }

  getUserUrl():string {
    return this.config.backendUrl + "/user"
  }

  getUserUrl_id(): string {
    return this.config.backendUrl + "/userid"
  }

  getUpdateUrl(): string {
    return this.config.backendUrl+ "/user/update"
  }
  
  getSockUrl() {
    return this.config.backendUrl+ "/sock";
  }

  postSockUrl(): string {
    return this.config.backendUrl+ "/sock";
  }
  
}
