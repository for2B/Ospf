import { Component, OnInit } from '@angular/core';
import {TitleBarService} from '../../../share/osp-infobar/titleBar.service';
import { TitleInfo } from '../../../services/chencanxin/title-info';
import {UserManagementService} from '../../../services/panjiawei/user-management.service';
import {User} from '../../../services/panjiawei/users';
import {NzMessageService, NzModalService} from 'ng-zorro-antd';
@Component({
  selector: 'app-user-management',
  templateUrl: './user-management.component.html',
  styleUrls: ['./user-management.component.css']
})
export class UserManagementComponent implements OnInit {

  dataSet = [];
  PageIndex = 1;
  PageSize = 20;
  checkedList = [];
  UsersCount = 199;
  isLoading = true;
  searchValue = '';
  userList: User[];
  allChecked = false;
  indeterminate = false;
  selectedStatValue = '4';  // 0=>禁用； 1=>启用 ；4=>不限（不要问我为什么没有2，因为2被我忘了）
  selectedPowerValue = '0'; // 0=>不限；1=>启用/登录；2=>承接项目；4=>发布项目；8=>评价
  displayData: Array<User> = [];
  titleinfo: TitleInfo = new TitleInfo('anticon anticon-user', '用户管理', '查看和管理用户的信息。');

  constructor(
    private titleBar: TitleBarService,
    private modalService: NzModalService,
    private message: NzMessageService,
    private user_management_service: UserManagementService) { }

  ngOnInit(): void { this.titleBar.set(this.titleinfo); this.getUserList(); }

  currentPageDataChange($event: Array<User>): void { this.displayData = $event; }

  refreshStatus(value: boolean, data: any): void { data.checked = value; }

  checkAll(value: boolean): void {
    this.displayData.forEach(data => {data.checked = value; });
    this.allChecked = true;
  }

  // 查找条重置数值
  reset() {
    this.selectedStatValue = '4';
    this.selectedPowerValue = '0';
    this.searchValue = '';
  }
  // 查找
  search() { this.DataChange(true); }

  // 页码变了，要换数据
  PageIndexChange() { this.DataChange(false); }

  // 每页显示行数变了，还是要换数据
  PageSizeChange() { this.DataChange(true); }
  // 获取数据
  DataChange(Pto1: boolean) {this.isLoading = true;
    this.cleanChecked();
    this.getUserList(Pto1);
  }

  // 获取用户列表，可用于搜索，Pto1：意思就是是否获取第一页的数据，如果不是的话就按当前的页码数请求咯
  getUserList(Pto1?: boolean) {
    this.isLoading = true;
    this.allChecked = false;
    if (Pto1) {this.PageIndex = 1; }
    this.user_management_service.getUsersList(this.PageIndex, this.PageSize, this.searchValue.trim(),
      this.selectedPowerValue.trim(),
     this.selectedStatValue.trim()).subscribe(
      data => {
        this.isLoading = false;
        console.log(data);
        if (data['data']['Alluser'] && data['data']['Count']) {
          this.dataSet = [];
          this.userList = data['data']['Alluser'];
          this.UsersCount = data['data']['Count'];
          for (let i = 0; i < this.userList.length; i++) {
            this.userList[i].checked = false;
            this.dataSet.push(this.userList[i]);
          }
        }
      });
  }

  // 清除全选
  cleanChecked() {
   for (let i = 0; i < this.dataSet.length; i++) {
     this.dataSet[i].checked = false;
   }
    this.allChecked = false;
  }
  /// 更改用户权限，参数不懂？？？ 看下面
  // operate_type 就是 添加add， 还有 移除remove
  // permission_type? 权限类型 有“AC_1”,“AC_2”“AC_4”“AC_8”
  //   AC_1: 1,  // 登录，启用，禁用权限，都是这个啦
  //   AC_2: 2,  // 承接项目权限，不知道加上有什么用，加上就对了
  //   AC_4: 4,  // 发布项目权限
  //   AC_8: 8   // 这个就是评论权限了，万一有差评师，得禁言
  // 不要吐槽，没办法了，智商有限，只能这设计了
  alterUser(operate_type: string, permission_type: string) {
    const theFun = function() {
      this.checkedList = [];
      for (let i = 0; i < this.dataSet.length; i++) {
        if (this.dataSet[i].checked) {this.checkedList.push(this.dataSet[i].Uid); }
      }
      if (this.checkedList.length === 0) {
        this.message.create('warning', `未指定用户。`);
        return;
      }
      console.log(this.checkedList);
      this.user_management_service.alterUserPermission(this.checkedList, operate_type, permission_type).subscribe(
        data => {
          console.log(data);
          if (data['data'] && data['data'] === 'ok') {
            this.getUserList();
            this.message.create('success', '操作成功。');
          } else {
            this.message.create('error', `出现了些问题x_x`);
          }
        }
      );
    };
    this.confirmModal(theFun, '确认修改已选择用户权限?');
  }
  // 禁用/启用开关
  switchChange(i: number, uid: number) {
    console.log(i);
    let operator = 'add';
    let tip = '确认启用该用户?';
    if (this.dataSet[i].State) {operator = 'remove';  tip = '确认禁用该用户?'; }
    const theFun = function() {
      this.user_management_service.alterUserPermission([uid], operator, 'AC_1').subscribe(
        data => {
          if (data['data'] === 'ok') {
            this.dataSet[i].State = !this.dataSet[i].State;
            this.message.create('success', `操作成功。`);
          } else {
            this.message.create('error', `出现了些问题x_x。`);
          }
        });
    };
    this.confirmModal(theFun, tip);
  }
  // 批量删除用户
  delUser() {
    const delUFun = function() {
      this.checkedList = [];
      let datalen = 0;
      for (let i = 0; i < this.dataSet.length; i++) {
        if (this.dataSet[i].checked) {
          this.checkedList.push(this.dataSet[i].Uid);
          datalen++;
        }
      }
      if (this.checkedList.length === 0) {
        this.message.create('warning', `未指定用户。`);
        return;
      }
      console.log(this.checkedList);
      this.user_management_service.deleteUsers(this.checkedList).subscribe(
        data => {
          if (data['data']) {
            if (this.dataSet.length === datalen && this.PageIndex > 1) {
              this.PageIndex -= 1;
            }
            this.getUserList();
            this.allChecked = false;
          }
        });
    };
    this.confirmModal(delUFun, '确认删除已选择用户?');
  }
  //  删除单个用户（实际上调用了 批量删除用户delUser()）
  delOneUser(uid: number) {
    const delOFun = function() {
      // this.isLoading = true;
    this.user_management_service.deleteUsers([uid]).subscribe(
      data => {
        if (data['data']) {
          if (this.dataSet.length === 1 && this.PageIndex > 1) {
            this.PageIndex -= 1;
          }
          this.getUserList();
        }
      }); };
    this.confirmModal(delOFun, '确认删除该用户?');
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
