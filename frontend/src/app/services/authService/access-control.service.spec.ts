import { TestBed } from '@angular/core/testing';

import { HttpClient, HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';


import { AccessControlService } from './access-control.service';
import { services } from '..';

describe('AccessControlService', () => {
  let service: AccessControlService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, HttpClientModule],
      providers: [services, HttpClient],
    });
    service = TestBed.inject(AccessControlService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
