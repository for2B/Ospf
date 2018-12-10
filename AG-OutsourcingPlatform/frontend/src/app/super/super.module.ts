import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { SuperRoutingModule } from './super-routing.module';
import { SetGzhmenuComponent } from './set-gzhmenu/set-gzhmenu.component';
import { BuildCodeComponent } from './build-code/build-code.component';
import {BuildCodeService} from './build-code/build-code.service';
import {ShareModule} from '../share/share.module';
import { SuperComponent } from './super.component';
import { JsonToGoStructComponent } from './json-to-go-struct/json-to-go-struct.component';

@NgModule({
  imports: [
    CommonModule,
    SuperRoutingModule,
    ShareModule,
  ],
  declarations: [
    SetGzhmenuComponent,
    BuildCodeComponent,
    SuperComponent,
    JsonToGoStructComponent,
  ],
  providers:
  [
    BuildCodeService,
  ]
})
export class SuperModule { }
