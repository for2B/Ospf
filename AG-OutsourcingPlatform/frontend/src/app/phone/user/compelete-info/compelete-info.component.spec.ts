import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CompeleteInfoComponent } from './compelete-info.component';

describe('CompeleteInfoComponent', () => {
  let component: CompeleteInfoComponent;
  let fixture: ComponentFixture<CompeleteInfoComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CompeleteInfoComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CompeleteInfoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
