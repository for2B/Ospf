import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {QualificationComponent} from 'src/app/pc/user/qualification/qualification.component';
import {Page404Component} from '../../share/page-404/page-404.component';
import { AddMemmerComponent } from './add-memmer/add-memmer.component';
import { PlatformProListComponent } from './platform-pro-list/platform-pro-list.component';
import { MyProListComponent } from './my-pro-list/my-pro-list.component';
import { PublishProjectComponent } from './publish-project/publish-project.component';
import { ProjectDetailComponent } from './project-detail/project-detail.component';
import { AmendprojectComponent } from './amendproject/amendproject.component';

const projectRoutes: Routes = [
  {
    path: '',
    redirectTo: 'pfprolist',
    pathMatch: 'full',
  },
  {
    path: 'addmember',
    component:AddMemmerComponent,
  },
  {
    path: 'pfprolist',
    component:PlatformProListComponent,
  },
  {
    path:'myprolist',
    component:MyProListComponent,
  },
  {
    path:'pubproject',
    component:PublishProjectComponent
  },
  {
    path:'detail',
    component:ProjectDetailComponent
  },
  {
    path:'amendproject',
    component:AmendprojectComponent
  },
  {
    path: '**',
    component: Page404Component
  }
];

@NgModule({
  imports: [RouterModule.forChild(projectRoutes)],
  exports: [RouterModule]
})
export class ProjectRoutingModule { }
