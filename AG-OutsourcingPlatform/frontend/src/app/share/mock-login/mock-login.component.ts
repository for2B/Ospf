import {Component, OnInit} from '@angular/core';
import {Title} from '@angular/platform-browser';
import {PubHttpService} from '../../services/HttpService/pub-http.service';
import {MockLoginService} from './mock-login.service';
import {ActivatedRoute, Route} from '@angular/router';
import {URL_MOCK_LOGIN} from '../../const/const_chencanxin'
@Component({
  selector: 'app-mock-login',
  templateUrl: './mock-login.component.html',
  styleUrls: ['./mock-login.component.css']
})
export class MockLoginComponent implements OnInit {
  redirectURL="";
  constructor(
    private title: Title,
              private http: PubHttpService,
              private mockLoginService: MockLoginService,
  private ActivatedRoute:ActivatedRoute
  ) { }

  ngOnInit() {
    this.title.setTitle('模拟登陆 ');
    this.redirectURL=this.ActivatedRoute.snapshot.paramMap.get('redirectURL');

  }

  radioValue = '3';
  isLoading = false;

  toMockLogin(): void {
    console.log('value=', this.radioValue);
    this.isLoading = true;
    setTimeout(_ => {
      this.isLoading = false;
    }, 800);
    let redirectURL=this.mockLoginService.redirectURL?this.mockLoginService.redirectURL:'/'
    if (redirectURL=="/mock_login")
    {
      redirectURL="/";
    }
    let url = URL_MOCK_LOGIN+'?redirectURL=' + redirectURL + '&userType=' + this.radioValue;
    console.log(url)

    window.open(url, '_self');
  }


}
