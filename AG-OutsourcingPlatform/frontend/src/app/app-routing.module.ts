import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {Page404Component} from './share/page-404/page-404.component';
import {MockLoginComponent} from './share/mock-login/mock-login.component';
import {LoginGuard} from './guard/login.guard';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'pc',
    pathMatch: 'full',
  },
  {
    path: 'pc',
    loadChildren: './pc/pc.module#PcModule',
    canActivate:[LoginGuard],
  },
  {
    path: 'phone',
    loadChildren: './phone/phone.module#PhoneModule'
  },
  {
    path: 'super',
    loadChildren: './super/super.module#SuperModule'
  },
  {
    path: 'mock_login',
    component:MockLoginComponent,
  },
  {
    path: '**',
    component: Page404Component
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {useHash: true})],
  exports: [RouterModule]
})
export class AppRoutingModule { }
