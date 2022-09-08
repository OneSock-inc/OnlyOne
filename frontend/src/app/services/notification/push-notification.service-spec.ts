import { TestBed } from '@angular/core/testing';
import { ServiceWorkerModule, SwPush } from '@angular/service-worker';
import {services} from "..";
import { PushNotificationService } from './push-notification.service';

describe('PushNotificationService', () => {
  let service: PushNotificationService;

  beforeEach(() => {
    TestBed.configureTestingModule({
        imports : [ServiceWorkerModule.register('ngsw-worker.js', { enabled: false })],
      providers : [services,SwPush, PushNotificationService]
    });
    service = TestBed.inject(PushNotificationService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
