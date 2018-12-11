import { Member } from "../services/panjiawei/users";

export const SITE_HOST_URL = "http://chuil.online:6616";
// export const SITE_HOST_URL = "http://localhost:6616";
export const UPLOAD_FILE = '/api/upload_file';
export const PUBLISH_PROJECT = '/api/publish_project';
export const GET_PLATFORM_PROJECT = '/api/get_project_list'
export const GET_MYPUB_PROJECT="/api/get_project_publish"
export const GET_MYCONT_PROJECT="/api/get_project_contract"
export const COUNT_COMPARTED_STATUS = "/api/count_comparted_status"
export const GET_TEAMMATE = "/api/get_project_teammate"
export const GET_BIDDINGLIST = "/api/get_bidding_list"
export const JOIN_BIDDING = "/api/join_bidding"
export const GIVE_UP_CONTRACTOR = "/api/give_up_contract"
export const UPDATE_PROJECT_STATUS = "/api/update_project_status"
export const SELECT_CONTRACTOR = "/api/select_project_contractor"
export const GET_PROJECT_IDENTITY = "/api/get_project_identity"
export const GIVE_UP_BIDDING = "/api/give_up_bidding"
export const MODIFICATION_PROJECT = "/api/modification_project"

//路由
export const ROUTES = {
    plafprolist:{name:"平台项目列表",route:"/pc/project/pfprolist"},
    mypublishlist:{name:"我的项目列表",route:"/pc/project/myprolist"},
    seeDetail:{name:"查看项目详情",route:"/pc/project/detail"},
    amendProject:{name:"修改项目",route:"pc/project/amendproject"}
}
export interface FbData {
    code: number;
    msg: string;
    data:any;
}
 
export class Project{
    project_id : number
    project_name:string
    publisher_id : number
    project_publish_time :string
    project_circle:number
    project_money:number
    project_detail:string
    project_status:number
    competitors_nums:number
    constructor(options:{
        pro_id?:number,
        pro_name?:string,
        pub_id?:number,
        pro_pub_time?:string,
        pro_circle?:number,
        pro_money?:number,
        pro_detail?:string,
        pro_status?:number,
        comp_nums?:number,
    }={}){
        this.project_id = options.pro_id || 1;
        this.project_name = options.pro_name || '测试发布项目';
        this.publisher_id = options.pub_id || 2;
        this.project_publish_time = options.pro_pub_time || '8102-1-1';
        this.project_circle = options.pro_circle || 30;
        this.project_money = options.pro_money || 4500;
        this.project_detail = options.pro_detail || '<p>无</p>';
        this.project_status = options.pro_status || 1;
    }
}

export class ProjectList{
    total:number
    project_list:Project[]
    constructor(options:{
        total?:number,
        projects?:Project[]
    }={}){
        this.total = options.total||0,
        this.project_list = options.projects || []
    }
}


export class NumberOfEachStatus{
    total : number
    status_one : number
    status_two : number
    status_three : number
    status_four : number
    status_five : number
    constructor(options:{
        total?:number,
        status_one?:number,
        status_two?:number,
        status_three?:number,
        status_four?:number,
        status_five?:number,
    }={}){
        this.total = options.total || 0;
        this.status_one = options.status_one || 0;
        this.status_two = options.status_two || 0;
        this.status_three = options.status_three || 0;
        this.status_four = options.status_four || 0;
        this.status_five = options.status_five || 0;
    }
}

export class Bider{
    user_id              :number
    head_img             :string
    name                 :string
    expertise_field      :string[]
    address              :string
    finish_project_count :number
    identity             :number
    constructor(options:{
        user_id?              :number,
        head_img?             :string,
        name?                 :string,
        expertise_field?      :string[],
        address?              :string,
        finish_project_count? :number,
        identity?             :number,
    }={}){
        this.user_id = options.user_id||1;
        this.head_img = options.head_img || '';
        this.name = options.name || '李狗蛋';
        this.expertise_field = options.expertise_field || ['Android','ios','java'];
        this.address = options.address || '广东省 广州市 番禺区';
        this.finish_project_count = options.finish_project_count || 10;
        this.identity = options.identity|| 1
    }
}

export let Biddinglist:Bider[] = [
    {user_id:1,head_img:"",name:"李狗蛋",expertise_field:['Android','ios','java'],address:"广东省 广州市 番禺区",finish_project_count:10,identity:0},
    {user_id:2,head_img:"",name:"王富贵",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:12,identity:0},
    {user_id:3,head_img:"",name:"张三",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:13,identity:0},
    {user_id:4,head_img:"",name:"李四",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:2,identity:0},
    {user_id:5,head_img:"",name:"王五",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:1,identity:0},
    {user_id:6,head_img:"",name:"赵六",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:0,identity:0},
    {user_id:8,head_img:"",name:"赛文7",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:7,identity:0},
    {user_id:9,head_img:"",name:"狂小狗",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:23,identity:0},
    {user_id:10,head_img:"",name:"马哥",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:3,identity:0},
    {user_id:11,head_img:"",name:"啊光",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:5,identity:0},
    {user_id:12,head_img:"",name:"接q辣舞",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:7,identity:0},
    {user_id:13,head_img:"",name:"啊让男",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:17,identity:0},
    {user_id:14,head_img:"",name:"虎大将军",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:15,identity:0},
    {user_id:15,head_img:"",name:"7酱",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:8,identity:0},
    {user_id:16,head_img:"",name:"嘤嘤嘤",expertise_field:['vue','go'],address:"广东深圳",finish_project_count:2,identity:0},
    ]

export class Teammate{
    Uid:number
    Identity:number
    IsMe:number // 1 0
    ButtonStyle:number
    constructor(option:{
        Uid?:number,
        Identity?:number
        IsMe?:number
        ButtonStyle?:number
    }={}){
        this.Uid = option.Uid || 0
        this.Identity = option.Identity || 0
        this.IsMe = option.IsMe || 0
        this.ButtonStyle = option.ButtonStyle || 0
    }
}

export let teammates:Teammate[] = [
    new Teammate({Uid:22,Identity:1,IsMe:0}),
    new Teammate({Uid:22,Identity:0,IsMe:0}),
    new Teammate({Uid:22,Identity:0,IsMe:1}),
    new Teammate({Uid:22,Identity:0,IsMe:0}),
    new Teammate({Uid:22,Identity:0,IsMe:0}),
]



export class CurrentClass{
    statusText:boolean
    selectStatus:boolean
}

export class StatusTab{
    statusText:string
    status:number
    projectAmount:number
    currentClass:CurrentClass
    constructor(statustext:string,status:number,projectamount?:number){
        this.statusText = statustext
        this.status = status
        this.projectAmount = projectamount
    }

    click(statustabgroup:StatusTabGroup){
        statustabgroup.selector = this
        this.currentClass = {'selectStatus':true,'statusText':false};
        statustabgroup.statusTabs.forEach(
            tab=>{
                if(tab.statusText!==this.statusText){
                    tab.currentClass = {'selectStatus':false,'statusText':true}
                }
            }
        )
    }
}

export class StatusTabGroup{
    statusTabs:StatusTab[]
    selector:StatusTab = new StatusTab("",0,0)
    constructor(statusTabs:StatusTab[]){
        this.statusTabs = statusTabs
        this.selector = statusTabs[0]
        this.selector.click(this)
    }
}


