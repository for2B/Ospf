import { Component, OnInit } from '@angular/core';
import {UserManagementService} from '../../../services/panjiawei/user-management.service';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';
import {UserFullDetail} from '../../../services/panjiawei/users';
import {TitleInfo} from '../../../services/chencanxin/title-info';
import {TitleBarService} from '../../../share/osp-infobar/titleBar.service';
import {NzMessageService, NzModalService} from 'ng-zorro-antd';
import {Title} from '@angular/platform-browser';

@Component({
  selector: 'app-user-setting',
  templateUrl: './user-setting.component.html',
  // styleUrls: ['./user-setting.component.css']
})
export class UserSettingComponent implements OnInit {
  DATA = ['没有数据'];
  isSpinning = true;
  titleinfo: TitleInfo = new TitleInfo('anticon anticon-setting', '用户账号设置', '查看和配置用户的权限。');
  uid = 0;
  switchValue: boolean[] = [false, false, false, false];
  constructor(
    private titleBar: TitleBarService,
    private route: ActivatedRoute,
    private router: Router,
    private title: Title,
    private modalService: NzModalService,
    private message: NzMessageService,
    private managerService: UserManagementService ) { }
  theUser: UserFullDetail;
  ngOnInit() {
    this.title.setTitle('用户管理');
    this.titleBar.set(this.titleinfo);
      this.route.paramMap.subscribe(
       (params: ParamMap) => {
         this.uid = Number(params.get('id'));
         this.managerService.getUser(Number(params.get('id'))).subscribe(data => {
           if (data['data']) {
             this.theUser = data['data'];
             this.DATA = [];
             console.log(this.theUser);
             let per = this.theUser.Permissions;
             for (let i = 3; i >= 0 ; i--) {
               let v = 1 << i;
               if (per - v >= 0) {
                 this.switchValue[i] = true;
                 per -= v;
               }
             }
             for (const key of Object.keys(this.theUser)) {
               if (this.theUser.hasOwnProperty(key)) {
                 const str = key + '~:' + this.theUser[key].toString();
                 this.DATA.push(str);
               }
             }
             setTimeout(() => {this.isSpinning = false; }, 300);
           }
         });
       });

  }

  permissionChange(i: number, on: boolean, permission_type: string) {
    let op_ty = 'remove';
    if (on) {op_ty = 'add'; }
    const theFun = function () {
      this.managerService.alterUserPermission([this.uid], op_ty, permission_type).subscribe(
        data => {
          console.log(data);
          if (data['data'] === 'ok') {
            this.switchValue[i] = !this.switchValue[i];
            this.message.create('success', `操作成功。`);
          } else {
            this.message.create('error', `出现了些问题x_x。`);
          }
        }
      );
    };
    this.confirmModal(theFun, '确认更改此用户权限');

  }

  // 确认提示框 theFun 确认时调用，title 提示语
  confirmModal(thFun: any, title: string) {
    this.modalService.confirm({
      nzTitle     : title,
      nzOkText    : '确认',
      nzOkType    : 'danger',
      nzOnOk      : () => thFun.call(this),
      nzCancelText: '取消',
      nzOnCancel  : () => console.log('Cancel')
    });
  }
}
