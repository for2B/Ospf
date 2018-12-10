import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { QualificationComponent } from './qualification/qualification.component';
import { LoginComponent } from './login/login.component';
import { CompeleteInfoComponent } from './compelete-info/compelete-info.component';
import { InvitationCodeComponent } from './invitation-code/invitation-code.component';
import {ShareModule} from '../../share/share.module';
import {UserRoutingModule} from './user-routing.module';

@NgModule({
  imports: [
    ShareModule,
    UserRoutingModule,
  ],
  declarations: [QualificationComponent, LoginComponent, CompeleteInfoComponent, InvitationCodeComponent]
})
export class UserModule { }
