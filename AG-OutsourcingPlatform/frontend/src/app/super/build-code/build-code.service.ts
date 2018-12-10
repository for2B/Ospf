import {Injectable} from '@angular/core';
import {ModalService} from '../../share/modal/modal.service';
import {BaseService} from '../../services/chencanxin/baseService';
import {URL_GENERATE_GOLANG_API_CODE} from '../../const/const_chencanxin';

@Injectable({
  providedIn: 'root'
})
export class BuildCodeService {

  constructor(
    private modal: ModalService,
    public base: BaseService,
  ) {
  }

  postCode(_controllerFileCode: string, _apiName: string, _modelFileCode: string, _registFunctionCode: string, _registStatement: string, _developer: string, next: () => void) {
    // check param
    if (!_controllerFileCode) {
      this.base.modal.error('need _controllerFileCode param.');
    }
    if (_controllerFileCode == '') {
      this.base.modal.error(' _controllerFileCode param can\'t be an empty string.');
    }
    if (!_modelFileCode) {
      this.base.modal.error('need _modelFileCode param.');
    }
    if (_modelFileCode == '') {
      this.base.modal.error(' _modelFileCode param can\'t be an empty string.');
    }
    if (!_registFunctionCode) {
      this.base.modal.error('need _registFunctionCode param.');
    }
    if (_registFunctionCode == '') {
      this.base.modal.error(' _registFunctionCode param can\'t be an empty string.');
    }
    if (!_registStatement) {
      this.base.modal.error('need _registStatement param.');
    }
    if (_registStatement == '') {
      this.base.modal.error(' _registStatement param can\'t be an empty string.');
    }
    if (!_developer) {
      this.base.modal.error('need _developer param.');
    }
    if (_developer == '') {
      this.base.modal.error(' _developer param can\'t be an empty string.');
    }
    if (!_apiName) {
      this.base.modal.error('need _apiName param.');
    }
    if (_apiName == '') {
      this.base.modal.error(' _apiName param can\'t be an empty string.');
    }


    let data = {
      apiName: _apiName,
      developer: _developer,
      controllorCode: _controllerFileCode,
      modelFileCode: _modelFileCode,
      registFunctionCode: _registFunctionCode,
      registStatement: _registStatement,
    };
    this.base.http.postjson(URL_GENERATE_GOLANG_API_CODE, data).subscribe(next);
  }

  //自动生产后端新接口的控制器代码
  buildControllerFile(name: string,developer:string) {
    //  check param
    if (!name) {
      this.modal.error('need name param.');
    }
    if (name.includes('Controller')) {
      this.modal.error('don\'t need controller end.');
    }

    let nameLower = name.toLowerCase();

    let baseCode = `
    package controller_${developer}
    
    import (
		"OutsourcingPlatform/ospf/model/model_${developer}"
		"net/http"
	  )
    
    type ${name}Controller struct {
      *model_${developer}.${name}Model
    }
    
    type ApiParamOf${name} struct {
      
    }
    
    type ApiResponseOf${name} struct {
      
    }
    
    func (c *${name}Controller) To${name}(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    
    
        //var apiParam ApiParamOf${name}
        //var apiResponse ApiResponseOf${name}
    	
    	
      	msg:="${name}Controller is work!"
      	return msg,nil
      	
    }
    `;
    if (baseCode == '') {
      this.modal.error('return output value is an empty string.');
    }
    return baseCode;
  }

  lowerFirstLetter(name){
    let firstLetter=name.slice(0,1);
    firstLetter=firstLetter.toLowerCase();
    name=firstLetter+name.slice(1);
    return name
  }
  //自动生产后端新接口的模型层代码
  buildModelFile(name: string,developer:string) {
    //  check param
    if (!name) {
      this.modal.error('need name param.');
    }
    if (name.includes('Model')) {
      this.modal.error('don\'t need "Model" in the end.');
    }
    if (!developer ) {
      this.modal.error('need developer param.');
    }

    let nameOfFirstLetterLower = this.lowerFirstLetter(name)

    let baseCode = `
  
      package model_${developer}
      
      import (
          "OutsourcingPlatform/ospf/infrastructure"
          "sync"
      )
      
      var once${name}Model sync.Once
      
      type ${name}Model struct {
        *infrastructure.Infrastructure
      }
      
      var ${nameOfFirstLetterLower}Model *${name}Model
      
      func Get${name}Model(i *infrastructure.Infrastructure) *${name}Model {
        once${name}Model.Do(func() {
          ${nameOfFirstLetterLower}Model = &${name}Model{
            Infrastructure: i,
          }
        })
        return ${nameOfFirstLetterLower}Model
      }
      
    `;
    if (baseCode == '') {
      this.modal.error('return output value is an empty string.');
    }
    return baseCode;
  }


  //自动生产后端新接口的注册函数体代码
  buildRegistCode(name: string, url: string, apiLevel: string,developer:string) {
    //  check param
    if (!name) {
      this.modal.error('need name param.');
    }
    if (name.includes('Model')) {
      this.modal.error('don\'t need "Model" in the end.');
    }

    if (!url) {
      this.modal.error('need url param.');
    }
    if (name.includes('api')) {
      this.modal.error('don\'t need "api" in the url param.');
    }

    if (!apiLevel) {
      this.modal.error('need apiLevel param.');
    }
    if (!apiLevel.includes('USER')) {
      this.modal.error('need "USER" in the start of apiUrl param.');
    }


    let baseCode = `
  
  func (s *httpServer) reg${name}Controller(r *mux.Router) {
      c := &controller_${developer}.${name}Controller{
        ${name}Model: model_${developer}.Get${name}Model(s.ctx.ospf.infrastructure),
      }
      r.Handle("/api/${url}", httpapi.Decorate(c.To${name}, 
        s.AllowOrigin,
        middleware.DefaultDecode, 
        middleware.Encode,
        middleware.Log(s.ctx.ospf.infrastructure.Logger),
        middleware.Session(s.ctx.ospf.infrastructure.SessionManager),
        middleware.Permission(constant.${apiLevel}),
        middleware.Online),
        )
  }
    `;
    if (baseCode == '') {
      this.modal.error('return output value is an empty string.');
    }
    return baseCode;
  }


  //自动生产后端新接口的注册函数体代码
  buildRegistStatementCode(name: string) {
    //  check param
    if (!name) {
      this.modal.error('need name param.');
    }
    if (name.includes('Model')) {
      this.modal.error('don\'t need "Model" in the end.');
    }

    let baseCode = `s.reg${name}Controller(s.router)`;
    if (baseCode == '') {
      this.modal.error('return output value is an empty string.');
    }
    return baseCode;
  }
}
