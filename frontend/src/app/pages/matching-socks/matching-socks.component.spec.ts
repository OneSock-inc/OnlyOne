import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MatchingSocksComponent } from './matching-socks.component';

describe('MatchingSocksComponent', () => {
  let component: MatchingSocksComponent;
  let fixture: ComponentFixture<MatchingSocksComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MatchingSocksComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MatchingSocksComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
