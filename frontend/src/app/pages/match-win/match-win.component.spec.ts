import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MatchWinComponent } from './match-win.component';

describe('MatchWinComponent', () => {
  let component: MatchWinComponent;
  let fixture: ComponentFixture<MatchWinComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MatchWinComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MatchWinComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
