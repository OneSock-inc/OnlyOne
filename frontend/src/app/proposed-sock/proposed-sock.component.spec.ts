import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProposedSockComponent } from './proposed-sock.component';

describe('ProposedSockComponent', () => {
  let component: ProposedSockComponent;
  let fixture: ComponentFixture<ProposedSockComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProposedSockComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ProposedSockComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
