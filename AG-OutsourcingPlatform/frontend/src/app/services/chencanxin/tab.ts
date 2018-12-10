
export class tabGroup {
  constructor(tabs: tab[]) {
    this.tabs = tabs;
    this.selected=tabs[0];
    this.selected.click(this)
  }
  tabs: tab[] =[];
  selected: tab=new tab("");
}

export class tab {

  name:string;
  currentClass:currentClass;
  selector: string;

  constructor(name:string ){
    this.name=name;
    this.currentClass={'active':false,'menu':true};
    this.selector="";
  }

  setSelector(selector:string ){
    this.selector=selector;
    return this
  }

  click(tabgroup:tabGroup){
    tabgroup.selected=this;
    this.currentClass={'active':true,'menu':false};
    tabgroup.tabs.forEach(
      t=>{
        if(t.name!==this.name){
          t.currentClass={'active':false,'menu':true};
        }
      }
    )

  }
}

export class currentClass{
  active:boolean;
  menu:boolean;
}
