import { Component, OnInit, Input, Output,EventEmitter, AfterViewInit } from '@angular/core';
import { StatusTabGroup, StatusTab, NumberOfEachStatus } from '../../../../const/const_chenhuiliang';

@Component({
  selector: 'app-statustab',
  templateUrl: './statustab.component.html',
  styleUrls: ['./statustab.component.css']
})
export class StatustabComponent implements OnInit,AfterViewInit {

  @Input()tabGroup:StatusTabGroup
  @Output() selectTab = new EventEmitter<number>();
  @Input() 
  set numberOfEachStatus(numberofeachstatus:NumberOfEachStatus){
    this.NumberOfEachStatus = numberofeachstatus
    for(var i = 0 ; i < this.tabGroup.statusTabs.length ; i++){
      switch(i){
        case 0:
          this.tabGroup.statusTabs[i].projectAmount = this.NumberOfEachStatus.total
          break;
        case 1:
          this.tabGroup.statusTabs[i].projectAmount = this.NumberOfEachStatus.status_one
          break;
        case 2:
          this.tabGroup.statusTabs[i].projectAmount = this.NumberOfEachStatus.status_two
          break;
        case 3:
          this.tabGroup.statusTabs[i].projectAmount = this.NumberOfEachStatus.status_three
          break;
        case 4:
          this.tabGroup.statusTabs[i].projectAmount = this.NumberOfEachStatus.status_four
          break;
        case 5:
          this.tabGroup.statusTabs[i].projectAmount = this.NumberOfEachStatus.status_five
          break; 
      }
    }
  }

  get numberOfEachStatus(){
    return this.NumberOfEachStatus
  } 

  NumberOfEachStatus:NumberOfEachStatus
  percentWidth:number
  constructor() { }

  ngOnInit() {
    this.percentWidth = 100/this.tabGroup.statusTabs.length
   
  }

  ngAfterViewInit(){
    var Tabs =  document.getElementsByClassName('statustab')
    for (var i = 0 ;i < Tabs.length; i++){
      var tab = <HTMLElement>Tabs[i]
      tab.style.width = `${this.percentWidth}%`
    }
  }


  tabclick(tab:StatusTab){
    tab.click(this.tabGroup)
    this.selectTab.emit(tab.status)
  }

}
