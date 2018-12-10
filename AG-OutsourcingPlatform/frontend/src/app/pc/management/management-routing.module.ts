import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {UserManagementComponent} from './user-management/user-management.component';

const managementroutes: Routes = [
  {
    path: '',
    redirectTo: 'usermanegement',
    pathMatch: 'full',
  },
  {
    path: 'usermanegement',
    component: UserManagementComponent,
  }
];

@NgModule({
  imports: [RouterModule.forChild(managementroutes)],
  exports: [RouterModule]
})
export class ManagementRoutingModule { }
