import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { PageNotFoundComponent } from './pages/page-not-found/page-not-found.component';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { SignupPageComponent } from './pages/signup-page/signup-page.component';

// Material
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from "@angular/material/form-field";
import {MatButtonModule} from '@angular/material/button'; 
import {MatIconModule} from '@angular/material/icon';
import {MatAutocompleteModule} from '@angular/material/autocomplete';
import {MatSliderModule} from '@angular/material/slider';
import {MatButtonToggleModule} from '@angular/material/button-toggle';
import {MatDividerModule} from '@angular/material/divider';
import { ColorPickerModule } from 'ngx-color-picker';
import {MatCardModule} from '@angular/material/card';
import {MatBadgeModule} from '@angular/material/badge';


import {RouterModule} from '@angular/router';
import { AddSockPageComponent } from './pages/add-sock-page/add-sock-page.component';
import { AppRoutingModule } from './app-routing.module';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { ConfigService } from './services/config/config.service';
import { LoaderComponent } from './loader/loader.component';
import { LoaderDirective } from './loader/loader.directive';
import { MessageBannerComponent } from './message-banner/message-banner.component';
import { MessageBannerDirective } from './message-banner/mesage-banner.directive';
import { HttpErrorService } from './services/http-interceptors/http-error.service';
import { AuthService } from './services/authService/auth.service';
import { AuthInterceptor } from './services/http-interceptors/auth-interceptor.service';
import { TokenService } from './services/authService/token-service.service';
import { AccessControlService } from './services/authService/access-control.service';
import { SignupFormComponent } from './forms/signup-form/signup-form.component';
import { LoginFormComponent } from './forms/login-form/login-form.component';
import { AddSockFormComponent } from './forms/add-sock-form/add-sock-form.component';
import { SockListComponent } from './pages/sock-list/sock-list.component';
import { MatchingSocksComponent } from './pages/matching-socks/matching-socks.component';
import { ProposedSockComponent } from './pages/matching-socks/proposed-sock/proposed-sock.component';
import { MyAccountComponent } from './pages/my-account/my-account.component';
import { MatchWinComponent } from './pages/match-win/match-win.component';
import { MatchLoseComponent } from './pages/match-lose/match-lose.component';
import { RegisteredSockComponent } from './pages/sock-list/registered-sock/registered-sock.component';
import { ServiceWorkerModule } from '@angular/service-worker';
import { environment } from '../environments/environment';
import {PushNotificationService} from './services/notification/push-notification.service';

@NgModule({
  declarations: [
    LoaderComponent,
    LoaderDirective,
    MessageBannerComponent,
    MessageBannerDirective,
    HomePageComponent,
    AppComponent,
    LoginPageComponent,
    PageNotFoundComponent,
    MessageBannerComponent,
    SignupPageComponent,
    AddSockPageComponent,
    SignupFormComponent,
    LoginFormComponent,
    AddSockFormComponent,
    SignupFormComponent,
    SockListComponent,
    MatchingSocksComponent,
    ProposedSockComponent,
    MyAccountComponent,
    MatchWinComponent,
    MatchLoseComponent,
    RegisteredSockComponent,
  ],
  imports: [

    BrowserModule,
    AppRoutingModule,
    FormsModule,
    BrowserAnimationsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
    HttpClientModule,
    MatAutocompleteModule,
    ReactiveFormsModule,
    MatSliderModule,
    MatButtonToggleModule,
    ColorPickerModule,
    MatDividerModule,
    MatCardModule,
    MatBadgeModule,
    ServiceWorkerModule.register('ngsw-worker.js', {
      enabled: environment.production ,      // Register the ServiceWorker as soon as the application is stable
      // or after 30 seconds (whichever comes first).
      registrationStrategy: 'registerWhenStable:30000'
    }),
    RouterModule,
  ],
  providers: [
    PushNotificationService,
    ConfigService,
    AuthService,
    TokenService,
    AccessControlService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: HttpErrorService,
      multi: true,
    },
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
