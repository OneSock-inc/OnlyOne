import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddSockFormComponent } from './add-sock-form.component';

describe('AddSockFormComponent', () => {
  let component: AddSockFormComponent;
  let fixture: ComponentFixture<AddSockFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddSockFormComponent ]
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
