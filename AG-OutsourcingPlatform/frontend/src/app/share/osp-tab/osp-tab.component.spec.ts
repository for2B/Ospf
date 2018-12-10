import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OspTabComponent } from './osp-tab.component';

describe('OspTabComponent', () => {
  let component: OspTabComponent;
  let fixture: ComponentFixture<OspTabComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OspTabComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OspTabComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
