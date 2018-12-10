import { Component, OnInit } from '@angular/core';
import {BuildCodeService} from './build-code.service';
import {ToastService} from '../../share/toast/toast.service';
import {Title} from '@angular/platform-browser';

@Component({
  selector: 'app-build-code',
  templateUrl: './build-code.component.html',
  styleUrls: ['./build-code.component.css']
})
export class BuildCodeComponent implements OnInit {

  isLoading=false;
  name="";
  controllerCode="";
  modelCode="";
  registCode="";
  registStatementCode="";
  url="";
  apiLevel="";
  developer="";
  constructor(
    private myService:BuildCodeService,
  ) { }

  ngOnInit() {
    this.myService.base.pagetitle.setTitle("新建后端API")
  }

  toBuildCode(){
    this.upperFirstLetter()

    this.controllerCode=this.myService.buildControllerFile(this.name,this.developer)
    this.modelCode=this.myService.buildModelFile(this.name,this.developer)
    this.registCode=this.myService.buildRegistCode(this.name,this.url,this.apiLevel,this.developer)
    this.registStatementCode=this.myService.buildRegistStatementCode(this.name)
    this.myService.postCode(this.controllerCode,this.name,
      this.modelCode,this.registCode,
      this.registStatementCode,this.developer,
      () =>{this.myService.base.modal.success("创建接口成功","请重启后端程序,新接口链接为:<br><a>http://localhost:6616/api/"+this.name+"</a>")}
      )
  }

  autoInput(){
    this.url=this.name;
  }

  upperFirstLetter(){
    let firstLetter=this.name.slice(0,1);
    firstLetter=firstLetter.toUpperCase();
    this.name=firstLetter+this.name.slice(1);
    console.log("this.name=",this.name);
    this.myService.base.toast.success("已成功为您格式化!")
  }
}
