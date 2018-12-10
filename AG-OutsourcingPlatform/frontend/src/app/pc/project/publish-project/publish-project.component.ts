import { Component, OnInit } from '@angular/core';
import {  FormBuilder,FormControl,FormGroup,Validators } from '@angular/forms';
import { TitleInfo } from '../../../services/chencanxin/title-info';
import { TitleBarService } from '../../../share/osp-infobar/titleBar.service';
import { PubHttpService } from '../../../services/HttpService/pub-http.service';
import {  PUBLISH_PROJECT, ROUTES } from '../../../const/const_chenhuiliang';
import { NzMessageService } from 'ng-zorro-antd';
import { Router } from '@angular/router';
import { ProjectService } from '../../../services/chenhuiliang/projectData.service';
@Component({
  selector: 'app-publish-project',
  templateUrl: './publish-project.component.html',
  styleUrls: ['./publish-project.component.css']
})
export class PublishProjectComponent implements OnInit {
  validateForm: FormGroup;
  ProText=""
  isLoading = false


  updateConfirmValidator(): void {
    /** wait for refresh value */
    Promise.resolve().then(() => this.validateForm.controls.checkPassword.updateValueAndValidity());
  }

  
 
  constructor(
    private fb: FormBuilder,
    private titleBar: TitleBarService,
    private message:NzMessageService,
    private route:Router,
    private ProjectService:ProjectService
  ) {}

  titleinfo: TitleInfo = new TitleInfo('anticon anticon-plus-circle-o' , '发布项目' , '发布项目需求到平台上等待承接');

  ngOnInit() {
    this.validateForm = this.fb.group({
      title            : [ null, [ Validators.required,Validators.maxLength(20) ] ],
      money            : [ null, [ Validators.required ] ],
      cycle            : [ null, [ Validators.required ] ],
    });
    this.titleBar.set(this.titleinfo);
    this.titleBar.setEvent.emit(this.titleinfo);
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
    this.ProjectService.PublishProject({
      project_name:this.validateForm.get("title").value,
      project_circle:this.validateForm.get("cycle").value,
      project_money:this.validateForm.get("money").value,
      project_detail:this.ProText
    }).subscribe(
      res=>{
        this.isLoading = false
        this.message.success("发布成功！", { nzDuration: 2000 })
        setTimeout(_=>{
          this.route.navigate([ROUTES.mypublishlist.route])
        },2000)
      },
      err=>{
        this.isLoading = false
        console.log(err)
      }
    )

  }
}
