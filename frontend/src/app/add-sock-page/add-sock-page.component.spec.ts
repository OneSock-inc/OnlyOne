import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddSockPageComponent } from './add-sock-page.component';

describe('AddSockPageComponent', () => {
  let component: AddSockPageComponent;
  let fixture: ComponentFixture<AddSockPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddSockPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddSockPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
