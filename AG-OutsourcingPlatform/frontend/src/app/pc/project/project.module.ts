import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AddMemmerComponent } from './add-memmer/add-memmer.component';
import {ShareModule} from '../../share/share.module';
import { PlatformProListComponent } from './platform-pro-list/platform-pro-list.component';
import { ProjectRoutingModule } from './project-routing.module';
import { ProjectitemComponent } from './share-component/projectitem/projectitem.component';
import { ProstatusComponent } from './share-component/prostatus/prostatus.component';
import { BottomButtonComponent } from './share-component/bottom-button/bottom-button.component';
import { SelectgroupComponent } from './share-component/selectgroup/selectgroup.component';
import { MyProListComponent } from './my-pro-list/my-pro-list.component';
import { MypublishProjectComponent } from './my-pro-list/mypublish-project/mypublish-project.component';
import { MycontractProjectComponent } from './my-pro-list/mycontract-project/mycontract-project.component';
import { PublishProjectComponent } from './publish-project/publish-project.component';
import {TeamMemberCardComponent} from './team-member-card/team-member-card.component';
import {MemberService} from '../../services/panjiawei/member.service';
import { ProjectDetailComponent } from './project-detail/project-detail.component';
import { BiddinglistComponent } from './share-component/biddinglist/biddinglist.component';
import { BidlistoperatebuttonComponent } from './share-component/bidlistoperatebutton/bidlistoperatebutton.component';
import { AmendprojectComponent } from './amendproject/amendproject.component';
import { MyparticipationComponent } from './my-pro-list/myparticipation/myparticipation.component';
import { StatustabComponent } from './share-component/statustab/statustab.component';

@NgModule({
  imports: [
    ShareModule,
    ProjectRoutingModule,

  ],
  declarations: [
    AddMemmerComponent, 
    PlatformProListComponent, 
    ProjectitemComponent, 
    ProstatusComponent, 
    BottomButtonComponent, 
    SelectgroupComponent,
    MyProListComponent,
    MypublishProjectComponent,
    MycontractProjectComponent,
    PublishProjectComponent,
    TeamMemberCardComponent,
    ProjectDetailComponent,
    BiddinglistComponent,
    BidlistoperatebuttonComponent,
    AmendprojectComponent,
    MyparticipationComponent,
    StatustabComponent,
    ],
  providers: [MemberService,]
})
export class ProjectModule { }
