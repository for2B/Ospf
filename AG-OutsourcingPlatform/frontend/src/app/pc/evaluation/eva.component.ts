import { Component, OnInit } from '@angular/core';
import { Router, provideRoutes } from '@angular/router';
import {TitleInfo} from './../../services/chencanxin/title-info';
import {TitleBarService} from './../../share/osp-infobar/titleBar.service';
import {tab, tabGroup} from './../../services/chencanxin/tab';
import {Title} from '@angular/platform-browser';


@Component({
    selector: 'app-eva',
    templateUrl: './eva.component.html',
    styleUrls: ['./eva.component.css'],
    }
)

export class EvaComponent implements OnInit {
    constructor(
      private titleBar: TitleBarService,
      private title: Title,

    ) { }
    titleinfo: TitleInfo = new TitleInfo('anticon anticon-solution' , '我的评价' , '在这里您可以看到与你有关的评价');

      tabs: tab[ ] = [
        new tab('邀请的评价'),
        new tab('对我的评价'),
      ];
      tabgroup: tabGroup = new tabGroup(this.tabs);

      ngOnInit() {
        this.title.setTitle('我的评价');
        this.titleBar.set(this.titleinfo);
        this.titleBar.setEvent.emit(this.titleinfo);
      }
}

