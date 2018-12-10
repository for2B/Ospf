import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {Page404Component} from '../share/page-404/page-404.component';
import {PcComponent} from './pc.component';
import {LoginGuard} from '../guard/login.guard';
const pcroutes: Routes = [
  {
    path: '',
    component: PcComponent,


  children:[
      {
        path: '',
        redirectTo: 'project',
        pathMatch: 'full',
      },
      {
        path: 'project',
        loadChildren: './project/project.module#ProjectModule'
      },
      {
        path: 'evaluate',
        loadChildren: './evaluation/evaluation.module#EvaluationModule'
      },
      {
        path: 'user',
        loadChildren: './user/user.module#UserModule'
      },
      {
        path: 'management',
        loadChildren: './management/management.module#ManagementModule'
      },
      {
        path: '**',
        component: Page404Component
      }
    ]
  }

];

@NgModule({
  imports: [RouterModule.forChild(pcroutes)],
  exports: [RouterModule]
})
export class PcRoutingModule { }
