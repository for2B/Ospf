import { Component, OnInit } from '@angular/core';
import { Router, provideRoutes } from '@angular/router';
import { TitleInfo } from '../../../services/chencanxin/title-info';
import { TitleBarService } from '../../../share/osp-infobar/titleBar.service';
import { testeva, para, Eva, imeva } from '../testdata';
import { GET_EVA_TO_OTHER, WRITE_EVALUATE } from '../../../const/const_guochuting';
import { PubHttpService } from '../../../services/HttpService/pub-http.service'
// import { Card } from 'antd';
// const { Meta } = Card;


@Component({
    selector: 'app-eva-invitied',
    templateUrl: './eva-invitied.component.html',
    styleUrls: ['./eva-invitied.component.css'],
    }
)

export class EvaInvitiedComponent implements OnInit {
    constructor(
        private titleBar: TitleBarService ,
        private httpserve : PubHttpService,
    ) {}
    // titleinfo:TitleInfo=new TitleInfo("anticon anticon-solution","我的评价","在这里您可以看到与你有关的评价")
    evadata = testeva;
    // models:model[] = [{moduleID:"",moduleTitle:"",sectionArray:this.learningcontents}];

    aeva : imeva =
    {
      estimator_id :19,
      estimatored_id : 33,
      evaluate_time: '20180102',
      evaluate_detail: "第一个评价",
      project_ssociated_id: 15,
      is_private :false
    }
    im_eva(){
        //立即评价
        console.log("ima eva")
        let url = WRITE_EVALUATE;//+"?evaluate_id="+this.aeva.EvaluateId+"?EstimatorId="+this.aeva.EstimatorId+"?EstimatoredId="+this.aeva.EstimatoredId+"?EvaluateTime="+this.aeva.EvaluateTime+"?EvaluateDetail="+this.aeva.EvaluateDetail+"?ProjectSsociatedId="+this.aeva.ProjectSsociatedId+"?IsPrivate="+this.aeva.IsPrivate;
        // console.log("url1=", url)
        this.httpserve.post(url, JSON.stringify(this.aeva)).subscribe(
            Response => {
                let res:any={data:{qrcode_url:""}}
                res = Response;
            },
            err => {
                console.log("mioamioamiao")
            }
            
        )
    }
    look_him(){
        //查看个人
        console.log("look himself")
    }
    search(){
        console.log("search")
    }
    UserId : number;
    KeyUserName :string;
    CurrentPage : number;
    ShowNums : number;
    eva : Eva[];
    totalPage : number;
    para : para = {
        UserId : 1001,
        KeyUserName : '李雷17',
        CurrentPage : 1,
        ShowNums : 20
    }
    flag : boolean = true;
    geteva(UserId, KeyUserName, CurrentPage, ShowNums){
        //还需根据openid来判断是对用户的评价
        let url = GET_EVA_TO_OTHER+'?key_user_name=' + this.para.KeyUserName + '&current_page=' +this.para.CurrentPage +'&show_nums='+this.para.ShowNums;
        // console.log("url1=", url)
        this.httpserve.get(url).subscribe(
            Response => {
                let res:any={data:{qrcode_url:""}}
                res = Response;
                // console.log("res=", res)
                this.evadata = res.data.userArray
                // console.log("this.evadata=", this.evadata)
                if(this.evadata == null){
                    this.flag = false;
                }
            },
            err => {
                console.log("gun")
            }
            
        )
    }
    ngOnInit() {
        this.flag = true;
        this.geteva(1001, '张三',1, 20);
    // this.titleBar.set(this.titleinfo);
    // this.titleBar.setEvent.emit(this.titleinfo);
  }
}
