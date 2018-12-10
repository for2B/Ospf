import { Component, OnInit } from '@angular/core';
import { TitleInfo } from '../../../services/chencanxin/title-info';
import { tabGroup, tab } from '../../../services/chencanxin/tab';
import { TitleBarService } from '../../../share/osp-infobar/titleBar.service';

@Component({
  selector: 'app-my-pro-list',
  templateUrl: './my-pro-list.component.html',
  styleUrls: ['./my-pro-list.component.css']
})
export class MyProListComponent implements OnInit {

  constructor(private titleBar: TitleBarService) { }
  titleinfo: TitleInfo = new TitleInfo('anticon anticon-switcher' , '我的项目' , '在这里查看与我相关的项目');
  
  
  tabs: tab[ ] = [
    new tab('我创建的'),
    new tab('我承包的'),
    new tab('我参与的')
  ];

  tabgroup: tabGroup = new tabGroup(this.tabs);

  ngOnInit() {
    this.titleBar.set(this.titleinfo);
    this.titleBar.setEvent.emit(this.titleinfo);
  }

}
