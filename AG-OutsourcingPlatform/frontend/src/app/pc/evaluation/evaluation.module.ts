import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ShareModule } from '../../share/share.module';
import { EvaluationgRoutingModule } from './evaluation-routing.module';
import { EvaInvitiedComponent } from './eva-invitied/eva-invitied.component';
import { EvaForMeComponent } from './eva-forme/eva-forme.component';
import { EvaComponent } from './eva.component';


@NgModule({
  imports: [
    ShareModule,
    EvaluationgRoutingModule,
  ],
  declarations: [
    EvaComponent,
    EvaForMeComponent,
    EvaInvitiedComponent,
  ]
})
export class EvaluationModule { }
