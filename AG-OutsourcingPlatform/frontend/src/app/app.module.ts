import {NgModule, NO_ERRORS_SCHEMA} from '@angular/core';
import {AppRoutingModule} from './app-routing.module';
import {registerLocaleData} from '@angular/common';
import {AppComponent} from './app.component';
import zh from '@angular/common/locales/zh';
import 'froala-editor/js/froala_editor.pkgd.min.js';
import {ShareModule} from './share/share.module';
import {BrowserModule} from '@angular/platform-browser';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {LoginGuard} from './guard/login.guard';
import { FormsModule } from '@angular/forms';
import {HttpInterceptorService} from './services/HttpInterceptor/http-interceptor.service';
import {HTTP_INTERCEPTORS} from '@angular/common/http';
import {PubHttpService} from './services/HttpService/pub-http.service';
import { LocalStorageModule } from 'angular-2-local-storage';

registerLocaleData(zh);

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
     FormsModule,
    ShareModule,
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    LocalStorageModule.withConfig({
      prefix: 'my-app',
      storageType: 'localStorage'
  })
  ],
  providers: [
    LoginGuard,
    PubHttpService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: HttpInterceptorService,
      multi: true
    },
  ],
  bootstrap: [AppComponent],
  schemas:[
    NO_ERRORS_SCHEMA,
  ]
})
export class AppModule {
}
