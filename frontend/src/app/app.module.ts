import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { TitleComponent } from './title/title.component';
import { ButtonComponent } from './button/button.component';
import { LoginPageComponent } from './login-page/login-page.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { HomePageComponent } from './home-page/home-page.component';

// Material
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AppRoutingModule } from './app-routing.module';
import { RouterModule } from '@angular/router';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { ConfigService } from './config/config.service';
import { LoaderComponent } from './loader/loader.component';
import { LoaderDirective } from './loader/loader.directive';
import { MessageBannerComponent } from './message-banner/message-banner.component';
import { MesageBannerDirective } from './message-banner/mesage-banner.directive';
import { HttpErrorService } from './services/http-interceptors/http-error.service';

@NgModule({
  declarations: [
    AppComponent,
    TitleComponent,
    ButtonComponent,
    LoginPageComponent,
    PageNotFoundComponent,
    HomePageComponent,
    LoaderComponent,
    LoaderDirective,
    MessageBannerComponent,
    MesageBannerDirective,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    RouterModule,
    FormsModule,
    ReactiveFormsModule,
    BrowserAnimationsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
    HttpClientModule,
  ],
  providers: [
    ConfigService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: HttpErrorService,
      multi: true,
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
