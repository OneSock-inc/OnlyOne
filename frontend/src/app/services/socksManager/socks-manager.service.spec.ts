import { HttpClient, HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';
import { services } from '..';

import { SocksManagerService } from './socks-manager.service';

describe('SocksManagerService', () => {
  let service: SocksManagerService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, HttpClientModule],
      providers: [services, HttpClient],
    });
    service = TestBed.inject(SocksManagerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
