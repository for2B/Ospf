import {NgModule} from '@angular/core';
import {ShareModule} from '../share/share.module';
import {PcRoutingModule} from './pc-routing.module';
import {PcComponent} from './pc.component';
import {TitleBarService} from '../share/osp-infobar/titleBar.service';

@NgModule({
  imports: [
    ShareModule,
    PcRoutingModule,
  ],
  declarations: [
    PcComponent,

  ]
  ,
  providers:[
    TitleBarService,

  ]
})
export class PcModule {
}
