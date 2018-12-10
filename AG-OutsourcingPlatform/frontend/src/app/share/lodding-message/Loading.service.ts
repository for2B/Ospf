
import {Injectable} from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd';

@Injectable({
    providedIn: 'root',
  })
export class LoadingService {
    constructor(private message: NzMessageService) { }
    id : any
    ishide:boolean
    hide() {
        if (!this.id) {
        return this;
        }
        this.message.remove(this.id);
        this.ishide = true
        return this;
    }

    success(){
        if (!this.id){
        return
        }
        this.message.create('success', "获取数据成功!");
    }

    show(content?: string) {
        content = content ? content : '正在获取数据...';
        this.id = this.message.loading(content, {nzDuration: 0}).messageId;
        this.ishide = false
        setTimeout(_=>{
            if (this.ishide == false){
                this.hide()
                //TODO 可能需要做一些错误提示
            }
        },7000)
    }
}
  