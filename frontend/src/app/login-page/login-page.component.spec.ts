import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginPageComponent } from './login-page.component';

import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { services } from '../services';
import { ReactiveFormsModule } from '@angular/forms';
import { LoaderComponent } from '../loader/loader.component';
import { LoaderDirective } from '../loader/loader.directive';
import { MesageBannerDirective } from '../message-banner/mesage-banner.directive';
import { MessageBannerComponent } from '../message-banner/message-banner.component';

describe('LoginPageComponent', () => {
  let component: LoginPageComponent;
  let fixture: ComponentFixture<LoginPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, HttpClientModule, ReactiveFormsModule],
      declarations: [
        LoginPageComponent,
        MesageBannerDirective,
        MessageBannerComponent,
        LoaderComponent,
        LoaderDirective,
      ],
      providers: services,
    }).compileComponents();

    fixture = TestBed.createComponent(LoginPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should not crash on first call of removeMessage()', () => {
    component.removeMessage();
    expect(component).toBeTruthy();
  });

  /**
   * integration tests
   */

  it('Schould create loader', () => {
    expect(
      component.dynamicChild?.viewContainerRef.createComponent(LoaderComponent)
    ).toBeTruthy();
  });

  it('should display message in banner', () => {
    component.displayMessage('test message');
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('.banner')?.textContent).toContain(
      'test message'
    );
  });
});
