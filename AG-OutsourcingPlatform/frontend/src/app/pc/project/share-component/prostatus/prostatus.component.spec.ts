import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ProstatusComponent } from './prostatus.component';

describe('ProstatusComponent', () => {
  let component: ProstatusComponent;
  let fixture: ComponentFixture<ProstatusComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ProstatusComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProstatusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
