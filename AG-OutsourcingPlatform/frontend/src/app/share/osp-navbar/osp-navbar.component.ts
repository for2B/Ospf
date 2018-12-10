import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-osp-navbar',
  templateUrl: './osp-navbar.component.html',
  styleUrls: ['./osp-navbar.component.css']
})
export class OspNavbarComponent implements OnInit {

  constructor(
  private  router : Router
  ) { 
  }
  username:string="陈灿鑫"
  ngOnInit() {

  }
  toeva(){
  }

  exit():void{
    this.router.navigateByUrl("pc/user/login")
  }
}
