import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Config } from '../../dataModel/config.model';

@Injectable( {providedIn: 'root'} )
export class ConfigService {
  constructor(private http: HttpClient) { 
    this.config = {
      backendUrl: 'https://api.jsch.ch',
    }
  }
  
  private config: Config;

  getConfig() {
    return this.config;
  }

}