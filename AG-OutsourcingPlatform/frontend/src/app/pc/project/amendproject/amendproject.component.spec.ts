import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AmendprojectComponent } from './amendproject.component';

describe('AmendprojectComponent', () => {
  let component: AmendprojectComponent;
  let fixture: ComponentFixture<AmendprojectComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AmendprojectComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AmendprojectComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
