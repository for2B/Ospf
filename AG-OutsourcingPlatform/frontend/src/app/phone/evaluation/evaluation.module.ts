import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PhEvaComponent } from './eva.component';
import { PhEvaInvitiedComponent } from './eva-inv/eva-inv.component'

@NgModule({
  imports: [
    CommonModule
  ],
  declarations: [
    PhEvaComponent,
    PhEvaInvitiedComponent,
  ]
})
export class EvaluationModule { }
