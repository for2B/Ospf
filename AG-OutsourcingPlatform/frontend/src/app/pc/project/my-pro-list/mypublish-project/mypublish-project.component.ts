import { Component, OnInit, OnDestroy } from '@angular/core';
import { PubHttpService } from '../../../../services/HttpService/pub-http.service';
import { GET_MYPUB_PROJECT, ProjectList, Project, NumberOfEachStatus, StatusTabGroup, StatusTab } from '../../../../const/const_chenhuiliang';
import { ProjectService } from '../../../../services/chenhuiliang/projectData.service';
import { ToastService } from '../../../../share/toast/toast.service';
import { LoadingService } from '../../../../share/lodding-message/Loading.service';
import { UpdateStatusService } from '../../../../services/chenhuiliang/update-status.service';
import { Observable, Subscribable, Subscription } from 'rxjs';

@Component({
  selector: 'app-mypublish-project',
  templateUrl: './mypublish-project.component.html',
  styleUrls: ['./mypublish-project.component.css']
})
export class MypublishProjectComponent implements OnInit ,OnDestroy{
  CurrentPage = 1
  ShowNums = 10
  TotalNums:number
  list_null= true
  status_nums = 12
  Project_Status = 0
  NumberWaitContract = 0
  NumberWaitSign = 0
  NumberSign = 0
  NumberEnd = 0
  projectlist:ProjectList
  project_list:Project[]
  ListenInMyPublishDataChange$:Subscription
  numberOfEachStatus = new NumberOfEachStatus()

  tabGroup = new StatusTabGroup([
    new StatusTab("全部",0),
    new StatusTab("等待承接",1),
    new StatusTab("等待签订合同",2),
    new StatusTab("签订合同",3),
    new StatusTab("项目结束",4)
  ])

  constructor(
    private ProjectService:ProjectService,
    private message:ToastService,
    private loading:LoadingService,
    private ListenInDataChange:UpdateStatusService,
  ) { }

  ngOnInit() {
    this.ShowNums = 10
    this.getMyPublishProjectList()
    this.getNumberOfEachStatus()
    this.listenDataChange()
  }

  //换页
  ChangePage(page:number){
    this.CurrentPage = page
    this.getMyPublishProjectList()
  }

  //变更显示条数
  ChangeSize(size:number){
    this.ShowNums = size
    this.CurrentPage = 1
    this.getMyPublishProjectList()
  }

  //改变显示状态
  changestatus(status:number){
    this.Project_Status = status
    this.CurrentPage = 1
    this.loading.hide()
    this.getMyPublishProjectList()
    this.SetTotalNumber()
  }


  //根据Project_Status获取项目数据
  getMyPublishProjectList(){
    this.loading.show()
    this.ProjectService.GetMyPublishProject(this.CurrentPage,this.ShowNums,this.Project_Status).subscribe(
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
    this.ProjectService.GetQuantityEachStatusMyPublish().subscribe(
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
      default:
          this.message.error("设置项目状态总数出错，无法匹配到对应的状态值")
          break;
    }
  }

  listenDataChange(){
    this.ListenInMyPublishDataChange$ = this.ListenInDataChange.ListenInMyPublishDataChange().subscribe(res=>{
      // this.getMyPublishProjectList()
    this.getNumberOfEachStatus()
    })
    console.log(this.ListenInMyPublishDataChange$)
  }

  ngOnDestroy(){
    this.ListenInMyPublishDataChange$.unsubscribe() //销毁取消订阅，防止内存泄漏
  }



}
