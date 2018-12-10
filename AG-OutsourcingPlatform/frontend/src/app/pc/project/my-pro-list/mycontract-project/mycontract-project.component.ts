import { Component, OnInit, OnDestroy } from '@angular/core';
import { Project, NumberOfEachStatus, StatusTabGroup, StatusTab } from '../../../../const/const_chenhuiliang';
import { ProjectService } from '../../../../services/chenhuiliang/projectData.service';
import { ToastService } from '../../../../share/toast/toast.service';
import { LoadingService } from '../../../../share/lodding-message/Loading.service';
import { UpdateStatusService } from '../../../../services/chenhuiliang/update-status.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-mycontract-project',
  templateUrl: './mycontract-project.component.html',
  styleUrls: ['./mycontract-project.component.css']
})
export class MycontractProjectComponent implements OnInit,OnDestroy {
  status_nums = 12
  CurrentPage = 1
  ShowNums : number
  TotalNums:number
  list_null= false
  numberOfEachStatus = new NumberOfEachStatus()
  Project_Status = 0
  project_list:Project[]
  ListenInMyContractDataChange$ :Subscription
  tabGroup = new StatusTabGroup([
    new StatusTab("全部",0),
    new StatusTab("参与竞标",1),
    new StatusTab("等待签订合同",2),
    new StatusTab("签订合同",3),
    new StatusTab("项目结束",4),
    new StatusTab("竞标失败",5)

  ])

  constructor(
    private ProjectService:ProjectService,
    private toast:ToastService,
    private message:ToastService,
    private loading:LoadingService,
    private ListenInDataChange:UpdateStatusService,
  ) { }

  ngOnInit() {
    this.ShowNums = 20
    this.getMycontractProjectList()
    this.getNumberOfEachStatus()
    this.listenDataChange()
  }


  changestatus(status:number){
    this.Project_Status = status
    this.loading.hide()
    this.CurrentPage = 1
    this.getMycontractProjectList()
  }

  ChangePage(page:number){
    this.CurrentPage = page
    this.getMycontractProjectList()
  }

  ChangeSize(size:number){
    this.ShowNums = size
    this.CurrentPage  = 1
    this.getMycontractProjectList()
  }


  //根据Project_Status获取项目列表
  getMycontractProjectList(){
    this.loading.show()
    this.ProjectService.GetMyContractProject(this.CurrentPage,this.ShowNums,this.Project_Status).subscribe(
      res=>{
        this.loading.hide()
        if(!res.data){
          this.list_null = true
          this.project_list = undefined
          return
        }
        this.list_null = false
        this.project_list = res.data
      }
    )
  }

  getNumberOfEachStatus(){
    this.ProjectService.GetQuantityEachStatusMyContract().subscribe(
      res=>{
        this.numberOfEachStatus = res.data
        this.SetTotalNumber()
      }
    )
  }

  SetTotalNumber(){
    switch (this.Project_Status) {
      case 0:
          this.TotalNums = this.numberOfEachStatus.total;
          break;
      case 1:
          this.TotalNums = this.numberOfEachStatus.status_one;
          break;
      case 2:
          this.TotalNums = this.numberOfEachStatus.status_two;
          break;
      case 3:
          this.TotalNums = this.numberOfEachStatus.status_three;
          break;
      case 4:
          this.TotalNums = this.numberOfEachStatus.status_four;
          break;
      case 5:
          this.TotalNums = this.numberOfEachStatus.status_five;
          break;
      default:
          this.message.error("设置项目状态总数出错，无法匹配到对应的状态值")
          break;
    }
  }

  listenDataChange(){
    this.ListenInMyContractDataChange$ = this.ListenInDataChange.ListenInMyContractDataChange().subscribe(res=>{
      console.log('ChangeMyContract')
      // this.getMycontractProjectList()
      this.getNumberOfEachStatus()
    })
  }

  ngOnDestroy(){
    this.ListenInMyContractDataChange$.unsubscribe()
  }

}
