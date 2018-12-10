import { Component, OnInit } from '@angular/core';
import {PubHttpService} from '../../services/HttpService/pub-http.service';

@Component({
  selector: 'app-set-gzhmenu',
  templateUrl: './set-gzhmenu.component.html',
  styleUrls: ['./set-gzhmenu.component.css']
})
export class SetGzhmenuComponent implements OnInit {

  constructor(private httpservice: PubHttpService) { }
  menujson: string = "";
  ngOnInit() {
    this.filljson();
  }
  submit() {
    localStorage.setItem("menujson",this.menujson);
    let data = {
      menujson: this.menujson
    }

    console.log("data=",data);

    this.httpservice.post("http://localhost:6616/api/test", JSON.stringify(data)).subscribe(
      res => {
        console.log(res);
      },
      error => {
        console.log(error);
      }
    );
  }

  filljson() {
    this.menujson = `{
      "button": [{
              "name": "我是老师",
              "sub_button": [{
                  "type": "view",
                  "name": "发布课程",
                  "url": "http://www.chencanxin.cn/courseplatform/teacher/addCourse"
              }]
          },
          {
              "name": "我是学生",
              "sub_button": [{
                  "type": "view",
                  "name": "购买课程",
                  "url": "http://www.chencanxin.cn/courseplatform/student/"
              }]
          },
          {
              "name": "我的",
              "sub_button": [{
                      "type": "view",
                      "name": "个人中心",
                      "url": "http://www.chencanxin.cn/courseplatform/api/WeChatLogin?redirectURL=auto_my"
                  },
                  {
                      "type": "view",
                      "name": "常见问题",
                      "url": "http://www.chencanxin.cn/courseplatform/student/"
                  },
                  {
                      "type": "view",
                      "name": "联系客服",
                      "url": "http://www.chencanxin.cn/courseplatform/student/"
                  }
              ]
          }
      ]
  }`;

  let store_menujson=localStorage.getItem("menujson")
  if (store_menujson&&store_menujson!="") {
    this.menujson=localStorage.getItem("menujson");
  }

  }


}
