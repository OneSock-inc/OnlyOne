import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './home-page/home-page.component';
import { LoaderComponent } from './loader/loader.component';
import { LoginPageComponent } from './login-page/login-page.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { AccessControlService } from './services/authService/access-control.service';
import { SignupPageComponent } from './signup-page/signup-page.component';

const routes: Routes = [
  { path: 'login', component: LoginPageComponent, canMatch: [AccessControlService] },
  { path: 'signup', component: SignupPageComponent },
  { path: 'home', component: HomePageComponent },
  { path: '', redirectTo: '/home', pathMatch: 'full' },
  { path: '**', component: PageNotFoundComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
