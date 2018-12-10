import {Injectable} from '@angular/core';
import {NzMessageService, NzModalService} from 'ng-zorro-antd';
import {isNullOrUndefined} from 'util';
import {ErrorLine, Line} from 'tslint/lib/verify/lines';
import {TSError} from 'ts-node';

@Injectable({
  providedIn: 'root',
})
export class ModalService {

  constructor(private modalService: NzModalService) {
  }

  info(title:string,content:string): void {
    this.modalService.info({
      nzTitle: title,
      nzContent: content,
      nzOnOk: () => {}
    });
  }

  success(title:string,content:string): void {
    this.modalService.success({
      nzTitle: title,
      nzContent: content,
    });
  }

  error(content:string): void {
    this.modalService.error({
      nzTitle: "程序错误",
      nzContent: `${content}`,
    });
    throw new Error(`程序错误 : ${content}`)
  }

  apiError(code:string,msg:string,apiurl:string): void {
    this.modalService.error({
      nzTitle: "服务器返回错误",
      nzContent: `接口地址 : <br>${apiurl}<br><br>
                  错误码 : ${code} <br><br>
                  详细信息 : ${msg} `,
    });
    throw new Error(`错误码:${code};${msg}`)
  }

  warning(title:string,content:string): void {
    this.modalService.warning({
      nzTitle: title,
      nzContent: content,
    });
  }

}
