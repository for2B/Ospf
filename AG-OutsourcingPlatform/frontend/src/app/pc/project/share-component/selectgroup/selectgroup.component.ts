import { Component, OnInit, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-selectgroup',
  templateUrl: './selectgroup.component.html',
  styleUrls: ['./selectgroup.component.css']
})
export class SelectgroupComponent implements OnInit {
  CycleSelectedValue : any
  MoneySelectedValue : any
  DateSelectedValue = [0,6311433600]
  @Output() SelectOption = new EventEmitter<any>();
  
  CycleOptions = [
    { label: '全部',   cycle_range_start:0,   cycle_range_end:999999999999},
    { label: '30天内', cycle_range_start:0,   cycle_range_end:30 },
    { label: '30~90天',cycle_range_start:30,  cycle_range_end:90 },
    { label: '90天以上',cycle_range_start:90, cycle_range_end:999999999999},
  ];
  MoneyOptions=[
    {label:'全部',    money_range_start:0,money_range_end:999999999999},
    {label:'1万以内', money_range_start:0,money_range_end:10000},
    {label:'1万~5万', money_range_start:10000,money_range_end:50000},
    {label:'5万~10万',money_range_start:50000,money_range_end:100000},
    {label:'10万以上',money_range_start:100000,money_range_end:999999999999},
  ]

  arr: number[] = [ 1, 2 , 3, 4, 5, 6, 7, 8, 9];

  constructor() { }

  ngOnInit() {
    this.CycleSelectedValue = this.CycleOptions[0]
    this.MoneySelectedValue = this.MoneyOptions[0]
  }

  CycleChange(value: { label: string, cycle_range_start:number,cycle_range_end:number }): void {
    this.CycleSelectedValue = value
    this.EmtiData()
  }
  MoneyChange(value: { label: string, money_range_start:number,money_range_end:number }): void {
    this.MoneySelectedValue = value
    this.EmtiData()
  }
  onChange(result: Date[]): void {
    this.DateSelectedValue[0] = Math.floor((result[0].getTime()/1000)-result[0].getHours()*3600-result[0].getMinutes()*60-result[0].getSeconds())
    this.DateSelectedValue[1] = Math.floor((result[1].getTime()/1000)-result[1].getHours()*3600-result[1].getMinutes()*60-result[1].getSeconds())
    this.EmtiData()
  }

  EmtiData(){ 
    this.SelectOption.emit({
      cycle_range_start:this.CycleSelectedValue.cycle_range_start,
      cycle_range_end:this.CycleSelectedValue.cycle_range_end,
      money_range_start:this.MoneySelectedValue.money_range_start,
      money_range_end:this.MoneySelectedValue.money_range_end,
      date_range_start:this.DateSelectedValue[0],
      date_range_end:this.DateSelectedValue[1],
    })
  }
}
