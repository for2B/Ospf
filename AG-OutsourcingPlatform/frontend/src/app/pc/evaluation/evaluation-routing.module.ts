import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {Page404Component} from '../../share/page-404/page-404.component';
import {MockLoginComponent} from '../../share/mock-login/mock-login.component';
import { EvaInvitiedComponent } from './eva-invitied/eva-invitied.component';
import { EvaForMeComponent } from './eva-forme/eva-forme.component';
import { EvaComponent } from './eva.component';

const evaRoutes: Routes = [
  {
    path: '',
    redirectTo: 'eva',
    pathMatch: 'full',

  },
  {
    path: 'eva',
    component:EvaComponent,
  },
  {
    path : 'eva_invitied',
    component: EvaInvitiedComponent,
  },
  {
    path: 'eva_fot_me',
    component:EvaForMeComponent,
  },
  {
    path: '**',
    component: Page404Component
  }
];

@NgModule({
  imports: [RouterModule.forChild(evaRoutes)],
  exports: [RouterModule]
})
export class EvaluationgRoutingModule { }
