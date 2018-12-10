import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OspNavbarComponent } from './osp-navbar.component';

describe('OspNavbarComponent', () => {
  let component: OspNavbarComponent;
  let fixture: ComponentFixture<OspNavbarComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OspNavbarComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OspNavbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
