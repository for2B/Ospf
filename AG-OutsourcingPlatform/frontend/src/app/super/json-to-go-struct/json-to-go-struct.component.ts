import { Component, OnInit } from '@angular/core';
import {BaseService} from '../../services/chencanxin/baseService';

@Component({
  selector: 'app-json-to-go-struct',
  templateUrl: './json-to-go-struct.component.html',
  styleUrls: ['./json-to-go-struct.component.css']
})
export class JsonToGoStructComponent implements OnInit {

  constructor(private base:BaseService) { }

  ngOnInit() {
    this.base.pagetitle.setTitle("Json To Go Struct")
    this.base.toast.warning("当前页面加载比较缓慢,请耐心等待")
  }

}
