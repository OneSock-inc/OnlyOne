import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SignupPageComponent } from './signup-page.component';

import { MatAutocomplete } from '@angular/material/autocomplete';
import { services } from '../../services';
import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { LoaderComponent } from '../../loader/loader.component';
import { LoaderDirective } from '../../loader/loader.directive';
import { LoginPageComponent } from '../login-page/login-page.component';
import { MessageBannerDirective } from '../../message-banner/mesage-banner.directive';
import { MessageBannerComponent } from '../../message-banner/message-banner.component';

// const newUser: NewUser = {
//   username: 'jaja',
//   password: 'A ver1 str0ng pa$$w0r!',
//   firstname: 'Janine',
//   surname: 'Paoli',
//   street: 'Ch. de la Rue',
//   country: 'Switzerland',
//   postalCode: '1000',
// };

describe('SignupPageComponent', () => {
  let component: SignupPageComponent;
  let fixture: ComponentFixture<SignupPageComponent>;

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
      providers: [services]
    }).compileComponents();

    fixture = TestBed.createComponent(SignupPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

});
