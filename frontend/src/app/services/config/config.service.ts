import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Config } from '../../dataModel/config.model';


interface Country {
  name: string;
  code:string;
}
@Injectable( {providedIn: 'root'} )
export class ConfigService {
  constructor(private http: HttpClient) { 
    this.config = {
      backendUrl: 'https://api.jsch.ch',
      passwordMinLength: 10
    }
   }

  private configUrl = 'assets/config.json';
  private config: Config;

  getConfig() {
    return this.config;
  }

  getPasswordMinLength(): number {
    return this.config.passwordMinLength;
  }



}


/*
Copyright Google LLC. All Rights Reserved.
Use of this source code is governed by an MIT-style license that
can be found in the LICENSE file at https://angular.io/license
*/