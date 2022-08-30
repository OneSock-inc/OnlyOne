import { TestBed } from '@angular/core/testing';

import { BackendLinkService } from './backend-link.service';

describe('BackendLinkService', () => {
  let service: BackendLinkService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(BackendLinkService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
