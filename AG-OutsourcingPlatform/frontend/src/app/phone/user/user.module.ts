import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CompeleteInfoComponent } from './compelete-info/compelete-info.component';
import { InvitationCodeComponent } from './invitation-code/invitation-code.component';
import { LoginComponent } from './login/login.component';
import { QualificationComponent } from './qualification/qualification.component';

@NgModule({
  imports: [
    CommonModule
  ],
  declarations: [CompeleteInfoComponent, InvitationCodeComponent, LoginComponent, QualificationComponent]
})
export class UserModule { }
