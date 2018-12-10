import {Injectable} from '@angular/core';
import {ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot} from '@angular/router';
import {environment} from '../../environments/environment';
import {CookieService} from 'ngx-cookie-service';
import {PC_ROUTE, URL_WECHAT_LOGIN} from '../const/const_chencanxin';
import {MockLoginService} from '../share/mock-login/mock-login.service';

@Injectable()
export class LoginGuard implements CanActivate {
  constructor(
    private router: Router,
    private cookie: CookieService,
    private mockLoginService: MockLoginService,

) {
  }

//
  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {

    // store the targer url
    this.mockLoginService.redirectURL = state.url;

    var ag_outsourcingplatform = this.cookie.get('ag_outsourcingplatform');

    if (ag_outsourcingplatform) {
      return true;
    }
    else {
      if (environment.BUILD_TYPE == 'dev') {
        this.router.navigate([PC_ROUTE.mock_login.url]);
        return false;
      }
      if (environment.BUILD_TYPE == 'deploy') {
        let url = URL_WECHAT_LOGIN + '?redirectURL=' + state.url;
        console.log(url)
        window.open(url, '_self');
        return false;
      }
    }
  }
}
