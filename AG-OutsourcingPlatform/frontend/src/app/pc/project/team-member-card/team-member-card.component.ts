import {Component, EventEmitter, Input, Output, OnInit} from '@angular/core';
import {Member} from '../../../services/panjiawei/users';
import {MemberService} from '../../../services/panjiawei/member.service';

@Component({
  selector: 'app-team-member-card',
  templateUrl: './team-member-card.component.html',
  styleUrls: ['./team-member-card.component.css']
})
export class TeamMemberCardComponent implements OnInit {


  @Input() uid: number;
  @Input() projectId: number;
  @Output() isDel = new EventEmitter<boolean>();
  @Input() buttonStyle: number;
  @Input() openPhone: boolean ;
  @Input() TeamIdentity:number //1 承包者 0普通成员
  TeammateName:string
  member: Member;
  color = ['magenta', 'geekblue', 'blue', 'cyan', 'green', 'lime', 'gold', 'orange', 'volcano', 'red'];
  constructor(private memberService: MemberService) { }
  ngOnInit() {
    this.memberService.getMember(this.uid).subscribe(res => {
      if (res['data']) {
        this.member = res['data'];
        console.log(this.member);
        this.spliceName()
      }
    });
  }
  DeleteMemberFromTeam() {
    this.isDel.emit(true);
  }
  spliceName(){
    if(this.TeamIdentity == 1){
      this.TeammateName = this.member.Name+"(承包者)"
    }else{
      this.TeammateName = this.member.Name
    }
  }

}
