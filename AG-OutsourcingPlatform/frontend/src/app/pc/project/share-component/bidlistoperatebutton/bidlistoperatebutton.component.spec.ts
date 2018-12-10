import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BidlistoperatebuttonComponent } from './bidlistoperatebutton.component';

describe('BidlistoperatebuttonComponent', () => {
  let component: BidlistoperatebuttonComponent;
  let fixture: ComponentFixture<BidlistoperatebuttonComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BidlistoperatebuttonComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BidlistoperatebuttonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
