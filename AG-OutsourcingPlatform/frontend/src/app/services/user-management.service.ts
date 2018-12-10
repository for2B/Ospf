import {Injectable, Provider} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {User} from './panjiawei/users';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class UserManagementService {

  constructor(private http: HttpClient) { }
  getUser(uid: number): Observable<User> {
    const options = {
      params: new HttpParams().set('uid', uid.toString())
    };
    const url = 'http://localhost:6616/api/getuser';
    return this.http.get<User>(url, options);
  }
  getUsersList(page: number, limit: number, skeyword: string, permission: string, state: string) {
    if ( skeyword === '' ) {
      skeyword = '%';
    }
    const options = {
      params: new HttpParams()
        .set('page', page.toString())
        .set('limit', limit.toString())
        .set('permission', permission)
        .set('skeyword', skeyword)
        .set('state', state)
    };
    const url = 'http://localhost:6616/api/alluser';
    return this.http.get(url, options);
  }

  // operate_type有 “add”和“remove”，permission_type有“AC_1”,"AC_2","AC_4","AC_8"
  alterUserPermission(user_id: number[], operate_type: string, permission_type: string){
    const options = {user_id: user_id, operate_type: operate_type, permission_type: permission_type};
    const url = 'http://localhost:6616/api/alter_permission';
    return this.http.post(url, options);
  }

  deleteUsers(user_id: number[]) {
    const options = {user_id: user_id};
    const url = 'http://localhost:6616/api/del_users';
    return this.http.post(url, options);
  }
}
