import {Injectable} from '@angular/core';
import {HttpClient, HttpErrorResponse, HttpRequest} from '@angular/common/http';

import {Observable, throwError} from 'rxjs';
import {catchError} from 'rxjs/operators';

import {FbData} from '../../const/const_chenhuiliang';
import {ModalService} from '../../share/modal/modal.service';

@Injectable({
  providedIn: 'root'
})


export class PubHttpService {

  constructor(
    private httpclient: HttpClient,
    private modal: ModalService,
  ) {
  }

  /**
   * GetDataByHttpGet
   * @param url apiurl
   * @returns {Observable<FbData>}
   */
  GetDataByHttpGet(url: string, _next:(data)=>void) {
    this.httpclient.get<FbData>(url, {withCredentials: true})
      .pipe(
        catchError(this.handleErr(this.modal)),
      )
      .subscribe(
        // check response.data
        (response) => {
          if(!response.data){
            this.modal.error('The response struct returned has not data attribute.');
            return;
          }

          _next(response.data);
        }
      );

  }


  /**
   * GetDataByPost
   * @param _url apiurl
   * @param _body postdata
   * @returns {Observable<FbData>}
   */
  GetDataByPost(_url: string, _body: object,_next:(data)=>void) {
    this.httpclient.post<FbData>(
      _url,
      JSON.stringify(_body),
      {withCredentials: true})
      .pipe(catchError(this.handleErr(this.modal)))
      .subscribe(
        (response) => {
          // check response.data
          if(!response.data){
            this.modal.error('The response struct returned has not data attribute.');
            return;
          }
          _next(response.data);
        }
      );
  }


  /**
   * get
   * @param url apiurl
   * @param httpparams url parameter
   * @returns {Observable<FbData>}
   */
  get(url: string,httpparams?:any) {
    return this.httpclient.get<FbData>(url, {withCredentials: true,params:httpparams})
      .pipe(
        catchError(this.handleErr(this.modal)),
      );
  }

  /**
   * post
   * @param url apiurl
   * @param body JSON
   * @returns {Observable<FbData>}
   */
  post(url: string, body: any) {
    return this.httpclient.post<FbData>(url, body, {
        reportProgress: true, withCredentials: true,

      }
      ).pipe(
        catchError(this.handleErr(this.modal))
      );
    }

 



  /**
   * postjson
   * @param url apiurl
   * @param body postdata
   * @returns {Observable<FbData>}
   */
  postjson(url: string, body: object) {
    return this.httpclient.post<FbData>(url, JSON.stringify(body), {withCredentials: true}).pipe(catchError(this.handleErr(this.modal)));
  }


  //上传Ex文件
  UploadEx(url: string, formdata: any) {
    const req = new HttpRequest('POST', url, formdata, {
      reportProgress: true
    });
    return this.httpclient.request(req);
  }

  getWithOption(url: string, option: any) {
    return this.httpclient.get(url, option).pipe(catchError(this.handleErr(this.modal)));
  }


  private handleErr(modal: ModalService) {
    return function handleErr(error: HttpErrorResponse) {
      if (error.status == 0) {
        modal.error('服务无响应,是否断网了?');
      }
      if (error.status == 400) {
        modal.error('请求错误-400:' + error.message);
      }
      if (error.status == 404) {
        modal.error('请求地址不存在-404!');
      }
      if (error.status == 500) {
        modal.error('服务器错误-500!');
      }
      // return an observable with a user-facing error message
       return throwError(error);
    };
  }
}