import { Injectable } from '@angular/core';
import { Subject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UpdateStatusService {

  private ProjectStatusSubject = new Subject<MonitorStatus>();
  private PlafDataUpdateSubject = new Subject<any>()
  private IdentitySubject = new Subject<MonitorIdentity>()
  private MyPublishListDataSubject = new Subject<any>();
  private MyContractListSubject = new Subject<any>();
  constructor() { }

  ListenInProjectStatus():Observable<MonitorStatus>{
    return this.ProjectStatusSubject.asObservable();
  }

  PlafListenUpdateMsg():Observable<any>{
    return this.PlafDataUpdateSubject.asObservable();
  }

  ListenInIdentity():Observable<MonitorIdentity>{
    return this.IdentitySubject.asObservable();
  }

  ListenInMyPublishDataChange():Observable<ShowPlace>{
    return this.MyPublishListDataSubject.asObservable();
  }

  ListenInMyContractDataChange():Observable<ShowPlace>{
    return this.MyContractListSubject.asObservable();
  }
 

  UpdateProjectStatusLocal(projectStatus:number,projectId:number,showPlace:string):void{
    this.ProjectStatusSubject.next({ProjectStatus:projectStatus,ProjectId:projectId,ShowPlacd:showPlace});
  }

  PlafDataUpdate(){
    this.PlafDataUpdateSubject.next();
  }

  UpdateIdentityLocal(Identity:number,projectId:number,showPlace:string):void{
    this.IdentitySubject.next({Identity:Identity,ProjectId:projectId,ShowPlacd:showPlace});
  }

  UpdateMyPublishList(showPlace:string){
    this.MyPublishListDataSubject.next({ShowPlace:showPlace});
  }

  UpdateMyContractorList(showPlace:string){
    this.MyContractListSubject.next({ShowPlace:showPlace});
  }


 


}

export class MonitorStatus{
  ProjectStatus:number
  ProjectId:number
  ShowPlacd:string
}

export class MonitorIdentity{
  Identity:number
  ProjectId:number
  ShowPlacd:string
}

export class ShowPlace{
  ShowPlace:string
}