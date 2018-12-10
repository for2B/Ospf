import {Component, OnInit} from '@angular/core';
import {tab, tabGroup} from '../../../services/chencanxin/tab';
import {Title} from '@angular/platform-browser';
import {TitleInfo} from '../../../services/chencanxin/title-info';
import {TitleBarService} from '../../../share/osp-infobar/titleBar.service';
import {PubHttpService} from '../../../services/HttpService/pub-http.service';
import {ToastService} from '../../../share/toast/toast.service';
import {InvitationCodeService} from './invitation-code.service';
import {NORMAL_QRCODE_TYPE, REFEREE_QRCODE_TYPE} from '../../../const/const_chencanxin';
import {ModalService} from '../../../share/modal/modal.service';

@Component({
  selector: 'app-invitation-code',
  templateUrl: './invitation-code.component.html',
  styleUrls: ['./invitation-code.component.css']
})
export class InvitationCodeComponent implements OnInit {


  constructor(
    private title: Title,
    private titleBar: TitleBarService,
    private http: PubHttpService,
    private toast: ToastService,
    private invitationService: InvitationCodeService,
    private modal:ModalService,
  ) {
  }

  qrcodeForNormalUser = '';
  qrcodeForReferee = '';

  titleinfo: TitleInfo = new TitleInfo(
    'anticon anticon-smile',
    '我的邀请码',
    '这里您可以分享您的二维码让更多的人加入吧'
  );

  tabs:tab[] = [
    new tab('邀请伙伴'),
    new tab('邀请证明人')
  ];

  tabgroup: tabGroup = new tabGroup(this.tabs);

  ngOnInit() {
    this.title.setTitle('我的邀请码');
    this.titleBar.set(this.titleinfo);

    this.invitationService.getQrcode(
      NORMAL_QRCODE_TYPE,
      (value)=>{
      this.qrcodeForNormalUser=value;
    })

    this.invitationService.getQrcode(
      REFEREE_QRCODE_TYPE,
      (value)=>{
        this.qrcodeForReferee=value;
      })
  }



}
