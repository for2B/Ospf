import { Component, OnInit } from '@angular/core';
import {Time} from '@angular/common';

@Component({
  selector: 'app-super',
  templateUrl: './super.component.html',
  styleUrls: ['./super.component.css']
})
export class SuperComponent implements OnInit {
  gridStyle = {
    width    : '25%',
    textAlign: 'center',
    cursor:'pointer',
  };
  isCollapsed = false;
  cardtitle="";
  constructor() { }

  ngOnInit() {
    let date=new Date()
    this.cardtitle=date.toLocaleString()
    setInterval(
      ()=>{
        let date=new Date()
        this.cardtitle=date.toLocaleString()
      },
      1000
    )
  }

}
