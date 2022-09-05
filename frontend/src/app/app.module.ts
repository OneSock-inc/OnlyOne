import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { TitleComponent } from './title/title.component';
import { ButtonComponent } from './button/button.component';
import { LoginPageComponent } from './login-page/login-page.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { HomePageComponent } from './home-page/home-page.component';
import { SignupPageComponent } from './signup-page/signup-page.component';

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


import {RouterModule} from '@angular/router';
import { AddSockPageComponent } from './add-sock-page/add-sock-page.component';
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
import { AuthenticationButtonComponent } from './authentication-button/authentication-button.component';
import { SignupFormComponent } from './forms/signup-form/signup-form.component';
import { LoginFormComponent } from './forms/login-form/login-form.component';
import { AddSockFormComponent } from './forms/add-sock-form/add-sock-form.component';
import { SockListComponent } from './sock-list/sock-list.component';
import { MatchingSocksComponent } from './matching-socks/matching-socks.component';
import { MySockComponent } from './my-sock/my-sock.component';
import { ProposedSockComponent } from './proposed-sock/proposed-sock.component';
import { MyAccountComponent } from './my-account/my-account.component';
import { MatchWinComponent } from './match-win/match-win.component';
import { MatchLoseComponent } from './match-lose/match-lose.component';

@NgModule({
  declarations: [
    LoaderComponent,
    LoaderDirective,
    MessageBannerComponent,
    MessageBannerDirective,
    HomePageComponent,
    AppComponent,
    TitleComponent,
    ButtonComponent,
    LoginPageComponent,
    PageNotFoundComponent,
    MessageBannerComponent,
    SignupPageComponent,
    AddSockPageComponent,
    AuthenticationButtonComponent,
    SignupFormComponent,
    LoginFormComponent,
    AddSockFormComponent,
    SignupFormComponent,
    SockListComponent,
    MatchingSocksComponent,
    MySockComponent,
    ProposedSockComponent,
    MyAccountComponent,
    MatchWinComponent,
    MatchLoseComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    RouterModule,
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
    MatCardModule
  ],
  providers: [
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
