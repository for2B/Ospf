import { Component, OnInit, Input, TemplateRef } from '@angular/core';
import { DeliveryprojectService } from '../../../../services/chenhuiliang/deliveryproject.service';
import { Project } from '../../../../const/const_chenhuiliang';
import { ProjectService } from '../../../../services/chenhuiliang/projectData.service';
import { NzModalService } from 'ng-zorro-antd';
import { UpdateStatusService } from '../../../../services/chenhuiliang/update-status.service';

@Component({
  selector: 'app-bidlistoperatebutton',
  templateUrl: 'bidlistoperatebutton.component.html',
  styleUrls: ['bidlistoperatebutton.component.css']
})
export class BidlistoperatebuttonComponent implements OnInit {

  @Input() BiderIdentity:number //1 承包者 0 竞标者
  @Input() BidderId:number 
  @Input() BidderName:string
  project :Project 
  ButtonType:string
  constructor(
    private deliverprojectService:DeliveryprojectService,
    private projectService : ProjectService,
    private modalService: NzModalService,
    private updatestatus:UpdateStatusService
  ) {
    this.project = deliverprojectService.getProject()
  }

  ngOnInit() { 
    this.ChooesButtonType()
  }

  ChooesButtonType(){
    //TODO 状态不为1或者2时，按钮消失，即此时已经不需要按钮操作了
    switch(this.project.project_status){//根据项目状态来判断当前按钮是否可选，如果项目状态为1，则此时还没有选择承包者，不为1 则已经选择了承包者
      case 1:
        this.ButtonType = "chooes"
        break;
      default:
        if(this.BiderIdentity == 1){
          this.ButtonType = "cancel"
        }
        if(this.BiderIdentity == 0){
          this.ButtonType = "forbidchooes"
        }
        break;
    }
  }

  chooes(ChooesContractor:TemplateRef<{}>){
    this.modalService.confirm({
      nzContent:ChooesContractor,
      nzOnOk:()=>{
        this.projectService.SelectContractor(this.BidderId,this.project.project_id).subscribe(res=>{
            this.updatestatus.UpdateProjectStatusLocal(2,this.project.project_id,'detail')
            this.project.project_status = 2
          }
        )
      }
    })
  } 

  cancel(CancelContractor:TemplateRef<{}>){
    this.modalService.confirm({
      nzContent:CancelContractor,
      nzOnOk:()=>{
        this.projectService.GiveUpContractor(this.project.project_id).subscribe(res=>{
          this.updatestatus.UpdateProjectStatusLocal(1,this.project.project_id,'detail')
          this.project.project_status = 1
          this.BiderIdentity = 0 //好像可以不用刷新竞标列表
        })
      }
    })
  }

}
