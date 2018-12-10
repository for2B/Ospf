import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-myparticipation',
  templateUrl: './myparticipation.component.html',
  styleUrls: ['./myparticipation.component.css']
})
export class MyparticipationComponent implements OnInit {

  statusOptions = ['等待承接','等待签订','签订合同','项目结束']
  nzSpan = 24/this.statusOptions.length
  constructor() { }

  ngOnInit() {
    console.log(this.nzSpan)
  }

}
