import {AfterContentInit, Component, OnInit} from '@angular/core';
import {TitleInfo} from '../services/chencanxin/title-info';
import {TitleBarService} from '../share/osp-infobar/titleBar.service';

@Component({
  templateUrl: './pc.component.html',
  styleUrls: ['./pc.component.css']
})
export class PcComponent implements OnInit {
  constructor(private titleBar: TitleBarService) {
  }

  titleinfo: TitleInfo = new TitleInfo('anticon anticon-solution', '标题初始化', '简介初始化');

  ngOnInit() {
     this.titleBar.get()
  }

}
