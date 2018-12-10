import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {QualificationComponent} from 'src/app/pc/user/qualification/qualification.component';
import {Page404Component} from '../../share/page-404/page-404.component';
import {LoginComponent} from './login/login.component';
import {InvitationCodeComponent} from './invitation-code/invitation-code.component';
import {CompeleteInfoComponent} from './compelete-info/compelete-info.component';
import {LoginGuard} from '../../guard/login.guard';

const userRoutes: Routes = [
  {
    path: '',
    redirectTo: 'qualification',
    pathMatch: 'full',
  },
  {
    path: 'qualification',
    component:QualificationComponent,
  },
  {
    path: 'invitation_code',
    component:InvitationCodeComponent,
  },
  {
    path: 'login',
    component:LoginComponent,
  },
  {
    path: 'complete_info',
    component:CompeleteInfoComponent,
  },
  {
    path: '**',
    component: Page404Component
  }
];

@NgModule({
  imports: [RouterModule.forChild(userRoutes)],
  exports: [RouterModule]
})
export class UserRoutingModule { }