import { Injectable } from '@angular/core';
import { PubHttpService } from '../HttpService/pub-http.service';
import { Observable } from 'rxjs';
import { FbData, SITE_HOST_URL, GET_BIDDINGLIST } from '../../const/const_chenhuiliang';
import { HttpParams } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class BiddingService {

  constructor(
    private httpService:PubHttpService
  ) { }

  GetBiddingList(ProjectId:number,CurrentPage:number,ShowNums:number,SortOrder?:string):Observable<FbData>{
    var httpparams
    if (SortOrder){
      httpparams = new HttpParams()
      .set("project_id",ProjectId.toString())
      .set("current_page",CurrentPage.toString())
      .set("show_nums",ShowNums.toString())
      .set("sort_order",SortOrder)
    }else{
      httpparams = new HttpParams()
      .set("project_id",ProjectId.toString())
      .set("current_page",CurrentPage.toString())
      .set("show_nums",ShowNums.toString())
    }
    return this.httpService.get(SITE_HOST_URL+GET_BIDDINGLIST,httpparams)
  }


}
