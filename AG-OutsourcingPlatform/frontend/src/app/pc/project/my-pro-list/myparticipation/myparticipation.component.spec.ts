import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MyparticipationComponent } from './myparticipation.component';

describe('MyparticipationComponent', () => {
  let component: MyparticipationComponent;
  let fixture: ComponentFixture<MyparticipationComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MyparticipationComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MyparticipationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
