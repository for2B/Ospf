import {Component, Input, OnInit} from '@angular/core';
import {tab, tabGroup} from '../../services/chencanxin/tab';

@Component({
  selector: 'app-ospf-tab',
  templateUrl: './ospf-tab.component.html',
  styleUrls: ['./ospf-tab.component.css']
})
export class OspfTabComponent implements OnInit {


  constructor() {
  }

  @Input()
  tabgroup: tabGroup = new tabGroup([
    new tab('初始值'),
    new tab('初始值'),
    new tab('初始值'),
  ]);

  ngOnInit() {

  }

}



