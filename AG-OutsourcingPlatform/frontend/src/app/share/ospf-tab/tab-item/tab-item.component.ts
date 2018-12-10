import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-tab-item',
  templateUrl: './tab-item.component.html',
  styleUrls: ['./tab-item.component.css']
})
export class TabItemComponent implements OnInit {

  constructor() { }

  @Input()

  ngOnInit() {
  }

}
