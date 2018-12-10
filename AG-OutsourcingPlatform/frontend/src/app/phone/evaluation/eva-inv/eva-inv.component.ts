import { Component, OnInit } from '@angular/core';
import { Router, provideRoutes } from '@angular/router';
import { TitleInfo } from '../../../services/chencanxin/title-info';
import { TitleBarService } from '../../../share/osp-infobar/titleBar.service';
import { testeva } from '../../../pc/evaluation/testdata';
// import { Card } from 'antd';
// const { Meta } = Card;


@Component({
    selector: 'app-eva-inv-ph',
    templateUrl: './eva-inv.component.html',
    styleUrls: ['./eva-inv.component.css'],
    }
)

export class PhEvaInvitiedComponent implements OnInit {
    // constructor(
    //     private titleBar: TitleBarService ,
    // ) {}
    // // titleinfo:TitleInfo=new TitleInfo("anticon anticon-solution","我的评价","在这里您可以看到与你有关的评价")
    // evadata = testeva;
    // im_eva(){
    //     //立即评价
    //     console.log("ima eva")
    // }
    // look_him(){
    //     //查看个人
    //     console.log("look himself")
    // }
    // // search(){
    // //     console.log("search")
    // // }
    ngOnInit() {
    // this.titleBar.set(this.titleinfo);
    // this.titleBar.setEvent.emit(this.titleinfo);
  }
}
