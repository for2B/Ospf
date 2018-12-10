import { Component, OnInit, OnChanges,Input, SimpleChange, OnDestroy } from '@angular/core';
import { ProjectService } from '../../../../services/chenhuiliang/projectData.service';
import { Bider, Biddinglist, Project } from '../../../../const/const_chenhuiliang';
import { BiddingService } from '../../../../services/chenhuiliang/bidding.service';
import { DeliveryprojectService } from '../../../../services/chenhuiliang/deliveryproject.service';
import { UpdateStatusService } from '../../../../services/chenhuiliang/update-status.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-biddinglist',
  templateUrl: './biddinglist.component.html',
  styleUrls: ['./biddinglist.component.css']
})
export class BiddinglistComponent implements OnInit,OnDestroy {

  @Input() 
  set Identity(identity:number){
      this.identity = identity
  }

  get Identity():number{
    return this.identity
  }

  identity:number//1 承包者 2 发布者 3普通成员 4路人（平台其他用户）
  project:Project
  pageIndex = 1;//当前页数
  pageSize = 10; //显示条数
  total = 1;  //总数
  biddingList :Bider[] = [];  //源数据
  loading = true; //是否加载中
  sortValue = null;
  sortKey = null;
  searchGenderList: string[] = [];
  ListenInProjectStatus$:Subscription
  ListenInIdentity$:Subscription

  //排序调用函数
  sort(sort: { key: string, value: string }): void {
    this.sortKey = sort.key;
    this.sortValue = sort.value;
    this.sortBiddlingList();
  }

  constructor(
    private projectService: ProjectService,
    private biddinglistService:BiddingService,
    private deliverProjectService:DeliveryprojectService,
    private upDateStatus:UpdateStatusService
  ) { 
    this.project = deliverProjectService.getProject()
  }

  ngOnInit() {
    this.getBiddingList();
    this.listenStatus();
    this.listenUpdateIdentity();
  }

  //获取竞标列表
  getBiddingList(sortorder?:string): void {
    this.loading = true;
    this.biddinglistService.GetBiddingList(this.project.project_id,this.pageIndex,this.pageSize,sortorder).subscribe(
      res=>{
        this.total = res.data.total
        this.loading = false
        if(!res.data.bidding_list){
          this.biddingList = []
        }else{
          this.biddingList = res.data.bidding_list
        }
      }
    )
  }

  //排序触发事件
  sortBiddlingList(reset: boolean = false):void{
    if (reset) {
      this.pageIndex = 1;
    }
    if(this.sortKey && this.sortValue){
      this.getBiddingList(this.sortValue)
    }else{
      this.getBiddingList()
    }
  }

  //监听项目状态变化
  listenStatus():void{
    this.ListenInProjectStatus$ = this.upDateStatus.ListenInProjectStatus().subscribe(
      res=>{
        if(this.project.project_id === res.ProjectId){
          this.project.project_status = res.ProjectStatus
          this.sortBiddlingList()
        }
      }
    )
  }

  //监听identity变化
  listenUpdateIdentity():void{
    this.ListenInIdentity$ = this.upDateStatus.ListenInIdentity().subscribe(
      res=>{
        if(this.project.project_id === res.ProjectId){
          this.Identity = res.Identity
          this.sortBiddlingList()
        }
      }
    )
  }

  ngOnDestroy(){
    this.ListenInIdentity$.unsubscribe()
    this.ListenInProjectStatus$.unsubscribe()
  }

}
