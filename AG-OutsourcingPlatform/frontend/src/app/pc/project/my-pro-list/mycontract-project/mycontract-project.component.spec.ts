import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MycontractProjectComponent } from './mycontract-project.component';

describe('MycontractProjectComponent', () => {
  let component: MycontractProjectComponent;
  let fixture: ComponentFixture<MycontractProjectComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MycontractProjectComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MycontractProjectComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
