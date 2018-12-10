import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { SetGzhmenuComponent } from './set-gzhmenu/set-gzhmenu.component';
import {BuildCodeComponent} from './build-code/build-code.component';
import {SuperComponent} from './super.component';
import {JsonToGoStructComponent} from './json-to-go-struct/json-to-go-struct.component';

const routes: Routes = [
  {
    path: '',
    component: SuperComponent,
    children:[
      {
        path: '',
        redirectTo: 'buildCode',
        pathMatch: 'full',
      },
      {
        path: 'setGZHmenu',
        component: SetGzhmenuComponent,
      },
      {
        path: 'buildCode',
        component: BuildCodeComponent,
      },
      {
        path: 'JsonToGoStruct',
        component: JsonToGoStructComponent,
      },

    ]
  },





];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class SuperRoutingModule { }
