import { TestBed } from '@angular/core/testing';

import { AuthInterceptor } from './auth-interceptor.service';
// https://ng-mocks.sudo.eu/
// https://stackblitz.com/github/help-me-mom/ng-mocks-sandbox/tree/tests?file=src/examples/TestHttpInterceptor/test.spec.ts
// -> tryed but does not work ...

describe('AuthInterceptorService', () => {
  let service: AuthInterceptor;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AuthInterceptor);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
