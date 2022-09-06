import { HttpClient, HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { services } from 'src/app/services';

import { RegisteredSockComponent } from './registered-sock.component';

describe('RegisteredSockComponent', () => {
  let component: RegisteredSockComponent;
  let fixture: ComponentFixture<RegisteredSockComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientModule],
      declarations: [ RegisteredSockComponent ],
      providers: [services, HttpClient]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RegisteredSockComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
