import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MockLoginComponent } from './mock-login.component';

describe('MockLoginComponent', () => {
  let component: MockLoginComponent;
  let fixture: ComponentFixture<MockLoginComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MockLoginComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MockLoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
