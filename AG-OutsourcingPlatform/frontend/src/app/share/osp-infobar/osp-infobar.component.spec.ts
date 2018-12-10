import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OspInfobarComponent } from './osp-infobar.component';

describe('OspInfobarComponent', () => {
  let component: OspInfobarComponent;
  let fixture: ComponentFixture<OspInfobarComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OspInfobarComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OspInfobarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
