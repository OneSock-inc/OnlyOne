import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisteredSockComponent } from './registered-sock.component';

describe('RegisteredSockComponent', () => {
  let component: RegisteredSockComponent;
  let fixture: ComponentFixture<RegisteredSockComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegisteredSockComponent ]
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
