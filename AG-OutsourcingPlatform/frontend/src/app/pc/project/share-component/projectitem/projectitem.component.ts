import { Component, OnInit, Input, OnDestroy } from '@angular/core';
import { Project, ROUTES } from '../../../../const/const_chenhuiliang';
import { ActivatedRoute, Router } from '@angular/router';
import { DeliveryprojectService } from '../../../../services/chenhuiliang/deliveryproject.service';
import { ProjectService } from '../../../../services/chenhuiliang/projectData.service';
import { UpdateStatusService } from '../../../../services/chenhuiliang/update-status.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-projectitem',
  templateUrl: './projectitem.component.html',
  styleUrls: ['./projectitem.component.css']
})
export class ProjectitemComponent implements OnInit,OnDestroy {
  @Input() project : Project
  @Input() Identity  : number
  money=1
  ListenInProjectStatus$:Subscription


  constructor(
    private router:Router,
    private deliveryprojectService: DeliveryprojectService,
    private projectService:ProjectService,
    private updateStatus:UpdateStatusService,
  ) { }

  ngOnInit() {
    this.listenStatusChange()
  }
  //查看详情
  seeDetail(){
      this.deliveryprojectService.setProject(this.project)
      this.router.navigate([ROUTES.seeDetail.route])

    }

  //监听状态变化
  listenStatusChange(){
    this.ListenInProjectStatus$ = this.updateStatus.ListenInProjectStatus().subscribe(
      res=>{
        if(this.project.project_id == res.ProjectId){
          this.project.project_status = res.ProjectStatus
        }
      }
    )
  }

  ngOnDestroy(){
    this.ListenInProjectStatus$.unsubscribe()
  }
}
