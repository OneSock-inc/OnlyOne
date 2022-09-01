import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginPageComponent } from './login-page.component';

import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { services } from '../services';

// import { AuthInterceptor } from './../services/http-interceptors/auth-interceptor.service';
// import { ConfigService } from "./../config/config.service";
// import { BackendLinkService } from "./../services/backendservice/backend-link.service";
// import { HttpErrorService } from './../services/http-interceptors/http-error.service';
// import { TokenService } from "./../services/authService/token-service.service";

describe('LoginPageComponent', () => {
  let component: LoginPageComponent;
  let fixture: ComponentFixture<LoginPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule,
                HttpClientModule,
              ],
      declarations: [ LoginPageComponent],
      providers: services
    })
    .compileComponents();

    fixture = TestBed.createComponent(LoginPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
