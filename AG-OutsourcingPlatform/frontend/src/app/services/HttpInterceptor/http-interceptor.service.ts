import {Injectable} from '@angular/core';
import {HttpErrorResponse, HttpEvent, HttpHandler, HttpRequest, HttpResponse} from '@angular/common/http';
import {Observable, throwError} from 'rxjs';
import {catchError, tap} from 'rxjs/operators';
import {Router} from '@angular/router';
import {ToastService} from '../../share/toast/toast.service';
import {ModalService} from '../../share/modal/modal.service';
import {FB, PC_ROUTE, URL_WECHAT_LOGIN} from '../../const/const_chencanxin';
import {environment} from '../../../environments/environment';
import {MockLoginService} from '../../share/mock-login/mock-login.service';


@Injectable({
  providedIn: 'root'
})
export class HttpInterceptorService {

  constructor(
    private rout: Router,
    private toast: ToastService,
    private route: Router,
    private modal: ModalService,
    private mockLoginService: MockLoginService,
  ) {
  }


  intercept(req: HttpRequest<any>, next: HttpHandler):
    Observable<HttpEvent<any>> {
    // check param
    if (!req) {
      this.modal.error('need req param.');
    }
    if (!next) {
      this.modal.error('need next param.');
    }

    const authReq = req.clone();
    return next.handle(authReq).pipe(
      tap(event => {
        if (event instanceof HttpResponse) {

          if (event.status == 200) {
            // check response
            if (!event.body) {
              this.modal.error('The server return nothing.');
              return;
            }
            if (!event.body.code) {
              if(event.body.code!=0){
                this.modal.error('The response struct returned has not code attribute.');

              }
              return;
            }

            /*  如果后端返回错误:
            *    1,如果是cookie为空,cookie过期,sid无效,则跳转至登陆页面
            *    2,其他错误弹出模态框打印错误码及相关信息
            * */

            switch (event.body.code) {
              case FB.GLOBAL_SUCCESS.code:
                // do nothing
                break;
              case FB.COOKIE_NULL.code:
              case FB.COOKIE_EMPTY.code:
              case FB.ERROR_INVALID_SID.code:
              case FB.SESSION_EXPIRED.code:
                this.toast.error('请重新登陆!')
                ;
                this.toLogin(this.route);
                return;
                break;

              default:
                this.modal.apiError(event.body.code.toString(), event.body.msg, event.url);
                return;
            }
          }
        }
        return event;
      })
      , catchError(err => {
        if (err instanceof HttpErrorResponse) {
          return throwError(err);
        }
        return throwError(err);
      })
    );
  }

  toLogin(route: any) {
    // check param
    if (!route) {
      this.modal.error('need the route param.');
    }

    if (environment.BUILD_TYPE == 'dev') {
      if (this.route.url != '/mock_login') {
        this.mockLoginService.redirectURL = this.route.url;
      }
      route.navigate([PC_ROUTE.mock_login.url]);
      return;
    }

    if (environment.BUILD_TYPE == 'deploy') {
      let url = URL_WECHAT_LOGIN + '?redirectURL=' + route.url;
      console.log(url);
      window.open(url, '_self');
      return;
    }

  }


}

