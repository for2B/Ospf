import {EventEmitter, Injectable} from '@angular/core';
import {TitleInfo} from '../../services/chencanxin/title-info';
import {ModalService} from '../../share/modal/modal.service';
import {ToastService} from '../../share/toast/toast.service';
import {Title} from '@angular/platform-browser';
import {PubHttpService} from '../HttpService/pub-http.service';

@Injectable({
  providedIn: 'root',
})
export class BaseService {

  constructor(
    public modal:ModalService,
    public toast:ToastService,
    public pagetitle:Title,
    public http:PubHttpService,
  ) { }


}
