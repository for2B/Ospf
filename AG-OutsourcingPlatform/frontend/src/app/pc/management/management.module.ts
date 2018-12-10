import {NgModule} from '@angular/core';
import {UserManagementComponent} from './user-management/user-management.component';
import {ManagementRoutingModule} from './management-routing.module';
import {ShareModule} from '../../share/share.module';
import {UserManagementService} from '../../services/panjiawei/user-management.service';
import { UserSettingComponent } from './user-setting/user-setting.component';

@NgModule({
  imports: [
    ShareModule,
    ManagementRoutingModule
  ],
  declarations: [
    UserManagementComponent,
    UserSettingComponent
  ]
  ,
  providers: [
    // UserManagementService,
  ]
})
export class ManagementModule {
}
