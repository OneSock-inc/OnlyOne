import { Injectable } from '@angular/core';

import { ConfigService } from '../config/config.service';
import { Config } from '../../dataModel/config.model';

@Injectable({
  providedIn: 'root'
})
export class BackendLinkService {

  config?: Config;
  error: any;

  constructor(private configService: ConfigService) { 
    this.configService.getConfig().subscribe({
      next: (data: Config) => this.config = { ...data }, // success path
      // error: error => this.error = error, // error path
    });
  }

  getLoginUrl(): string {
    return this.config?.backendUrl + "/user/login"
  }

  getLogoutUrl(): string {
    return this.config?.backendUrl + "/user/logout";
  }

}
