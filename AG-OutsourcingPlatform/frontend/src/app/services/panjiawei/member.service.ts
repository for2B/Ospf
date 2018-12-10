import { Injectable } from '@angular/core';
import {HttpParams} from '@angular/common/http';
import {PubHttpService} from '../HttpService/pub-http.service';

@Injectable()
export class MemberService {

  constructor(private http: PubHttpService) { }
  getMember(uid: number) {
    const options = {
      params: new HttpParams().set('uid', uid.toString())
    };
    const url = 'http://localhost:6616/api/getmember';
    return this.http.getWithOption(url, options);
  }
}



