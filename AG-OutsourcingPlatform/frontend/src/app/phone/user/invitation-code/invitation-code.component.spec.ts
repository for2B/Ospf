import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { InvitationCodeComponent } from './invitation-code.component';

describe('InvitationCodeComponent', () => {
  let component: InvitationCodeComponent;
  let fixture: ComponentFixture<InvitationCodeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ InvitationCodeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(InvitationCodeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
