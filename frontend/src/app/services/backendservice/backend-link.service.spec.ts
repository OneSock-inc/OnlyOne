import { TestBed } from '@angular/core/testing';
import { services } from '..';
import { HttpClient } from '@angular/common/http';
import { HttpClientModule } from '@angular/common/http';
import { BackendLinkService } from './backend-link.service';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { identifierName } from '@angular/compiler';

describe('BackendLinkService', () => {
  let service: BackendLinkService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule,
        HttpClientModule,
      ],
      providers: [services, HttpClient]
    });
    service = TestBed.inject(BackendLinkService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('getLoginUrl should not return undefined', () => {
    expect(service.getLoginUrl()).toBeDefined();
  })

  it('getRegisterUrl should not return undefined', () => {
    expect(service.getRegisterUrl()).toBeDefined();
  })
  
});
