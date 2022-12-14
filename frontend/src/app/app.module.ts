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
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatSliderModule } from '@angular/material/slider';
import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { MatDividerModule } from '@angular/material/divider';
import { ColorPickerModule } from 'ngx-color-picker';
import { MatCardModule } from '@angular/material/card';
import { MatBadgeModule } from '@angular/material/badge';
import { RouterModule } from '@angular/router';
import { AppRoutingModule } from './app-routing.module';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { MessageBannerDirective } from './message-banner/mesage-banner.directive';
import { MessageBannerComponent } from './message-banner/message-banner.component';
import { LoaderComponent } from './loader/loader.component';
import { LoaderDirective } from './loader/loader.directive';

import { SignupFormComponent } from './forms/signup-form/signup-form.component';
import { AddSockFormComponent } from './forms/add-sock-form/add-sock-form.component';
import { LoginFormComponent } from './forms/login-form/login-form.component';

import { AddSockPageComponent } from './pages/add-sock-page/add-sock-page.component';
import { SockListComponent } from './pages/sock-list/sock-list.component';
import { MatchingSocksComponent } from './pages/matching-socks/matching-socks.component';
import { ProposedSockComponent } from './pages/matching-socks/proposed-sock/proposed-sock.component';
import { MyAccountComponent } from './pages/my-account/my-account.component';
import { MatchWinComponent } from './pages/match-win/match-win.component';
import { MatchLoseComponent } from './pages/match-lose/match-lose.component';
import { RegisteredSockComponent } from './pages/sock-list/registered-sock/registered-sock.component';
import { BackbuttonComponent } from './backbutton/backbutton.component';
import { MatchResultComponent } from './match-result/match-result.component';

import { ConfigService } from './services/config/config.service';
import { HttpErrorService } from './services/http-interceptors/http-error.service';
import { AuthInterceptor } from './services/http-interceptors/auth-interceptor.service';
import { TokenService } from './services/authService/token-service.service';
import { AccessControlMatchPage, AccessControlService } from './services/authService/access-control.service';
import { UserService } from './services/userService/user-service.service';
import { ArrowDownComponent } from './arrow-down/arrow-down.component';

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
    BackbuttonComponent,
    MatchResultComponent,
    ArrowDownComponent,
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
    RouterModule,
  ],
  providers: [
    ConfigService,
    TokenService,
    AccessControlService,
    AccessControlMatchPage,
    UserService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: HttpErrorService,
      multi: true,
    },
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true,
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
