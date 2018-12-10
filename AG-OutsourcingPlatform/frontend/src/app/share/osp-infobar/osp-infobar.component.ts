import {
  AfterContentChecked,
  AfterContentInit,
  AfterViewChecked,
  AfterViewInit,
  Component,
  DoCheck,
  Input,
  OnChanges,
  OnInit
} from '@angular/core';
import {TitleInfo} from '../../services/chencanxin/title-info';
import {TitleBarService} from './titleBar.service';

@Component({
  selector: 'app-osp-infobar',
  templateUrl: './osp-infobar.component.html',
  styleUrls: ['./osp-infobar.component.css']
})
export class OspInfobarComponent implements OnInit {
constructor(
  private barinfo:TitleBarService,
  ){}
  titleInfo: TitleInfo=new TitleInfo("init","init","init");

  ngOnInit() {
    this.titleInfo=this.barinfo.setEvent.subscribe(
      info=>this.titleInfo=info
    )
  }
}
