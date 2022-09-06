import { HttpClient, HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { ServiceWorkerModule } from "@angular/service-worker"
import { SwPush } from '@angular/service-worker';

import { services } from 'src/app/services';
import { PushNotificationService } from 'src/app/services/notification/push-notification.service';

import { LoginFormComponent } from './login-form.component';

describe('LoginFormComponent', () => {
  let component: LoginFormComponent;
  let fixture: ComponentFixture<LoginFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, HttpClientModule, ReactiveFormsModule, MatAutocompleteModule,ServiceWorkerModule.register('ngsw-worker.js', { enabled: false })],
      declarations: [ LoginFormComponent],
      providers: [services, HttpClient,SwPush,PushNotificationService]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LoginFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  
});
