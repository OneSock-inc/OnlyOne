import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { services } from '../services';
import { MatchService } from '../services/match/match-service.service';

import { MatchResultComponent } from './match-result.component';

describe('MatchResultComponent', () => {
  let component: MatchResultComponent;
  let fixture: ComponentFixture<MatchResultComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MatchResultComponent ],
      providers: [ services, MatchService ],
      imports: [ HttpClientModule ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MatchResultComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
