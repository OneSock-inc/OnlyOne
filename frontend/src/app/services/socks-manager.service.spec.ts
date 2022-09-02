import { TestBed } from '@angular/core/testing';

import { SocksManagerService } from './socks-manager.service';

describe('SocksManagerService', () => {
  let service: SocksManagerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SocksManagerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
