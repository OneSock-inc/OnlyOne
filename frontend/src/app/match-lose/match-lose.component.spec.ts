import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MatchLoseComponent } from './match-lose.component';

describe('MatchLoseComponent', () => {
  let component: MatchLoseComponent;
  let fixture: ComponentFixture<MatchLoseComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MatchLoseComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MatchLoseComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
