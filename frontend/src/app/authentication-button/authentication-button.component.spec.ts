import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { services } from '../services';

import { AuthenticationButtonComponent } from './authentication-button.component';

describe('AuthenticationButtonComponent', () => {
  let component: AuthenticationButtonComponent;
  let fixture: ComponentFixture<AuthenticationButtonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, HttpClientModule],
      declarations: [ AuthenticationButtonComponent ],
      providers: services
    })
    .compileComponents();

    fixture = TestBed.createComponent(AuthenticationButtonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
