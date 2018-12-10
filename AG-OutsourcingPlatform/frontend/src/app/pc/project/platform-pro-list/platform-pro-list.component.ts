import { Component, OnInit, OnDestroy } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { TitleBarService } from '../../../share/osp-infobar/titleBar.service';
import { PubHttpService } from '../../../services/HttpService/pub-http.service';
import { TitleInfo } from '../../../services/chencanxin/title-info';
import {  GET_PLATFORM_PROJECT, ProjectList, Project } from '../../../const/const_chenhuiliang';
import { ProjectService } from '../../../services/chenhuiliang/projectData.service';
import { UpdateStatusService } from '../../../services/chenhuiliang/update-status.service';
import { Subscription } from 'rxjs';
import { LoadingService } from '../../../share/lodding-message/Loading.service';

@Component({
  selector: 'app-platform-pro-list',
  templateUrl: './platform-pro-list.component.html',
  styleUrls: ['./platform-pro-list.component.css'],
})
export class PlatformProListComponent implements OnInit,OnDestroy {
  projectlist:ProjectList
  project_list:Project[]
  list_null= false
  CurrentPage = 1
  ShowNums = 10
  TotalNums:number
  GetOption = {
    cycle_range_end:999999999999,
    cycle_range_start:0,
    date_range_end:6311433600,
    date_range_start:0,
    money_range_end:999999999999,
    money_range_start:0,
    current_page:this.CurrentPage,
    show_nums:this.ShowNums
  }
  PlafListenUpdateMsg$:Subscription
  
  titleinfo: TitleInfo = new TitleInfo('anticon anticon-appstore-o', '所有项目', '在这里你可以看到所有等待被承接的项目');
  
  constructor(   
    private title: Title,
    private titleBar: TitleBarService,
    private ProjectService:ProjectService,
    private UpdateProject:UpdateStatusService,
    private loading:LoadingService,
  ) { }

  ngOnInit() {
    this.title.setTitle('所有项目')
    this.titleBar.set(this.titleinfo)
    this.list_null = true
    this.getList()
    this.ListenIfContract()

  }




  ChangePage(page:number){
    this.GetOption.current_page = page
    this.getList()
  }

  ChangeSize(size:number){
    this.GetOption.show_nums = size
    this.getList()
  }

  getselect(GetOption:any){
    this.GetOption.cycle_range_start = GetOption.cycle_range_start
    this.GetOption.cycle_range_end = GetOption.cycle_range_end
    this.GetOption.money_range_start = GetOption.money_range_start
    this.GetOption.money_range_end = GetOption.money_range_end
    this.GetOption.date_range_start = GetOption.date_range_start
    this.GetOption.date_range_end = GetOption.date_range_end
    this.GetOption.current_page = 1
    this.getList()
  }

  getList(){
    this.loading.show()
    this.ProjectService.GetPlatformProjectList(this.GetOption).subscribe(
      res=>{
        this.loading.hide()
        if(res.data == undefined){
          this.list_null = true
          this.project_list = undefined
          return
        }
        if(res.data.total == 0){
          this.list_null = true
          this.project_list = undefined
          return
        }
        this.list_null = false
        this.project_list = res.data.project_list
        this.TotalNums = res.data.total
        console.log(this.TotalNums)
      },
      err=>{
      }
    )
  }

  ListenIfContract(){
    this.PlafListenUpdateMsg$ = this.UpdateProject.PlafListenUpdateMsg().subscribe(
      res=>{
        this.getList()
      }
    )
  }

  ngOnDestroy(){
    this.PlafListenUpdateMsg$.unsubscribe()
  }

}
