import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MyProListComponent } from './my-pro-list.component';

describe('MyProListComponent', () => {
  let component: MyProListComponent;
  let fixture: ComponentFixture<MyProListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MyProListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MyProListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
