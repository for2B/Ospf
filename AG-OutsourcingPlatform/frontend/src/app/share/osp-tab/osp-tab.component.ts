import {Component, Input, OnChanges, OnInit} from '@angular/core';
import {forEach} from '@angular/router/src/utils/collection';
import {tab, tabGroup} from '../../services/chencanxin/tab';

@Component({
  selector: 'app-osp-tab',
  templateUrl: './osp-tab.component.html',
  styleUrls: ['./osp-tab.component.css']
})
export class OspTabComponent implements OnInit {
  constructor() {
  }
  @Input()
  tabs :tab[]= [
    new tab('初始值'),
    new tab('初始值'),
    new tab('初始值'),
  ];

  @Input()
  tabgroup: tabGroup = new tabGroup(this.tabs);
  ngOnInit() {

  }

}
