import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { PageNotFoundComponent } from './pages/page-not-found/page-not-found.component';
import { AccessControlMatchPage, AccessControlService } from './services/authService/access-control.service';
import { SignupPageComponent } from './pages/signup-page/signup-page.component';
import { AddSockPageComponent } from './pages/add-sock-page/add-sock-page.component';
import { MatchingSocksComponent } from './pages/matching-socks/matching-socks.component';
import { MyAccountComponent } from './pages/my-account/my-account.component';
import { SockListComponent } from './pages/sock-list/sock-list.component';
import { MatchWinComponent } from './pages/match-win/match-win.component';
import { MatchLoseComponent } from './pages/match-lose/match-lose.component';

// TODO : add "canMatch" to routes
const routes: Routes = [
  {
    path: 'login',
    component: LoginPageComponent,
    canMatch: [AccessControlService],
  },
  { path: 'signup', component: SignupPageComponent },
  { path: 'home', component: HomePageComponent },
  { path: 'add-sock', component: AddSockPageComponent },
  { path: 'matching-socks', component: MatchingSocksComponent },
  { path: 'my-account', component: MyAccountComponent },
  {
    path: 'sock-list',
    component: SockListComponent,
  },
  {
    path: 'sock/:id',
    component: MatchingSocksComponent,
  },
  { path: 'match-win', component: MatchWinComponent, canMatch: [AccessControlMatchPage] },
  { path: 'match-lose', component: MatchLoseComponent, canMatch: [AccessControlMatchPage] },
  { path: '', redirectTo: '/home', pathMatch: 'full' },
  { path: '**', component: PageNotFoundComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
