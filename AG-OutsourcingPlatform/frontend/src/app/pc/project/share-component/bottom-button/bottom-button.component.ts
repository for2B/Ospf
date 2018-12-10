import { Component, OnInit, Input, Output, EventEmitter, TemplateRef, OnDestroy } from '@angular/core';
import { Project, ROUTES } from '../../../../const/const_chenhuiliang';
import { Router } from '@angular/router';
import { DeliveryprojectService } from '../../../../services/chenhuiliang/deliveryproject.service';
import { ProjectService } from '../../../../services/chenhuiliang/projectData.service';
import { ModalService } from '../../../../share/modal/modal.service';
import { NzModalService } from 'ng-zorro-antd';
import { ToastService } from '../../../../share/toast/toast.service';
import { UpdateStatusService } from '../../../../services/chenhuiliang/update-status.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-bottom-button',
  templateUrl: './bottom-button.component.html',
  styleUrls: ['./bottom-button.component.css']
})
export class BottomButtonComponent implements OnInit,OnDestroy {
  ButtonGroupType:string //contract-deatil;amend-end;sign-end;end;abandon;download-end;
  @Input() Identity:number  //
  @Input() showPlace:string //item（在项目卡片上显示）  detail（在项目详情中显示）
  @Input() project:Project
  ListenInProjectStatus$:Subscription
  ListenInIdentity$:Subscription
  
  

  constructor(
    private router:Router,
    private deliveryprojectService: DeliveryprojectService,
    private projectService:ProjectService,
    private modalService: NzModalService,
    private toast:ToastService,
    private updateStatus:UpdateStatusService
  ) { }

  ngOnInit() {
    if(!this.Identity){
      this.projectService.GetUserIdentity(this.project.project_id).subscribe(
        res=>{
          this.Identity = res.data.identity
          this.PickButtonType()
        }
      )
    }else{
      this.PickButtonType()
    }
    this.ListenStatus()
    this.ListenUpdateIdentity()
  }

 

  PickButtonType(){
    switch(this.Identity){
      case 1: //承包者
        switch(this.project.project_status){
          case 1:
          this.ButtonGroupType = "abandon"
          break
          case 2:
          this.ButtonGroupType = "download-end"
          break
          default:
          this.ButtonGroupType = ""
          break
        }
        break
      case 2: //发布者
        switch(this.project.project_status){
          case 1:
            this.ButtonGroupType = "amend-end"
            break
          case 2:
            this.ButtonGroupType = "sign-end"
            break
          case 3:
            this.ButtonGroupType = "end"
            break
          default:
            this.ButtonGroupType = ""
            break
        }
        break
      case 3://我参与的
        this.ButtonGroupType = "contract-deatil"
        break
      case 4: //平台上所有（其他人身份）
        this.ButtonGroupType = "contract-deatil"
        break
    }
  }

  //查看详情
  seeDetail(){
    this.deliveryprojectService.setProject(this.project)
    this.router.navigate([ROUTES.seeDetail.route])
  }


  //承接项目(参与竞标)
  contractProject(confirmcontract: TemplateRef<{}>):void{
    this.modalService.confirm({
      nzContent: confirmcontract,
      nzOnOk:()=>{
        this.projectService.JoinBidding(this.project.project_id).subscribe(
          res=>{
            this.toast.success(
              `<p><span style=\"font-size: 10px;\">系统提示：恭喜你参与竞选承接该项目，当前有
              <strong>${this.project.competitors_nums}</strong>人竞选承接该项目，项目需求方会根据系统里的
              【<strong>个人能力资质</strong>】确认最终承接人，请等候佳音！</span></p>`
            )
            this.updateStatus.PlafDataUpdate() 
            this.updateStatus.UpdateIdentityLocal(1,this.project.project_id,this.showPlace)
          }
        )        
      }
    })
  }

  //修改项目
  amendProjec():void{
    console.log(this.project)
    this.deliveryprojectService.setProject(this.project)
    this.router.navigate([ROUTES.amendProject.route])
    
  }

  //签订合同
  signPaction(updatesign: TemplateRef<{}>):void{
    this.modalService.confirm({
      nzContent:updatesign,
      nzOnOk:()=>{
        this.projectService.UpdateProjectStatus(this.project.project_id,3).subscribe(
          res=>{
            this.toast.success(`项目【<strong>${this.project.project_name}</strong>】状态已更新为【<strong>签订合同</strong>】`)
            this.project.project_status = 3  //修改本project的status，以便点击item进入详情页的时候能保存最新的项目状态
            this.updateStatus.UpdateProjectStatusLocal(3,this.project.project_id,this.showPlace)  //更新项目状态的值，通知所有监听该值的地方更新样式和数据
            this.updateStatus.UpdateMyPublishList(this.showPlace)//通知我发布的项目列表拉取最新的数据
            if(this.showPlace == 'detail'){
              this.deliveryprojectService.setProject(this.project) //在detail中修改status时要保持状态都本地，以便刷新后状态值为最新的
            }
          }
        )
      }
    })
  }

  //下载合同
  downloadPaction(){
    console.log("TODO downloadPaction"+this.project.project_id)
    //TODO something
  }

  //结束项目
  endProject(){
    console.log("TODO endProject"+this.project.project_id)
    this.modalService.confirm({
      nzContent:"是否确定结束项目？",
      nzOnOk:()=>{
        this.projectService.UpdateProjectStatus(this.project.project_id,4).subscribe(
          res=>{
            this.toast.success(`项目【<strong>${this.project.project_name}</strong>】状态已更新为【<strong>项目结束</strong>】`)
            this.project.project_status = 4
            this.updateStatus.UpdateProjectStatusLocal(4,this.project.project_id,this.showPlace)
            this.updateStatus.UpdateMyPublishList(this.showPlace)
            if(this.showPlace == 'detail'){
              this.deliveryprojectService.setProject(this.project)
            }
          }
        )
      }
    })
  }

  //放弃竞标
  abandonBidding(){
    this.modalService.confirm({
      nzContent:"是否确定放弃竞标该项目?",
      nzOnOk:()=>{
        this.projectService.GiveUpBidding(this.project.project_id).subscribe(
          res=>{
            this.project.project_status = 1
            this.updateStatus.UpdateProjectStatusLocal(1,this.project.project_id,this.showPlace)
            this.toast.success("成功放弃竞标！")
            this.updateStatus.UpdateMyContractorList(this.showPlace) //通知我发布的项目列表更新数据
            this.updateStatus.UpdateIdentityLocal(4,this.project.project_id,this.showPlace) //通知该id的项目，此用户身份已经改变,更新数据和样式
            if(this.showPlace == 'detail'){
              this.deliveryprojectService.setProject(this.project)
            }
          }
        )
       
      }
    })
    //TODO something
  }

  //放弃承接该项目
  abandonProject(){
    this.modalService.confirm({
      nzContent:"是否确定放弃承包该项目?",
      nzOnOk:()=>{
        this.projectService.GiveUpContractor(this.project.project_id).subscribe(res=>{
          this.updateStatus.UpdateProjectStatusLocal(1,this.project.project_id,this.showPlace)
          this.project.project_status = 1
          if(this.showPlace == 'detail'){
            this.deliveryprojectService.setProject(this.project)
          }
        })
      }
    })
  }

  //监听状态值变化
  ListenStatus(){
    this.ListenInProjectStatus$ = this.updateStatus.ListenInProjectStatus().subscribe(
      res=>{
        if(this.project.project_id == res.ProjectId){
          this.project.project_status = res.ProjectStatus
          // if(res.ShowPlacd == 'detail'){
            this.PickButtonType()
          // }
        }
      }
    )
  }

  //监听身份是否更新
  ListenUpdateIdentity(){
    this.ListenInIdentity$ = this.updateStatus.ListenInIdentity().subscribe(
      res=>{
        if(this.project.project_id == res.ProjectId){
          this.Identity = res.Identity
          // if(res.ShowPlacd == 'detail'){
            this.PickButtonType()
          // }
        }
      }
    )
  }
  
  ngOnDestroy(){
    this.ListenInIdentity$.unsubscribe()
    this.ListenInProjectStatus$.unsubscribe()
  }

}
    