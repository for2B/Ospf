export class User {
  Uid: number;
  Name: string;
  Phone: string;
  Location: string;
  Create_time: Date;
  Change_time: Date;
  Permissions: number;
  State: number;
  checked: boolean;
}

export class UserFullDetail {
  User_id: number;
  Invited_id: number;
  User_name:	string;
  Openid:	string;
  Head_img:	string;
  User_tel:	string;
  User_sex: string;
  User_idcard: string;
  User_bank: string;
  Permissions: number;
  Regist_time: Date;
  Personal_update_time: Date;
  Permissiongs_update_time: Date;
  Prove_data: string;
  Expertise_field: string;
  Wechat: string;
  Address: string;
  Level: number;
  User_state: number;
  Invitation_qrcode_token_normal:	string;
  Invitation_qrcode_token_witness: string;
}

export class Member {
  Uid: number;
  Name: string;
  Phone: string;
  Location: string;
  Rating: number;
  Head_img:	string;
  Expertise_field: string[];
}

export const ACCEPT = {
  AC_1: 1,  // 登录，启用，禁用权限，都是这个啦
  AC_2: 2,  // 承接项目权限，不知道加上有什么用，加上就对了
  AC_4: 4,  // 发布项目权限
  AC_8: 8   // 这个就是评论权限了，万一有差评师，得禁言
};

