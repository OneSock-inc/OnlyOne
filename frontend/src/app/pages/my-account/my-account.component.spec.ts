import { HttpClient, HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { services } from 'src/app/services';

import { MyAccountComponent } from './my-account.component';

describe('MyAccountComponent', () => {
  let component: MyAccountComponent;
  let fixture: ComponentFixture<MyAccountComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientModule],
      declarations: [ MyAccountComponent ],
      providers: [services, HttpClient]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MyAccountComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
