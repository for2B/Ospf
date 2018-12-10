import {EventEmitter, Injectable} from '@angular/core';
import {TitleInfo} from '../../services/chencanxin/title-info';

@Injectable({
  providedIn: 'root',
})
export class TitleBarService {
  titleInfo:TitleInfo=new TitleInfo("anticon anticon-smile-o","未设置标题","未设置提示")
  constructor() { }

  setEvent:EventEmitter<TitleInfo>=new EventEmitter<TitleInfo>();
  set(_titleInfo:TitleInfo){
    this.setEvent.emit(_titleInfo)
  }
  get(){
    this.setEvent.subscribe(
      info=>this.titleInfo=info
    )
  }
}
