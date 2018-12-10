import {Injectable} from '@angular/core';
import {NzMessageService} from 'ng-zorro-antd';
import {ModalService} from '../modal/modal.service';

@Injectable({
  providedIn: 'root',
})
export class ToastService {
  constructor(
    private message: NzMessageService,
    private modal:ModalService,
  ) {
  }

  error(content: string): void {
    this.message.create('error', content,{nzDuration:3000});
  };

  success(content: string): void {
    this.message.create('success', content);
  };

  warning(content: string): void {
    this.message.create('warning', content);
  };

}

 