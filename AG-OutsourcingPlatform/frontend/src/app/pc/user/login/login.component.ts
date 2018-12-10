import { ToastService } from './../../../share/toast/toast.service';
import { Router } from '@angular/router';
import { PubHttpService } from './../../../services/HttpService/pub-http.service';
import { Component, OnInit } from '@angular/core';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  
  constructor(
   private httpService:PubHttpService,
   private router:Router,
   private toast:ToastService
  ) {}
 
  ngOnInit(){
    this.toast.warning("11")
  }
  phone:String;
  password:String;
 login(phone: String,password: String):void{
     const body={"Phone":phone,"Password":password}
    
     this.httpService.post("http://localhost:6616/api/password_login",body).subscribe(rep=>{
       console.log(rep)
       if(rep.code==200){
         this.router.navigateByUrl("")
       }
       else{
          this.toast.error(rep.msg)
       }
     })
 
 }
}

class User {
  Phone:String;
  Password:String;
}
