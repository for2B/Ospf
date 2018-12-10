import {Injectable} from '@angular/core';
import {PubHttpService} from '../../../services/HttpService/pub-http.service';
import {NORMAL_QRCODE_TYPE, REFEREE_QRCODE_TYPE, URL_FETCH_QRCODE} from '../../../const/const_chencanxin';
import {ModalService} from '../../../share/modal/modal.service';
import {Error} from 'tslint/lib/error';

@Injectable({
  providedIn: 'root',
})
export class InvitationCodeService {

  constructor(
    private http: PubHttpService,
    private modal: ModalService,
  ) {
  }

  /**
   * 获取用于邀请指定类型用户的二维码图片链接
   */
  getQrcode(_UserType: string, _next: (value: any) => void) {
    // check input
    if (_UserType != REFEREE_QRCODE_TYPE && _UserType != NORMAL_QRCODE_TYPE) {
      throw new Error(`the param can only be "NORMAL" or "REFEREE".`);
    }
    this.http.GetDataByHttpGet(
      URL_FETCH_QRCODE + _UserType,
      _data => {
        if (!_data.qrcode_url) {
          this.modal.error('The data structure returned by the server has no qrcode_url attribute.');
        }
        if (_data.qrcode_url == '') {
          this.modal.error('The qrcode_url value returned by the server is an empty string.');
        }
        _next(unescape(_data.qrcode_url));
      }
    );
  }

}
