import { HttpClientModule } from '@angular/common/http';
import { TestBed } from '@angular/core/testing';
import { services } from '..';

import { MatchService } from './match-service.service';

describe('MatchServiceService', () => {
  let service: MatchService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientModule],
      providers: services
    });
    service = TestBed.inject(MatchService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

});
