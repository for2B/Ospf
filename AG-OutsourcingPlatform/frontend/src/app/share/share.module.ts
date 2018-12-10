import {NgModule, NO_ERRORS_SCHEMA} from '@angular/core';
import {CommonModule} from '@angular/common';
import {AddIconComponent} from './add-icon/add-icon.component';
import {Page404Component} from './page-404/page-404.component';
import {FroalaEditorComponent} from './froala-editor/froala-editor.component';
import {FroalaEditorModule, FroalaViewModule} from 'angular-froala-wysiwyg';
import {RouterModule} from '@angular/router';
import {HttpClientModule} from '@angular/common/http';
import {NgZorroAntdModule, NZ_I18N, NZ_MESSAGE_CONFIG, NzMessageService, zh_CN} from 'ng-zorro-antd';
import {MockLoginComponent} from './mock-login/mock-login.component';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {OspNavbarComponent} from './osp-navbar/osp-navbar.component';
import {OspInfobarComponent} from './osp-infobar/osp-infobar.component';
import {TopLayerComponent} from './top-layer/top-layer.component';
import {OspTabComponent} from './osp-tab/osp-tab.component';
import {CookieService} from 'ngx-cookie-service';
import {MockLoginService} from './mock-login/mock-login.service';
import {OspfTabComponent} from './ospf-tab/ospf-tab.component';
import {TabItemComponent} from './ospf-tab/tab-item/tab-item.component';
import {ToastService} from './toast/toast.service';
import {ModalService} from './modal/modal.service';
import {BaseService} from '../services/chencanxin/baseService';

@NgModule({
  imports: [
    CommonModule,
    FroalaEditorModule.forRoot(),
    FroalaViewModule.forRoot(),
    RouterModule,
    HttpClientModule,
    NgZorroAntdModule,
    FormsModule,
    ReactiveFormsModule,
  ],
  declarations: [
    Page404Component,
    AddIconComponent,
    FroalaEditorComponent,
    MockLoginComponent,
    OspNavbarComponent,
    OspInfobarComponent,
    AddIconComponent,
    TopLayerComponent,
    OspTabComponent,
    OspfTabComponent,
    TabItemComponent,
  ],
  exports: [
    AddIconComponent,
    Page404Component,
    FroalaEditorComponent,
    FroalaViewModule,
    CommonModule,
    HttpClientModule,
    NgZorroAntdModule,
    FormsModule,
    ReactiveFormsModule,
    OspNavbarComponent,
    OspInfobarComponent,
    AddIconComponent,
    OspTabComponent,
    TopLayerComponent,
    OspfTabComponent,
    TabItemComponent,
  ],
  providers: [
    ToastService,
    // NzMessageService,
    MockLoginService,
    {
      provide: NZ_I18N,
      useValue: zh_CN,
    },
    ModalService,
    { provide: NZ_MESSAGE_CONFIG,
      useValue: { nzMaxStack:8 }},
    CookieService,

    BaseService,
  ],
  schemas:[
    NO_ERRORS_SCHEMA,
  ]
})
export class ShareModule {
}
