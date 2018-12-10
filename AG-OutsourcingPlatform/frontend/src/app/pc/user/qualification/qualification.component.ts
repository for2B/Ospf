import {Component, OnInit} from '@angular/core';
import {TitleInfo} from '../../../services/chencanxin/title-info';
import {TitleBarService} from '../../../share/osp-infobar/titleBar.service';
import {tab, tabGroup} from '../../../services/chencanxin/tab';
import {Title} from '@angular/platform-browser';
import {ActivatedRoute, Params} from '@angular/router';
import {ToastService} from '../../../share/toast/toast.service';
import {Member} from '../../../services/panjiawei/users';

@Component({
  selector: 'app-qualification',
  templateUrl: './qualification.component.html',
  styleUrls: ['./qualification.component.css']
})
export class QualificationComponent implements OnInit {


  constructor(
    private title: Title,
    private titleBar: TitleBarService,
    private toast:ToastService,
  ) { }
   
  titleinfo: TitleInfo = new TitleInfo('anticon anticon-smile', '个人能力资质', '这里您可以查看/更新个人资料以及能力资质证明');


  tabgroup: tabGroup = new tabGroup([
    new tab('基础信息').setSelector('.tab1'),
    new tab('擅长领域').setSelector('.tab2'),
    new tab('证明资料').setSelector('.tab3'),
  ]);

  ngOnInit() {
    this.title.setTitle('个人能力资质');
    this.titleBar.set(this.titleinfo);
  }


}
