import { Component, OnInit, Input, OnDestroy } from '@angular/core';
import { Project } from '../../../../const/const_chenhuiliang';
import { UpdateStatusService } from '../../../../services/chenhuiliang/update-status.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-prostatus',
  templateUrl: './prostatus.component.html',
  styleUrls: ['./prostatus.component.css']
})
export class ProstatusComponent implements OnInit,OnDestroy{
  @Input() Status : number;
  @Input() Identity : number; //1 承包者 2  发布者  3我参与的 4路人   ’promulgator’(发布者) ‘contractor’ (承包者) ‘teammate’ (普通成员)’others’ (路人）
  @Input() ProjectId:number;
  @Input() showPlace:string //item（在项目卡片上显示）  detail（在项目详情中显示）
  ListenInProjectStatus$:Subscription
  ListenInIdentity$:Subscription

  config : {text:"",color:""}
  ProjectStatus_publisher = new Map()
  ProjectStatus_contractor = new Map()
  text : string
  backgoundc : string
  constructor(
    private updatestatus:UpdateStatusService
  ) { 
    this.ProjectStatus_publisher.set(1,{text:"等待承接",color:"#F95E5A"})
    this.ProjectStatus_publisher.set(2,{text:"等待签订合同",color:"#FE9400"})
    this.ProjectStatus_publisher.set(3,{text:"签订合同",color:"#4BD863"})
    this.ProjectStatus_publisher.set(4,{text:"项目结束",color:"#464C5B"})
    this.ProjectStatus_publisher.set(5,{text:"",color:"#ffffff00"})

    this.ProjectStatus_contractor.set(1,{text:"参与竞标",color:"#F95E5A"})
    this.ProjectStatus_contractor.set(2,{text:"等待签订合同",color:"#FE9400"})
    this.ProjectStatus_contractor.set(3,{text:"签订合同",color:"#4BD863"})
    this.ProjectStatus_contractor.set(4,{text:"项目结束",color:"#464C5B"})
    this.ProjectStatus_contractor.set(5,{text:"竞标失败",color:"#464C5B"})
  }
  ngOnInit() {
    this.ChooesStatusType()
    this.ListenUpdateStatus()
    this.ListenUpdateIdentity()
  }

  ChooesStatusType(){
    switch(this.Identity){
      case 1:
        this.config = this.ProjectStatus_contractor.get(this.Status);
        break;
      case 4:
        this.config = this.ProjectStatus_publisher.get(1);
        break;
      default:
        this.config = this.ProjectStatus_publisher.get(this.Status);
        break;
    }
   this.backgoundc =this.config.color
   this.text = this.config.text
  }

  //监听状态值是否更新
    ListenUpdateStatus(){
    this.ListenInProjectStatus$ = this.updatestatus.ListenInProjectStatus().subscribe(
      res=>{
        if(this.ProjectId == res.ProjectId){
            this.Status = res.ProjectStatus
            // if(res.ShowPlacd == 'detail'){
              this.ChooesStatusType()
            // }
        }
      }
    )
  }

  //监听身份是否更新
  ListenUpdateIdentity(){
    this.ListenInIdentity$ = this.updatestatus.ListenInIdentity().subscribe(
      res=>{
        if(this.ProjectId == res.ProjectId){
          this.Identity = res.Identity
          // if(res.ShowPlacd == 'detail'){
            this.ChooesStatusType()
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


