import { Component, OnInit, OnDestroy } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { TitleBarService } from '../../../share/osp-infobar/titleBar.service';
import { NzMessageService } from 'ng-zorro-antd';
import { Router } from '@angular/router';
import { ProjectService } from '../../../services/chenhuiliang/projectData.service';
import { TitleInfo } from '../../../services/chencanxin/title-info';
import { ROUTES, Project, FbData } from '../../../const/const_chenhuiliang';
import { DeliveryprojectService } from '../../../services/chenhuiliang/deliveryproject.service';
import { Observable, Subscription } from 'rxjs';

@Component({
  selector: 'app-amendproject',
  templateUrl: './amendproject.component.html',
  styleUrls: ['./amendproject.component.css']
})
export class AmendprojectComponent implements OnInit,OnDestroy {
  validateForm: FormGroup;
  ProText=""
  isLoading = false
  project:Project
  AmendProject$:Subscription


  updateConfirmValidator(): void {
    /** wait for refresh value */
    Promise.resolve().then(() => this.validateForm.controls.checkPassword.updateValueAndValidity());
  }

  constructor(
    private fb: FormBuilder,
    private titleBar: TitleBarService,
    private message:NzMessageService,
    private route:Router,
    private ProjectService:ProjectService,
    private deliveryproject:DeliveryprojectService
  ) { }
  titleinfo: TitleInfo = new TitleInfo('anticon anticon-plus-circle-o' , '修改项目' , '完善项目信息');

  ngOnInit() {
    this.project = this.deliveryproject.project
    console.log(this.project)
    this.validateForm = this.fb.group({
      title            : [ null, [ Validators.required,Validators.maxLength(20) ] ],
      money            : [ null, [ Validators.required ] ],
      cycle            : [ null, [ Validators.required ] ],
    });
    this.titleBar.set(this.titleinfo);
    this.titleBar.setEvent.emit(this.titleinfo);
    this.validateForm.setValue({title:this.project.project_name,money:this.project.project_money,cycle:this.project.project_circle})
    this.ProText = this.project.project_detail
    this.sotreDataBeforUoload()
  }

  
  
  getText(protext:string){
    this.ProText = protext
  }
  
  submitForm(): void {
    var valid = true
    for (const i in this.validateForm.controls) {
      this.validateForm.controls[ i ].markAsDirty();
      this.validateForm.controls[ i ].updateValueAndValidity();
      if(this.validateForm.controls[i].status==="INVALID"){
        valid = false 
      }
    }
    if (!valid){
      return
    }

    if (this.ProText == ""){
      this.ProText = "<p>无</p>"
    }
    this.isLoading = true
    this.ProjectService.AmendProject({
      project_id:this.project.project_id,
      project_name:this.validateForm.get("title").value,
      project_circle:this.validateForm.get("cycle").value,
      project_money:this.validateForm.get("money").value,
      project_detail:this.ProText
    }).subscribe(
      res=>{
        this.isLoading = false
        this.message.success("修改成功", { nzDuration: 1000 })
        this.updateLocalProjectData()
        setTimeout(_=>{
          this.route.navigate([ROUTES.seeDetail.route])
        },1000)
      },
      err=>{
        this.isLoading = false
        console.log(err)
      }
    )
  }

  //将修改的最新数据写入到本地
  updateLocalProjectData(){
    this.project.project_name = this.validateForm.get("title").value
    this.project.project_circle = this.validateForm.get("cycle").value
    this.project.project_money = this.validateForm.get("money").value
    this.project.project_detail = this.ProText
    this.deliveryproject.setProject(this.project)
  }

  //刷新保存数据
  sotreDataBeforUoload(){
    window.onbeforeunload = (event)=>{
      this.updateLocalProjectData()
    }
  }

  ngOnDestroy(){
    this. AmendProject$.unsubscribe()
  }

}

