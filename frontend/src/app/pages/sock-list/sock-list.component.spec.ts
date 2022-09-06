import { HttpClient, HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { services } from 'src/app/services';

import { SockListComponent } from './sock-list.component';

describe('SockListComponent', () => {
  let component: SockListComponent;
  let fixture: ComponentFixture<SockListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientModule],
      declarations: [ SockListComponent],
      providers: [services, HttpClient]
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
