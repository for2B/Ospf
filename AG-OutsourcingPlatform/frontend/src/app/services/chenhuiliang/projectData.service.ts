import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { FbData, SITE_HOST_URL, GET_PLATFORM_PROJECT ,GET_MYPUB_PROJECT, GET_MYCONT_PROJECT, PUBLISH_PROJECT, COUNT_COMPARTED_STATUS, GET_TEAMMATE, Project, JOIN_BIDDING, GIVE_UP_CONTRACTOR, UPDATE_PROJECT_STATUS, SELECT_CONTRACTOR, GET_PROJECT_IDENTITY, GIVE_UP_BIDDING, MODIFICATION_PROJECT } from '../../const/const_chenhuiliang';
import { PubHttpService } from '../HttpService/pub-http.service';
import { HttpParams } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ProjectService {

  GlobalProject:Project //全局Project变量

  constructor(
    private http:PubHttpService

  ) {}

  //获取平台列表
  GetPlatformProjectList(SiftOption:any):Observable<FbData>{
    return this.http.postjson(SITE_HOST_URL+GET_PLATFORM_PROJECT,SiftOption)
  }

  //获取的发布的项目列表
  GetMyPublishProject(CurrentPage:number,ShowNums:number,ProjectStatus?:number):Observable<FbData>{
    var httpParams  

    if (ProjectStatus) {
      httpParams = new HttpParams()
                    .set('project_status',ProjectStatus.toString())
                    .set('current_page', CurrentPage.toString())
                    .set('show_nums', ShowNums.toString())

    }else{
      httpParams  = new HttpParams()
                    .set('current_page', CurrentPage.toString())
                    .set('show_nums', ShowNums.toString())
    }

    return this.http.get(SITE_HOST_URL+GET_MYPUB_PROJECT, httpParams)
  }

  //获取我承接的项目列表
  GetMyContractProject(CurrentPage:number,ShowNums:number,ProjectStatus?:number):Observable<FbData>{
    var httpParams  

    if (ProjectStatus) {
      httpParams = new HttpParams()
                    .set('project_status',ProjectStatus.toString())
                    .set('current_page', CurrentPage.toString())
                    .set('show_nums', ShowNums.toString())

    }else{
      httpParams  = new HttpParams()
                    .set('current_page', CurrentPage.toString())
                    .set('show_nums', ShowNums.toString())
    }
    return this.http.get(SITE_HOST_URL+GET_MYCONT_PROJECT, httpParams)
  }

  //发布项目
  PublishProject(ProjectBaseInfo:any) :Observable<FbData>{
    return this.http.post(SITE_HOST_URL+PUBLISH_PROJECT,ProjectBaseInfo)
  }

  //修改项目
  AmendProject(ProjectChangeInfo:any):Observable<FbData>{
    return this.http.post(SITE_HOST_URL+MODIFICATION_PROJECT,ProjectChangeInfo)
  }

  //获取我发布的项目中个状态项目总数
  GetQuantityEachStatusMyPublish():Observable<FbData>{
    var httpParams = new HttpParams().set("list_type","my_publish")
    return this.http.get(SITE_HOST_URL+COUNT_COMPARTED_STATUS,httpParams)
  }

  //获取我承接的项目中个状态项目总数
  GetQuantityEachStatusMyContract():Observable<FbData>{
    var httpParams = new HttpParams().set("list_type","my_contract")
    return this.http.get(SITE_HOST_URL+COUNT_COMPARTED_STATUS,httpParams)
  }

  //获取竞标列表
  GetBiddingList(pageIndex: number = 1, pageSize: number = 10, sortField: string, sortOrder: string, genders: string[]): Observable<{}> {
    
    let params = new HttpParams()
    .append('page', `${pageIndex}`)
    .append('results', `${pageSize}`)
    .append('sortField', sortField)
    .append('sortOrder', sortOrder);
    genders.forEach(gender => {
      params = params.append('gender', gender);
    });
    return this.http.get('https://api.randomuser.me/', {
      params
    });
  }

  //获取团队成员
  GetTeammate(ProjectId:number):Observable<FbData>{
    var httpParams = new HttpParams().set("project_id",ProjectId.toString())
    return this.http.get(SITE_HOST_URL+GET_TEAMMATE,httpParams)
  }

  //参与竞标
  JoinBidding(ProjectId:number):Observable<FbData>{
    var httpparams = new HttpParams().set("project_id",ProjectId.toString())
    return this.http.get(SITE_HOST_URL+JOIN_BIDDING,httpparams)
  }

  //选择承包者
  SelectContractor(UserId:number,ProjectId:number):Observable<FbData>{
    var httpparams = new HttpParams()
    .set("user_id",UserId.toString())
    .set("project_id",ProjectId.toString())
    return this.http.get(SITE_HOST_URL+SELECT_CONTRACTOR,httpparams)
  }

  //放弃（取消）当前承接人
  GiveUpContractor(ProjectId:number):Observable<FbData>{
    var httpParams = new HttpParams().set("project_id",ProjectId.toString())
    return this.http.get(SITE_HOST_URL+GIVE_UP_CONTRACTOR,httpParams)
  }

  //更新项目状态
  UpdateProjectStatus(ProjectId:number,Status:number):Observable<FbData>{
    var httpparams = new HttpParams()
    .set("project_id",ProjectId.toString())
    .set("project_status",Status.toString())
    return this.http.get(SITE_HOST_URL+UPDATE_PROJECT_STATUS,httpparams)
  }

  //获取用户对ProjectId对应的项目的身份
  GetUserIdentity(ProjectId:number):Observable<FbData>{
    var httpparams = new HttpParams()
    .set("project_id",ProjectId.toString())
    return this.http.get(SITE_HOST_URL+GET_PROJECT_IDENTITY,httpparams)
  }

  //放弃竞标
  GiveUpBidding(ProjectId:number):Observable<FbData>{
    var httpparams = new HttpParams()
    .set("project_id",ProjectId.toString())
    return this.http.get(SITE_HOST_URL+GIVE_UP_BIDDING,httpparams)
  }


}
