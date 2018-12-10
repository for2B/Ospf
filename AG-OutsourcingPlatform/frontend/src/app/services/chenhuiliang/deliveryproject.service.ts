import { Injectable } from '@angular/core';
import { Project } from '../../const/const_chenhuiliang';
import { LocalStorageService } from 'angular-2-local-storage';

@Injectable({
  providedIn: 'root'
})
export class DeliveryprojectService {

  project:Project = new Project()
  // Identity:number
  constructor( private localStorageService: LocalStorageService) { 
    this.project = this.localStorageService.get("project")
    // this.Identity = this.localStorageService.get("identity")
  }

  getProject():Project{
    return this.project
  }

  // getIdentity():number{
  //   return this.Identity
  // }

  setProject(project:Project){
    this.project = project
    this.localStorageService.set("project",this.project)
  }
  // setIdentity(Identity:number){
  //   this.Identity = Identity
  //   this.localStorageService.set("identity",this.Identity)
  // }
  

}
