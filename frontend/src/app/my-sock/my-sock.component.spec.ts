import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MySockComponent } from './my-sock.component';

describe('MySockComponent', () => {
  let component: MySockComponent;
  let fixture: ComponentFixture<MySockComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MySockComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MySockComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
