import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BiddinglistComponent } from './biddinglist.component';

describe('BiddinglistComponent', () => {
  let component: BiddinglistComponent;
  let fixture: ComponentFixture<BiddinglistComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BiddinglistComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BiddinglistComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
