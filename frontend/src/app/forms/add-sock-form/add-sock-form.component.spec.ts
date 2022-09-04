import { HttpClient, HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { LoaderComponent } from 'src/app/loader/loader.component';
import { LoaderDirective } from 'src/app/loader/loader.directive';
import { LoginPageComponent } from 'src/app/login-page/login-page.component';
import { MesageBannerDirective } from 'src/app/message-banner/mesage-banner.directive';
import { MessageBannerComponent } from 'src/app/message-banner/message-banner.component';
import { SignupFormComponent } from '../signup-form/signup-form.component';

import { AddSockFormComponent } from './add-sock-form.component';

describe('AddSockFormComponent', () => {
  let component: AddSockFormComponent;
  let fixture: ComponentFixture<AddSockFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, HttpClientModule, ReactiveFormsModule, MatAutocompleteModule],
      declarations: [
        SignupFormComponent,
        LoginPageComponent,
        MesageBannerDirective,
        MessageBannerComponent,
        LoaderComponent,
        LoaderDirective,
      ],
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddSockFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
