import { Component, OnInit, OnDestroy } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { ProjectService } from '../../../services/chenhuiliang/projectData.service';
import { TitleInfo } from '../../../services/chencanxin/title-info';
import { TitleBarService } from '../../../share/osp-infobar/titleBar.service';
import {  Router } from '@angular/router';
import { Project, ROUTES, teammates, Teammate } from '../../../const/const_chenhuiliang';
import { DeliveryprojectService } from '../../../services/chenhuiliang/deliveryproject.service';
import { UpdateStatusService } from '../../../services/chenhuiliang/update-status.service';
import { Subscription } from 'rxjs';
@Component({
  selector: 'app-Project-detail',
  templateUrl: './project-detail.component.html',
  styleUrls: ['./project-detail.component.css']
})
export class ProjectDetailComponent implements OnInit,OnDestroy {

  titleInfo: TitleInfo
  project:Project = new Project()
  userIdentity:number //1 承包者 2 发布者 3普通成员 4 路人
  breadCrumbText:string
  teammates :Teammate[]
  ListenInIdentity$:Subscription
  ListenInProjectStatus$:Subscription

  constructor(
    private title:Title,
    private projectService:ProjectService,
    private titleBar: TitleBarService,
    private route : Router,
    private deliveryprojectService:DeliveryprojectService,
    private updateStatus:UpdateStatusService,
  ) {
  }

  ngOnInit() {
    this.project = this.deliveryprojectService.getProject()
    this.title.setTitle("项目详情")
    this.titleInfo = new TitleInfo('anticon anticon-scan', this.project.project_name, '在这里可以查看到有关项目的所有信息');
    this.titleBar.set(this.titleInfo)
    this.getUserIdentityForThisProject()
    this.listenUpdateIdentity()
    this.listenUpdateStatus()
    this.getTeammates()
  }

  //根据Useridentity来返回上一个页面
  goBack(){
    //根据当前用户的身份返回到对应的地方，所有项目，我创建的，我承接的
    switch (this.userIdentity){
      case 1:
          this.route.navigate([ROUTES.mypublishlist.route])
          break;
      case 2:
          this.route.navigate([ROUTES.mypublishlist.route])
          break;
      case 3:
          this.route.navigate([ROUTES.mypublishlist.route])
          break;
      case 4:
          this.route.navigate([ROUTES.plafprolist.route])
          break;
    }
  }

  //根据UserIdentity来设置面包屑
  setBreadcrumb(){
    switch (this.userIdentity){
      case 1:
          this.breadCrumbText = '我承包的'
          break;
      case 2:
          this.breadCrumbText = '我创建的'
          break;
      case 3:
          this.breadCrumbText = '我参与的'
          break;
      case 4:
          this.breadCrumbText = '所有项目'
          break;
    }
  }

  //获取团队成员
  getTeammates(){
    this.projectService.GetTeammate(this.project.project_id).subscribe(
      res=>{
        this.teammates = res.data
        this.setTeammateButton()
      },
      err=>{}
    )
  }


  setTeammateButton(){
    for(var index in this.teammates){
      if(this.teammates[index].IsMe === 1){ 
        this.teammates[index].ButtonStyle = 4//4   查看
      }else{
        switch(this.userIdentity){
          case 1:
            this.teammates[index].ButtonStyle = 1; //   1   评价，查看，移除 
            break;
          case 2:
            this.teammates[index].ButtonStyle = 3;// 3   评价和查看
            break;
          case 3:
            this.teammates[index].ButtonStyle = 3;
            break;
        }
      }
    }
  }

  //获取用户相对于该项目的身份 1 承包者 2发布者 3普通成员 4 路人
  getUserIdentityForThisProject(){
    this.projectService.GetUserIdentity(this.project.project_id).subscribe(
      res=>{
        this.userIdentity = res.data.identity    
        this.setBreadcrumb()
      }
    )
  }

  listenUpdateIdentity(){
    this.ListenInIdentity$ = this.updateStatus.ListenInIdentity().subscribe(
      res=>{
        if(this.project.project_id == res.ProjectId){
          this.userIdentity = res.Identity
          this.setBreadcrumb()
        }
      }
    )
  }

  listenUpdateStatus(){
    this.ListenInProjectStatus$ =  this.updateStatus.ListenInProjectStatus().subscribe(res=>{
      console.log(res)
      if(res.ShowPlacd == 'detail'&&res.ProjectId == this.project.project_id&&res.ProjectStatus!=1){
        this.getTeammates()
      }
    })
  }

  ngOnDestroy(){
    this.ListenInIdentity$.unsubscribe()
    this.ListenInProjectStatus$.unsubscribe()
  }

}
 