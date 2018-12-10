import { Component, OnInit } from '@angular/core';
import { Router, provideRoutes } from '@angular/router';
import {TitleInfo} from '../../../services/chencanxin/title-info';
import {TitleBarService} from '../../../share/osp-infobar/titleBar.service';
import { for_me_eva } from '../for_me.test';


@Component({
    selector: 'app-eva-forme',
    templateUrl: './eva-forme.component.html',
    styleUrls: ['./eva-forme.component.css'],
    }
)

export class EvaForMeComponent implements OnInit {
    constructor(
        // private titleBar:TitleBarService
    ) { }
    // titleinfo:TitleInfo=new TitleInfo("anticon anticon-solution","我的评价","在这里您可以看到与你有关的评价")
    evadata = for_me_eva;
    search(){
        //搜索
        console.log("search")
    }
    ngOnInit() {
    // this.titleBar.set(this.titleinfo);
    // this.titleBar.setEvent.emit(this.titleinfo);
  }
}
