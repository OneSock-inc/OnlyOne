import { TestBed } from '@angular/core/testing';

import { MessageDisplayerService } from './message-displayer.service';

describe('MessageDisplayerService', () => {
  let service: MessageDisplayerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MessageDisplayerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
