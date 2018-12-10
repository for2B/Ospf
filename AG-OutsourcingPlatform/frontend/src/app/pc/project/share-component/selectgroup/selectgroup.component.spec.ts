import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SelectgroupComponent } from './selectgroup.component';

describe('SelectgroupComponent', () => {
  let component: SelectgroupComponent;
  let fixture: ComponentFixture<SelectgroupComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SelectgroupComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SelectgroupComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
