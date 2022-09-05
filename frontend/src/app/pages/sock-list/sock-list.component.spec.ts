import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SockListComponent } from './sock-list.component';

describe('SockListComponent', () => {
  let component: SockListComponent;
  let fixture: ComponentFixture<SockListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SockListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SockListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
