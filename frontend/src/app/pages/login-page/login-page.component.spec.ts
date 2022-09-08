import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginPageComponent } from './login-page.component';

import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { services } from '../../services';
import { ReactiveFormsModule } from '@angular/forms';
import { LoaderComponent } from '../../loader/loader.component';
import { LoaderDirective } from '../../loader/loader.directive';
import { MessageBannerDirective } from '../../message-banner/mesage-banner.directive';
import { MessageBannerComponent } from '../../message-banner/message-banner.component';

describe('LoginPageComponent', () => {
  let component: LoginPageComponent;
  let fixture: ComponentFixture<LoginPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, HttpClientModule, ReactiveFormsModule],
      declarations: [
        LoginPageComponent,
        MessageBannerDirective,
        MessageBannerComponent,
        LoaderComponent,
        LoaderDirective,
      ],
      providers: services,
    }).compileComponents();

    fixture = TestBed.createComponent(LoginPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

});
