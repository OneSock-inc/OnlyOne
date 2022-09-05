import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { LoaderComponent } from 'src/app/loader/loader.component';
import { LoaderDirective } from 'src/app/loader/loader.directive';
import { LoginPageComponent } from 'src/app/pages/login-page/login-page.component';
import { MessageBannerDirective } from 'src/app/message-banner/mesage-banner.directive';
import { MessageBannerComponent } from 'src/app/message-banner/message-banner.component';
import { services } from 'src/app/services';

import { SignupFormComponent } from './signup-form.component';

describe('SignupFormComponent', () => {
  let component: SignupFormComponent;
  let fixture: ComponentFixture<SignupFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, HttpClientModule, ReactiveFormsModule, MatAutocompleteModule],
      declarations: [
        SignupFormComponent,
        LoginPageComponent,
        MessageBannerDirective,
        MessageBannerComponent,
        LoaderComponent,
        LoaderDirective,
      ],
      providers: [services]
    }).compileComponents();

    fixture = TestBed.createComponent(SignupFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
